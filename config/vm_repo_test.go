package config_test

import (
	. "github.com/cloudfoundry/bosh-cli/config"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	fakesys "github.com/cloudfoundry/bosh-utils/system/fakes"
	fakeuuid "github.com/cloudfoundry/bosh-utils/uuid/fakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("VMRepo", func() {
	var (
		repo                   VMRepo
		deploymentStateService DeploymentStateService
		fs                     *fakesys.FakeFileSystem
		fakeUUIDGenerator      *fakeuuid.FakeGenerator
	)

	BeforeEach(func() {
		logger := boshlog.NewLogger(boshlog.LevelNone)
		fs = fakesys.NewFakeFileSystem()
		fakeUUIDGenerator = &fakeuuid.FakeGenerator{}
		deploymentStateService = NewFileSystemDeploymentStateService(fs, fakeUUIDGenerator, logger, "/fake/path")
		repo = NewVMRepo(deploymentStateService)
	})

	Describe("FindCurrent", func() {
		Context("when a current vm cid is set", func() {
			BeforeEach(func() {
				err := repo.UpdateCurrent("fake-vm-cid")
				Expect(err).ToNot(HaveOccurred())
			})

			It("returns current manifest sha1", func() {
				record, found, err := repo.FindCurrent()
				Expect(err).ToNot(HaveOccurred())
				Expect(found).To(BeTrue())
				Expect(record).To(Equal("fake-vm-cid"))
			})
		})

		Context("when a current vm cid is not set", func() {
			It("returns false", func() {
				_, found, err := repo.FindCurrent()
				Expect(err).ToNot(HaveOccurred())
				Expect(found).To(BeFalse())
			})
		})
	})

	Describe("UpdateCurrent", func() {
		It("updates vm cid", func() {
			err := repo.UpdateCurrent("fake-vm-cid")
			Expect(err).ToNot(HaveOccurred())

			deploymentState, err := deploymentStateService.Load()
			Expect(err).ToNot(HaveOccurred())

			expectedConfig := DeploymentState{
				DirectorID:   "fake-uuid-0",
				CurrentVMCID: "fake-vm-cid",
			}
			Expect(deploymentState).To(Equal(expectedConfig))
		})
	})

	Describe("ClearCurrent", func() {
		It("updates vm cid", func() {
			err := repo.ClearCurrent()
			Expect(err).ToNot(HaveOccurred())

			deploymentState, err := deploymentStateService.Load()
			Expect(err).ToNot(HaveOccurred())

			expectedConfig := DeploymentState{
				DirectorID:   "fake-uuid-0",
				CurrentVMCID: "",
			}
			Expect(deploymentState).To(Equal(expectedConfig))

			_, found, err := repo.FindCurrent()
			Expect(err).ToNot(HaveOccurred())
			Expect(found).To(BeFalse())
		})
	})

	Describe("FindCurrentIP", func() {
		Context("when current IP is specified", func() {
			BeforeEach(func() {
				err := repo.UpdateCurrentIP("10.10.1.3")
				Expect(err).ToNot(HaveOccurred())
			})

			It("returns current IP", func() {
				ip, found, err := repo.FindCurrentIP()
				Expect(err).ToNot(HaveOccurred())
				Expect(found).To(BeTrue())
				Expect(ip).To(Equal("10.10.1.3"))
			})
		})

		Context("when current IP is not valid", func() {
			BeforeEach(func() {
				err := repo.UpdateCurrentIP("fake-ip-address")
				Expect(err).ToNot(HaveOccurred())
			})

			It("returns an error", func() {
				ip, found, err := repo.FindCurrentIP()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("is not a valid IP address"))
				Expect(found).To(BeFalse())
				Expect(ip).To(Equal(""))
			})
		})

		Context("when current IP is not specified", func() {
			It("returns false", func() {
				ip, found, err := repo.FindCurrentIP()
				Expect(err).ToNot(HaveOccurred())
				Expect(found).To(BeFalse())
				Expect(ip).To(Equal(""))
			})
		})
	})

	Describe("UpdateCurrentIP", func() {
		It("updates vm current IP", func() {
			err := repo.UpdateCurrentIP("10.10.1.3")
			Expect(err).ToNot(HaveOccurred())

			deploymentState, err := deploymentStateService.Load()
			Expect(err).ToNot(HaveOccurred())

			expectedConfig := DeploymentState{
				DirectorID: "fake-uuid-0",
				CurrentIP:  "10.10.1.3",
			}
			Expect(deploymentState).To(Equal(expectedConfig))
		})
	})
})
