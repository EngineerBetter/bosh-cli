// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cloudfoundry/bosh-init/director"
)

type FakeReleaseSeries struct {
	NameStub        func() string
	nameMutex       sync.RWMutex
	nameArgsForCall []struct{}
	nameReturns     struct {
		result1 string
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

func (fake *FakeReleaseSeries) Name() string {
	fake.nameMutex.Lock()
	fake.nameArgsForCall = append(fake.nameArgsForCall, struct{}{})
	fake.nameMutex.Unlock()
	if fake.NameStub != nil {
		return fake.NameStub()
	} else {
		return fake.nameReturns.result1
	}
}

func (fake *FakeReleaseSeries) NameCallCount() int {
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	return len(fake.nameArgsForCall)
}

func (fake *FakeReleaseSeries) NameReturns(result1 string) {
	fake.NameStub = nil
	fake.nameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeReleaseSeries) Delete(force bool) error {
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

func (fake *FakeReleaseSeries) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeReleaseSeries) DeleteArgsForCall(i int) bool {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return fake.deleteArgsForCall[i].force
}

func (fake *FakeReleaseSeries) DeleteReturns(result1 error) {
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

var _ director.ReleaseSeries = new(FakeReleaseSeries)
