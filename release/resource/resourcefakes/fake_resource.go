// This file was generated by counterfeiter
package resourcefakes

import (
	"sync"

	"github.com/cloudfoundry/bosh-cli/crypto"
	"github.com/cloudfoundry/bosh-cli/release/resource"
)

type FakeResource struct {
	NameStub        func() string
	nameMutex       sync.RWMutex
	nameArgsForCall []struct{}
	nameReturns     struct {
		result1 string
	}
	FingerprintStub        func() string
	fingerprintMutex       sync.RWMutex
	fingerprintArgsForCall []struct{}
	fingerprintReturns     struct {
		result1 string
	}
	ArchivePathStub        func() string
	archivePathMutex       sync.RWMutex
	archivePathArgsForCall []struct{}
	archivePathReturns     struct {
		result1 string
	}
	ArchiveSHA1Stub        func() string
	archiveSHA1Mutex       sync.RWMutex
	archiveSHA1ArgsForCall []struct{}
	archiveSHA1Returns     struct {
		result1 string
	}
	BuildStub        func(dev, final resource.ArchiveIndex) error
	buildMutex       sync.RWMutex
	buildArgsForCall []struct {
		dev   resource.ArchiveIndex
		final resource.ArchiveIndex
	}
	buildReturns struct {
		result1 error
	}
	FinalizeStub        func(final resource.ArchiveIndex) error
	finalizeMutex       sync.RWMutex
	finalizeArgsForCall []struct {
		final resource.ArchiveIndex
	}
	finalizeReturns struct {
		result1 error
	}
	RehashWithCalculatorStub        func(calculator crypto.DigestCalculator) (resource.Resource, error)
	rehashWithCalculatorMutex       sync.RWMutex
	rehashWithCalculatorArgsForCall []struct {
		calculator crypto.DigestCalculator
	}
	rehashWithCalculatorReturns struct {
		result1 resource.Resource
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeResource) Name() string {
	fake.nameMutex.Lock()
	fake.nameArgsForCall = append(fake.nameArgsForCall, struct{}{})
	fake.recordInvocation("Name", []interface{}{})
	fake.nameMutex.Unlock()
	if fake.NameStub != nil {
		return fake.NameStub()
	} else {
		return fake.nameReturns.result1
	}
}

func (fake *FakeResource) NameCallCount() int {
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	return len(fake.nameArgsForCall)
}

func (fake *FakeResource) NameReturns(result1 string) {
	fake.NameStub = nil
	fake.nameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeResource) Fingerprint() string {
	fake.fingerprintMutex.Lock()
	fake.fingerprintArgsForCall = append(fake.fingerprintArgsForCall, struct{}{})
	fake.recordInvocation("Fingerprint", []interface{}{})
	fake.fingerprintMutex.Unlock()
	if fake.FingerprintStub != nil {
		return fake.FingerprintStub()
	} else {
		return fake.fingerprintReturns.result1
	}
}

func (fake *FakeResource) FingerprintCallCount() int {
	fake.fingerprintMutex.RLock()
	defer fake.fingerprintMutex.RUnlock()
	return len(fake.fingerprintArgsForCall)
}

func (fake *FakeResource) FingerprintReturns(result1 string) {
	fake.FingerprintStub = nil
	fake.fingerprintReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeResource) ArchivePath() string {
	fake.archivePathMutex.Lock()
	fake.archivePathArgsForCall = append(fake.archivePathArgsForCall, struct{}{})
	fake.recordInvocation("ArchivePath", []interface{}{})
	fake.archivePathMutex.Unlock()
	if fake.ArchivePathStub != nil {
		return fake.ArchivePathStub()
	} else {
		return fake.archivePathReturns.result1
	}
}

func (fake *FakeResource) ArchivePathCallCount() int {
	fake.archivePathMutex.RLock()
	defer fake.archivePathMutex.RUnlock()
	return len(fake.archivePathArgsForCall)
}

func (fake *FakeResource) ArchivePathReturns(result1 string) {
	fake.ArchivePathStub = nil
	fake.archivePathReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeResource) ArchiveSHA1() string {
	fake.archiveSHA1Mutex.Lock()
	fake.archiveSHA1ArgsForCall = append(fake.archiveSHA1ArgsForCall, struct{}{})
	fake.recordInvocation("ArchiveSHA1", []interface{}{})
	fake.archiveSHA1Mutex.Unlock()
	if fake.ArchiveSHA1Stub != nil {
		return fake.ArchiveSHA1Stub()
	} else {
		return fake.archiveSHA1Returns.result1
	}
}

func (fake *FakeResource) ArchiveSHA1CallCount() int {
	fake.archiveSHA1Mutex.RLock()
	defer fake.archiveSHA1Mutex.RUnlock()
	return len(fake.archiveSHA1ArgsForCall)
}

func (fake *FakeResource) ArchiveSHA1Returns(result1 string) {
	fake.ArchiveSHA1Stub = nil
	fake.archiveSHA1Returns = struct {
		result1 string
	}{result1}
}

func (fake *FakeResource) Build(dev resource.ArchiveIndex, final resource.ArchiveIndex) error {
	fake.buildMutex.Lock()
	fake.buildArgsForCall = append(fake.buildArgsForCall, struct {
		dev   resource.ArchiveIndex
		final resource.ArchiveIndex
	}{dev, final})
	fake.recordInvocation("Build", []interface{}{dev, final})
	fake.buildMutex.Unlock()
	if fake.BuildStub != nil {
		return fake.BuildStub(dev, final)
	} else {
		return fake.buildReturns.result1
	}
}

func (fake *FakeResource) BuildCallCount() int {
	fake.buildMutex.RLock()
	defer fake.buildMutex.RUnlock()
	return len(fake.buildArgsForCall)
}

func (fake *FakeResource) BuildArgsForCall(i int) (resource.ArchiveIndex, resource.ArchiveIndex) {
	fake.buildMutex.RLock()
	defer fake.buildMutex.RUnlock()
	return fake.buildArgsForCall[i].dev, fake.buildArgsForCall[i].final
}

func (fake *FakeResource) BuildReturns(result1 error) {
	fake.BuildStub = nil
	fake.buildReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeResource) Finalize(final resource.ArchiveIndex) error {
	fake.finalizeMutex.Lock()
	fake.finalizeArgsForCall = append(fake.finalizeArgsForCall, struct {
		final resource.ArchiveIndex
	}{final})
	fake.recordInvocation("Finalize", []interface{}{final})
	fake.finalizeMutex.Unlock()
	if fake.FinalizeStub != nil {
		return fake.FinalizeStub(final)
	} else {
		return fake.finalizeReturns.result1
	}
}

func (fake *FakeResource) FinalizeCallCount() int {
	fake.finalizeMutex.RLock()
	defer fake.finalizeMutex.RUnlock()
	return len(fake.finalizeArgsForCall)
}

func (fake *FakeResource) FinalizeArgsForCall(i int) resource.ArchiveIndex {
	fake.finalizeMutex.RLock()
	defer fake.finalizeMutex.RUnlock()
	return fake.finalizeArgsForCall[i].final
}

func (fake *FakeResource) FinalizeReturns(result1 error) {
	fake.FinalizeStub = nil
	fake.finalizeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeResource) RehashWithCalculator(calculator crypto.DigestCalculator) (resource.Resource, error) {
	fake.rehashWithCalculatorMutex.Lock()
	fake.rehashWithCalculatorArgsForCall = append(fake.rehashWithCalculatorArgsForCall, struct {
		calculator crypto.DigestCalculator
	}{calculator})
	fake.recordInvocation("RehashWithCalculator", []interface{}{calculator})
	fake.rehashWithCalculatorMutex.Unlock()
	if fake.RehashWithCalculatorStub != nil {
		return fake.RehashWithCalculatorStub(calculator)
	} else {
		return fake.rehashWithCalculatorReturns.result1, fake.rehashWithCalculatorReturns.result2
	}
}

func (fake *FakeResource) RehashWithCalculatorCallCount() int {
	fake.rehashWithCalculatorMutex.RLock()
	defer fake.rehashWithCalculatorMutex.RUnlock()
	return len(fake.rehashWithCalculatorArgsForCall)
}

func (fake *FakeResource) RehashWithCalculatorArgsForCall(i int) crypto.DigestCalculator {
	fake.rehashWithCalculatorMutex.RLock()
	defer fake.rehashWithCalculatorMutex.RUnlock()
	return fake.rehashWithCalculatorArgsForCall[i].calculator
}

func (fake *FakeResource) RehashWithCalculatorReturns(result1 resource.Resource, result2 error) {
	fake.RehashWithCalculatorStub = nil
	fake.rehashWithCalculatorReturns = struct {
		result1 resource.Resource
		result2 error
	}{result1, result2}
}

func (fake *FakeResource) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	fake.fingerprintMutex.RLock()
	defer fake.fingerprintMutex.RUnlock()
	fake.archivePathMutex.RLock()
	defer fake.archivePathMutex.RUnlock()
	fake.archiveSHA1Mutex.RLock()
	defer fake.archiveSHA1Mutex.RUnlock()
	fake.buildMutex.RLock()
	defer fake.buildMutex.RUnlock()
	fake.finalizeMutex.RLock()
	defer fake.finalizeMutex.RUnlock()
	fake.rehashWithCalculatorMutex.RLock()
	defer fake.rehashWithCalculatorMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeResource) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ resource.Resource = new(FakeResource)
