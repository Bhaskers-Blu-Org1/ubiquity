// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/IBM/ubiquity/remote/mounter/initiator"
)

type FakeBaseInitiator struct {
	FlushMultipathStub        func(string)
	flushMultipathMutex       sync.RWMutex
	flushMultipathArgsForCall []struct {
		arg1 string
	}
	RemoveSCSIDeviceStub        func(string) error
	removeSCSIDeviceMutex       sync.RWMutex
	removeSCSIDeviceArgsForCall []struct {
		arg1 string
	}
	removeSCSIDeviceReturns struct {
		result1 error
	}
	removeSCSIDeviceReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBaseInitiator) FlushMultipath(arg1 string) {
	fake.flushMultipathMutex.Lock()
	fake.flushMultipathArgsForCall = append(fake.flushMultipathArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("FlushMultipath", []interface{}{arg1})
	fake.flushMultipathMutex.Unlock()
	if fake.FlushMultipathStub != nil {
		fake.FlushMultipathStub(arg1)
	}
}

func (fake *FakeBaseInitiator) FlushMultipathCallCount() int {
	fake.flushMultipathMutex.RLock()
	defer fake.flushMultipathMutex.RUnlock()
	return len(fake.flushMultipathArgsForCall)
}

func (fake *FakeBaseInitiator) FlushMultipathCalls(stub func(string)) {
	fake.flushMultipathMutex.Lock()
	defer fake.flushMultipathMutex.Unlock()
	fake.FlushMultipathStub = stub
}

func (fake *FakeBaseInitiator) FlushMultipathArgsForCall(i int) string {
	fake.flushMultipathMutex.RLock()
	defer fake.flushMultipathMutex.RUnlock()
	argsForCall := fake.flushMultipathArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeBaseInitiator) RemoveSCSIDevice(arg1 string) error {
	fake.removeSCSIDeviceMutex.Lock()
	ret, specificReturn := fake.removeSCSIDeviceReturnsOnCall[len(fake.removeSCSIDeviceArgsForCall)]
	fake.removeSCSIDeviceArgsForCall = append(fake.removeSCSIDeviceArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("RemoveSCSIDevice", []interface{}{arg1})
	fake.removeSCSIDeviceMutex.Unlock()
	if fake.RemoveSCSIDeviceStub != nil {
		return fake.RemoveSCSIDeviceStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.removeSCSIDeviceReturns
	return fakeReturns.result1
}

func (fake *FakeBaseInitiator) RemoveSCSIDeviceCallCount() int {
	fake.removeSCSIDeviceMutex.RLock()
	defer fake.removeSCSIDeviceMutex.RUnlock()
	return len(fake.removeSCSIDeviceArgsForCall)
}

func (fake *FakeBaseInitiator) RemoveSCSIDeviceCalls(stub func(string) error) {
	fake.removeSCSIDeviceMutex.Lock()
	defer fake.removeSCSIDeviceMutex.Unlock()
	fake.RemoveSCSIDeviceStub = stub
}

func (fake *FakeBaseInitiator) RemoveSCSIDeviceArgsForCall(i int) string {
	fake.removeSCSIDeviceMutex.RLock()
	defer fake.removeSCSIDeviceMutex.RUnlock()
	argsForCall := fake.removeSCSIDeviceArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeBaseInitiator) RemoveSCSIDeviceReturns(result1 error) {
	fake.removeSCSIDeviceMutex.Lock()
	defer fake.removeSCSIDeviceMutex.Unlock()
	fake.RemoveSCSIDeviceStub = nil
	fake.removeSCSIDeviceReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBaseInitiator) RemoveSCSIDeviceReturnsOnCall(i int, result1 error) {
	fake.removeSCSIDeviceMutex.Lock()
	defer fake.removeSCSIDeviceMutex.Unlock()
	fake.RemoveSCSIDeviceStub = nil
	if fake.removeSCSIDeviceReturnsOnCall == nil {
		fake.removeSCSIDeviceReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.removeSCSIDeviceReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeBaseInitiator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.flushMultipathMutex.RLock()
	defer fake.flushMultipathMutex.RUnlock()
	fake.removeSCSIDeviceMutex.RLock()
	defer fake.removeSCSIDeviceMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeBaseInitiator) recordInvocation(key string, args []interface{}) {
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

var _ initiator.BaseInitiator = new(FakeBaseInitiator)