package vm

import (
	"time"

	bicloud "github.com/cloudfoundry/bosh-init/cloud"
	biconfig "github.com/cloudfoundry/bosh-init/config"
	biagentclient "github.com/cloudfoundry/bosh-init/deployment/agentclient"
	bias "github.com/cloudfoundry/bosh-init/deployment/applyspec"
	bidisk "github.com/cloudfoundry/bosh-init/deployment/disk"
	bideplmanifest "github.com/cloudfoundry/bosh-init/deployment/manifest"
	biui "github.com/cloudfoundry/bosh-init/ui"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	boshretry "github.com/cloudfoundry/bosh-utils/retrystrategy"
	boshsys "github.com/cloudfoundry/bosh-utils/system"
	"github.com/pivotal-golang/clock"
)

type VM interface {
	CID() string
	Exists() (bool, error)
	AgentClient() biagentclient.AgentClient
	WaitUntilReady(timeout time.Duration, delay time.Duration) error
	Start() error
	Stop() error
	Apply(bias.ApplySpec) error
	UpdateDisks(bideplmanifest.DiskPool, biui.Stage) ([]bidisk.Disk, error)
	WaitToBeRunning(maxAttempts int, delay time.Duration) error
	AttachDisk(bidisk.Disk) error
	DetachDisk(bidisk.Disk) error
	Disks() ([]bidisk.Disk, error)
	UnmountDisk(bidisk.Disk) error
	MigrateDisk() error
	Delete() error
}

type vm struct {
	cid          string
	vmRepo       biconfig.VMRepo
	stemcellRepo biconfig.StemcellRepo
	diskDeployer DiskDeployer
	agentClient  biagentclient.AgentClient
	cloud        bicloud.Cloud
	fs           boshsys.FileSystem
	logger       boshlog.Logger
	logTag       string
}

func NewVM(
	cid string,
	vmRepo biconfig.VMRepo,
	stemcellRepo biconfig.StemcellRepo,
	diskDeployer DiskDeployer,
	agentClient biagentclient.AgentClient,
	cloud bicloud.Cloud,
	fs boshsys.FileSystem,
	logger boshlog.Logger,
) VM {
	return &vm{
		cid:          cid,
		vmRepo:       vmRepo,
		stemcellRepo: stemcellRepo,
		diskDeployer: diskDeployer,
		agentClient:  agentClient,
		cloud:        cloud,
		fs:           fs,
		logger:       logger,
		logTag:       "vm",
	}
}

func (vm *vm) CID() string {
	return vm.cid
}

func (vm *vm) Exists() (bool, error) {
	exists, err := vm.cloud.HasVM(vm.cid)
	if err != nil {
		return false, bosherr.WrapErrorf(err, "Checking existance of VM '%s'", vm.cid)
	}
	return exists, nil
}

func (vm *vm) AgentClient() biagentclient.AgentClient {
	return vm.agentClient
}

func (vm *vm) WaitUntilReady(timeout time.Duration, delay time.Duration) error {
	agentPingRetryable := biagentclient.NewPingRetryable(vm.agentClient)
	timeService := clock.NewClock() //TODO: inject timeService
	agentPingRetryStrategy := boshretry.NewTimeoutRetryStrategy(timeout, delay, agentPingRetryable, timeService, vm.logger)
	return agentPingRetryStrategy.Try()
}

func (vm *vm) Start() error {
	vm.logger.Debug(vm.logTag, "Starting agent")
	err := vm.agentClient.Start()
	if err != nil {
		return bosherr.WrapError(err, "Starting agent")
	}

	return nil
}

func (vm *vm) Stop() error {
	vm.logger.Debug(vm.logTag, "Stopping agent")
	err := vm.agentClient.Stop()
	if err != nil {
		return bosherr.WrapError(err, "Stopping agent")
	}

	return nil
}

func (vm *vm) Apply(newState bias.ApplySpec) error {
	vm.logger.Debug(vm.logTag, "Sending apply message to the agent with '%#v'", newState)
	err := vm.agentClient.Apply(newState)
	if err != nil {
		return bosherr.WrapError(err, "Sending apply spec to agent")
	}

	return nil
}

func (vm *vm) UpdateDisks(diskPool bideplmanifest.DiskPool, eventLoggerStage biui.Stage) ([]bidisk.Disk, error) {
	disks, err := vm.diskDeployer.Deploy(diskPool, vm.cloud, vm, eventLoggerStage)
	if err != nil {
		return disks, bosherr.WrapError(err, "Deploying disk")
	}
	return disks, nil
}

func (vm *vm) WaitToBeRunning(maxAttempts int, delay time.Duration) error {
	agentGetStateRetryable := biagentclient.NewGetStateRetryable(vm.agentClient)
	agentGetStateRetryStrategy := boshretry.NewAttemptRetryStrategy(maxAttempts, delay, agentGetStateRetryable, vm.logger)
	return agentGetStateRetryStrategy.Try()
}

func (vm *vm) AttachDisk(disk bidisk.Disk) error {
	err := vm.cloud.AttachDisk(vm.cid, disk.CID())
	if err != nil {
		return bosherr.WrapError(err, "Attaching disk in the cloud")
	}

	err = vm.agentClient.MountDisk(disk.CID())
	if err != nil {
		return bosherr.WrapError(err, "Mounting disk")
	}

	return nil
}

func (vm *vm) DetachDisk(disk bidisk.Disk) error {
	err := vm.cloud.DetachDisk(vm.cid, disk.CID())
	if err != nil {
		return bosherr.WrapError(err, "Detaching disk in the cloud")
	}

	return nil
}

func (vm *vm) Disks() ([]bidisk.Disk, error) {
	result := []bidisk.Disk{}

	disks, err := vm.agentClient.ListDisk()
	if err != nil {
		return result, bosherr.WrapError(err, "Listing vm disks")
	}

	for _, diskCID := range disks {
		disk := bidisk.NewDisk(biconfig.DiskRecord{CID: diskCID}, nil, nil)
		result = append(result, disk)
	}

	return result, nil
}

func (vm *vm) UnmountDisk(disk bidisk.Disk) error {
	return vm.agentClient.UnmountDisk(disk.CID())
}

func (vm *vm) MigrateDisk() error {
	return vm.agentClient.MigrateDisk()
}

func (vm *vm) Delete() error {
	deleteErr := vm.cloud.DeleteVM(vm.cid)
	if deleteErr != nil {
		// allow VMNotFoundError for idempotency
		cloudErr, ok := deleteErr.(bicloud.Error)
		if !ok || cloudErr.Type() != bicloud.VMNotFoundError {
			return bosherr.WrapError(deleteErr, "Deleting vm in the cloud")
		}
	}

	err := vm.vmRepo.ClearCurrent()
	if err != nil {
		return bosherr.WrapError(err, "Deleting vm from vm repo")
	}

	err = vm.stemcellRepo.ClearCurrent()
	if err != nil {
		return bosherr.WrapError(err, "Clearing current stemcell from stemcell repo")
	}

	// returns bicloud.Error only if it is a VMNotFoundError
	return deleteErr
}
