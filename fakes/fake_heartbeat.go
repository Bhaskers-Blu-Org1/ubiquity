// This file was generated by counterfeiter
package fakes

import (
	"sync"
	"time"

	"github.com/ibm/ubiquity/utils"
)

type FakeHeartbeat struct {
	ExistsStub        func() (bool, error)
	existsMutex       sync.RWMutex
	existsArgsForCall []struct{}
	existsReturns     struct {
		result1 bool
		result2 error
	}
	CreateStub        func() error
	createMutex       sync.RWMutex
	createArgsForCall []struct{}
	createReturns     struct {
		result1 error
	}
	UpdateStub        func() error
	updateMutex       sync.RWMutex
	updateArgsForCall []struct{}
	updateReturns     struct {
		result1 error
	}
	GetLastUpdateTimestampStub        func() (time.Time, error)
	getLastUpdateTimestampMutex       sync.RWMutex
	getLastUpdateTimestampArgsForCall []struct{}
	getLastUpdateTimestampReturns     struct {
		result1 time.Time
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeHeartbeat) Exists() (bool, error) {
	fake.existsMutex.Lock()
	fake.existsArgsForCall = append(fake.existsArgsForCall, struct{}{})
	fake.recordInvocation("Exists", []interface{}{})
	fake.existsMutex.Unlock()
	if fake.ExistsStub != nil {
		return fake.ExistsStub()
	}
	return fake.existsReturns.result1, fake.existsReturns.result2
}

func (fake *FakeHeartbeat) ExistsCallCount() int {
	fake.existsMutex.RLock()
	defer fake.existsMutex.RUnlock()
	return len(fake.existsArgsForCall)
}

func (fake *FakeHeartbeat) ExistsReturns(result1 bool, result2 error) {
	fake.ExistsStub = nil
	fake.existsReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeHeartbeat) Create() error {
	fake.createMutex.Lock()
	fake.createArgsForCall = append(fake.createArgsForCall, struct{}{})
	fake.recordInvocation("Create", []interface{}{})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub()
	}
	return fake.createReturns.result1
}

func (fake *FakeHeartbeat) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeHeartbeat) CreateReturns(result1 error) {
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeHeartbeat) Update() error {
	fake.updateMutex.Lock()
	fake.updateArgsForCall = append(fake.updateArgsForCall, struct{}{})
	fake.recordInvocation("Update", []interface{}{})
	fake.updateMutex.Unlock()
	if fake.UpdateStub != nil {
		return fake.UpdateStub()
	}
	return fake.updateReturns.result1
}

func (fake *FakeHeartbeat) UpdateCallCount() int {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return len(fake.updateArgsForCall)
}

func (fake *FakeHeartbeat) UpdateReturns(result1 error) {
	fake.UpdateStub = nil
	fake.updateReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeHeartbeat) GetLastUpdateTimestamp() (time.Time, error) {
	fake.getLastUpdateTimestampMutex.Lock()
	fake.getLastUpdateTimestampArgsForCall = append(fake.getLastUpdateTimestampArgsForCall, struct{}{})
	fake.recordInvocation("GetLastUpdateTimestamp", []interface{}{})
	fake.getLastUpdateTimestampMutex.Unlock()
	if fake.GetLastUpdateTimestampStub != nil {
		return fake.GetLastUpdateTimestampStub()
	}
	return fake.getLastUpdateTimestampReturns.result1, fake.getLastUpdateTimestampReturns.result2
}

func (fake *FakeHeartbeat) GetLastUpdateTimestampCallCount() int {
	fake.getLastUpdateTimestampMutex.RLock()
	defer fake.getLastUpdateTimestampMutex.RUnlock()
	return len(fake.getLastUpdateTimestampArgsForCall)
}

func (fake *FakeHeartbeat) GetLastUpdateTimestampReturns(result1 time.Time, result2 error) {
	fake.GetLastUpdateTimestampStub = nil
	fake.getLastUpdateTimestampReturns = struct {
		result1 time.Time
		result2 error
	}{result1, result2}
}

func (fake *FakeHeartbeat) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.existsMutex.RLock()
	defer fake.existsMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	fake.getLastUpdateTimestampMutex.RLock()
	defer fake.getLastUpdateTimestampMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeHeartbeat) recordInvocation(key string, args []interface{}) {
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

var _ utils.Heartbeat = new(FakeHeartbeat)
