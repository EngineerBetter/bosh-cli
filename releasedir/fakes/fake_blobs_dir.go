// This file was generated by counterfeiter
package fakes

import (
	"io"
	"sync"

	"github.com/cloudfoundry/bosh-init/releasedir"
)

type FakeBlobsDir struct {
	InitStub        func() error
	initMutex       sync.RWMutex
	initArgsForCall []struct{}
	initReturns     struct {
		result1 error
	}
	BlobsStub        func() ([]releasedir.Blob, error)
	blobsMutex       sync.RWMutex
	blobsArgsForCall []struct{}
	blobsReturns     struct {
		result1 []releasedir.Blob
		result2 error
	}
	DownloadBlobsStub        func() error
	downloadBlobsMutex       sync.RWMutex
	downloadBlobsArgsForCall []struct{}
	downloadBlobsReturns     struct {
		result1 error
	}
	UploadBlobsStub        func() error
	uploadBlobsMutex       sync.RWMutex
	uploadBlobsArgsForCall []struct{}
	uploadBlobsReturns     struct {
		result1 error
	}
	TrackBlobStub        func(string, io.ReadCloser) (releasedir.Blob, error)
	trackBlobMutex       sync.RWMutex
	trackBlobArgsForCall []struct {
		arg1 string
		arg2 io.ReadCloser
	}
	trackBlobReturns struct {
		result1 releasedir.Blob
		result2 error
	}
	UntrackBlobStub        func(string) error
	untrackBlobMutex       sync.RWMutex
	untrackBlobArgsForCall []struct {
		arg1 string
	}
	untrackBlobReturns struct {
		result1 error
	}
}

func (fake *FakeBlobsDir) Init() error {
	fake.initMutex.Lock()
	fake.initArgsForCall = append(fake.initArgsForCall, struct{}{})
	fake.initMutex.Unlock()
	if fake.InitStub != nil {
		return fake.InitStub()
	} else {
		return fake.initReturns.result1
	}
}

func (fake *FakeBlobsDir) InitCallCount() int {
	fake.initMutex.RLock()
	defer fake.initMutex.RUnlock()
	return len(fake.initArgsForCall)
}

func (fake *FakeBlobsDir) InitReturns(result1 error) {
	fake.InitStub = nil
	fake.initReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBlobsDir) Blobs() ([]releasedir.Blob, error) {
	fake.blobsMutex.Lock()
	fake.blobsArgsForCall = append(fake.blobsArgsForCall, struct{}{})
	fake.blobsMutex.Unlock()
	if fake.BlobsStub != nil {
		return fake.BlobsStub()
	} else {
		return fake.blobsReturns.result1, fake.blobsReturns.result2
	}
}

func (fake *FakeBlobsDir) BlobsCallCount() int {
	fake.blobsMutex.RLock()
	defer fake.blobsMutex.RUnlock()
	return len(fake.blobsArgsForCall)
}

func (fake *FakeBlobsDir) BlobsReturns(result1 []releasedir.Blob, result2 error) {
	fake.BlobsStub = nil
	fake.blobsReturns = struct {
		result1 []releasedir.Blob
		result2 error
	}{result1, result2}
}

func (fake *FakeBlobsDir) DownloadBlobs() error {
	fake.downloadBlobsMutex.Lock()
	fake.downloadBlobsArgsForCall = append(fake.downloadBlobsArgsForCall, struct{}{})
	fake.downloadBlobsMutex.Unlock()
	if fake.DownloadBlobsStub != nil {
		return fake.DownloadBlobsStub()
	} else {
		return fake.downloadBlobsReturns.result1
	}
}

func (fake *FakeBlobsDir) DownloadBlobsCallCount() int {
	fake.downloadBlobsMutex.RLock()
	defer fake.downloadBlobsMutex.RUnlock()
	return len(fake.downloadBlobsArgsForCall)
}

func (fake *FakeBlobsDir) DownloadBlobsReturns(result1 error) {
	fake.DownloadBlobsStub = nil
	fake.downloadBlobsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBlobsDir) UploadBlobs() error {
	fake.uploadBlobsMutex.Lock()
	fake.uploadBlobsArgsForCall = append(fake.uploadBlobsArgsForCall, struct{}{})
	fake.uploadBlobsMutex.Unlock()
	if fake.UploadBlobsStub != nil {
		return fake.UploadBlobsStub()
	} else {
		return fake.uploadBlobsReturns.result1
	}
}

func (fake *FakeBlobsDir) UploadBlobsCallCount() int {
	fake.uploadBlobsMutex.RLock()
	defer fake.uploadBlobsMutex.RUnlock()
	return len(fake.uploadBlobsArgsForCall)
}

func (fake *FakeBlobsDir) UploadBlobsReturns(result1 error) {
	fake.UploadBlobsStub = nil
	fake.uploadBlobsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBlobsDir) TrackBlob(arg1 string, arg2 io.ReadCloser) (releasedir.Blob, error) {
	fake.trackBlobMutex.Lock()
	fake.trackBlobArgsForCall = append(fake.trackBlobArgsForCall, struct {
		arg1 string
		arg2 io.ReadCloser
	}{arg1, arg2})
	fake.trackBlobMutex.Unlock()
	if fake.TrackBlobStub != nil {
		return fake.TrackBlobStub(arg1, arg2)
	} else {
		return fake.trackBlobReturns.result1, fake.trackBlobReturns.result2
	}
}

func (fake *FakeBlobsDir) TrackBlobCallCount() int {
	fake.trackBlobMutex.RLock()
	defer fake.trackBlobMutex.RUnlock()
	return len(fake.trackBlobArgsForCall)
}

func (fake *FakeBlobsDir) TrackBlobArgsForCall(i int) (string, io.ReadCloser) {
	fake.trackBlobMutex.RLock()
	defer fake.trackBlobMutex.RUnlock()
	return fake.trackBlobArgsForCall[i].arg1, fake.trackBlobArgsForCall[i].arg2
}

func (fake *FakeBlobsDir) TrackBlobReturns(result1 releasedir.Blob, result2 error) {
	fake.TrackBlobStub = nil
	fake.trackBlobReturns = struct {
		result1 releasedir.Blob
		result2 error
	}{result1, result2}
}

func (fake *FakeBlobsDir) UntrackBlob(arg1 string) error {
	fake.untrackBlobMutex.Lock()
	fake.untrackBlobArgsForCall = append(fake.untrackBlobArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.untrackBlobMutex.Unlock()
	if fake.UntrackBlobStub != nil {
		return fake.UntrackBlobStub(arg1)
	} else {
		return fake.untrackBlobReturns.result1
	}
}

func (fake *FakeBlobsDir) UntrackBlobCallCount() int {
	fake.untrackBlobMutex.RLock()
	defer fake.untrackBlobMutex.RUnlock()
	return len(fake.untrackBlobArgsForCall)
}

func (fake *FakeBlobsDir) UntrackBlobArgsForCall(i int) string {
	fake.untrackBlobMutex.RLock()
	defer fake.untrackBlobMutex.RUnlock()
	return fake.untrackBlobArgsForCall[i].arg1
}

func (fake *FakeBlobsDir) UntrackBlobReturns(result1 error) {
	fake.UntrackBlobStub = nil
	fake.untrackBlobReturns = struct {
		result1 error
	}{result1}
}

var _ releasedir.BlobsDir = new(FakeBlobsDir)
