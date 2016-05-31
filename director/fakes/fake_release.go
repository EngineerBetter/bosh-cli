// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cloudfoundry/bosh-init/director"
	semver "github.com/cppforlife/go-semi-semantic/version"
)

type FakeRelease struct {
	NameStub        func() string
	nameMutex       sync.RWMutex
	nameArgsForCall []struct{}
	nameReturns     struct {
		result1 string
	}
	VersionStub        func() semver.Version
	versionMutex       sync.RWMutex
	versionArgsForCall []struct{}
	versionReturns     struct {
		result1 semver.Version
	}
	VersionMarkStub        func(mark string) string
	versionMarkMutex       sync.RWMutex
	versionMarkArgsForCall []struct {
		mark string
	}
	versionMarkReturns struct {
		result1 string
	}
	CommitHashWithMarkStub        func(mark string) string
	commitHashWithMarkMutex       sync.RWMutex
	commitHashWithMarkArgsForCall []struct {
		mark string
	}
	commitHashWithMarkReturns struct {
		result1 string
	}
	JobsStub        func() ([]director.Job, error)
	jobsMutex       sync.RWMutex
	jobsArgsForCall []struct{}
	jobsReturns     struct {
		result1 []director.Job
		result2 error
	}
	PackagesStub        func() ([]director.Package, error)
	packagesMutex       sync.RWMutex
	packagesArgsForCall []struct{}
	packagesReturns     struct {
		result1 []director.Package
		result2 error
	}
	DeleteStub        func(force bool) error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		force bool
	}
	deleteReturns struct {
		result1 error
	}
}

func (fake *FakeRelease) Name() string {
	fake.nameMutex.Lock()
	fake.nameArgsForCall = append(fake.nameArgsForCall, struct{}{})
	fake.nameMutex.Unlock()
	if fake.NameStub != nil {
		return fake.NameStub()
	} else {
		return fake.nameReturns.result1
	}
}

func (fake *FakeRelease) NameCallCount() int {
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	return len(fake.nameArgsForCall)
}

func (fake *FakeRelease) NameReturns(result1 string) {
	fake.NameStub = nil
	fake.nameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeRelease) Version() semver.Version {
	fake.versionMutex.Lock()
	fake.versionArgsForCall = append(fake.versionArgsForCall, struct{}{})
	fake.versionMutex.Unlock()
	if fake.VersionStub != nil {
		return fake.VersionStub()
	} else {
		return fake.versionReturns.result1
	}
}

func (fake *FakeRelease) VersionCallCount() int {
	fake.versionMutex.RLock()
	defer fake.versionMutex.RUnlock()
	return len(fake.versionArgsForCall)
}

func (fake *FakeRelease) VersionReturns(result1 semver.Version) {
	fake.VersionStub = nil
	fake.versionReturns = struct {
		result1 semver.Version
	}{result1}
}

func (fake *FakeRelease) VersionMark(mark string) string {
	fake.versionMarkMutex.Lock()
	fake.versionMarkArgsForCall = append(fake.versionMarkArgsForCall, struct {
		mark string
	}{mark})
	fake.versionMarkMutex.Unlock()
	if fake.VersionMarkStub != nil {
		return fake.VersionMarkStub(mark)
	} else {
		return fake.versionMarkReturns.result1
	}
}

func (fake *FakeRelease) VersionMarkCallCount() int {
	fake.versionMarkMutex.RLock()
	defer fake.versionMarkMutex.RUnlock()
	return len(fake.versionMarkArgsForCall)
}

func (fake *FakeRelease) VersionMarkArgsForCall(i int) string {
	fake.versionMarkMutex.RLock()
	defer fake.versionMarkMutex.RUnlock()
	return fake.versionMarkArgsForCall[i].mark
}

func (fake *FakeRelease) VersionMarkReturns(result1 string) {
	fake.VersionMarkStub = nil
	fake.versionMarkReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeRelease) CommitHashWithMark(mark string) string {
	fake.commitHashWithMarkMutex.Lock()
	fake.commitHashWithMarkArgsForCall = append(fake.commitHashWithMarkArgsForCall, struct {
		mark string
	}{mark})
	fake.commitHashWithMarkMutex.Unlock()
	if fake.CommitHashWithMarkStub != nil {
		return fake.CommitHashWithMarkStub(mark)
	} else {
		return fake.commitHashWithMarkReturns.result1
	}
}

func (fake *FakeRelease) CommitHashWithMarkCallCount() int {
	fake.commitHashWithMarkMutex.RLock()
	defer fake.commitHashWithMarkMutex.RUnlock()
	return len(fake.commitHashWithMarkArgsForCall)
}

func (fake *FakeRelease) CommitHashWithMarkArgsForCall(i int) string {
	fake.commitHashWithMarkMutex.RLock()
	defer fake.commitHashWithMarkMutex.RUnlock()
	return fake.commitHashWithMarkArgsForCall[i].mark
}

func (fake *FakeRelease) CommitHashWithMarkReturns(result1 string) {
	fake.CommitHashWithMarkStub = nil
	fake.commitHashWithMarkReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeRelease) Jobs() ([]director.Job, error) {
	fake.jobsMutex.Lock()
	fake.jobsArgsForCall = append(fake.jobsArgsForCall, struct{}{})
	fake.jobsMutex.Unlock()
	if fake.JobsStub != nil {
		return fake.JobsStub()
	} else {
		return fake.jobsReturns.result1, fake.jobsReturns.result2
	}
}

func (fake *FakeRelease) JobsCallCount() int {
	fake.jobsMutex.RLock()
	defer fake.jobsMutex.RUnlock()
	return len(fake.jobsArgsForCall)
}

func (fake *FakeRelease) JobsReturns(result1 []director.Job, result2 error) {
	fake.JobsStub = nil
	fake.jobsReturns = struct {
		result1 []director.Job
		result2 error
	}{result1, result2}
}

func (fake *FakeRelease) Packages() ([]director.Package, error) {
	fake.packagesMutex.Lock()
	fake.packagesArgsForCall = append(fake.packagesArgsForCall, struct{}{})
	fake.packagesMutex.Unlock()
	if fake.PackagesStub != nil {
		return fake.PackagesStub()
	} else {
		return fake.packagesReturns.result1, fake.packagesReturns.result2
	}
}

func (fake *FakeRelease) PackagesCallCount() int {
	fake.packagesMutex.RLock()
	defer fake.packagesMutex.RUnlock()
	return len(fake.packagesArgsForCall)
}

func (fake *FakeRelease) PackagesReturns(result1 []director.Package, result2 error) {
	fake.PackagesStub = nil
	fake.packagesReturns = struct {
		result1 []director.Package
		result2 error
	}{result1, result2}
}

func (fake *FakeRelease) Delete(force bool) error {
	fake.deleteMutex.Lock()
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		force bool
	}{force})
	fake.deleteMutex.Unlock()
	if fake.DeleteStub != nil {
		return fake.DeleteStub(force)
	} else {
		return fake.deleteReturns.result1
	}
}

func (fake *FakeRelease) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeRelease) DeleteArgsForCall(i int) bool {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return fake.deleteArgsForCall[i].force
}

func (fake *FakeRelease) DeleteReturns(result1 error) {
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

var _ director.Release = new(FakeRelease)
