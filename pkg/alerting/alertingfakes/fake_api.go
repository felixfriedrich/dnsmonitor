// Code generated by counterfeiter. DO NOT EDIT.
package alertingfakes

import (
	"dnsmonitor/pkg/alerting"
	"sync"
)

type FakeAPI struct {
	SendSMSStub        func(string) error
	sendSMSMutex       sync.RWMutex
	sendSMSArgsForCall []struct {
		arg1 string
	}
	sendSMSReturns struct {
		result1 error
	}
	sendSMSReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeAPI) SendSMS(arg1 string) error {
	fake.sendSMSMutex.Lock()
	ret, specificReturn := fake.sendSMSReturnsOnCall[len(fake.sendSMSArgsForCall)]
	fake.sendSMSArgsForCall = append(fake.sendSMSArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.SendSMSStub
	fakeReturns := fake.sendSMSReturns
	fake.recordInvocation("SendSMS", []interface{}{arg1})
	fake.sendSMSMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeAPI) SendSMSCallCount() int {
	fake.sendSMSMutex.RLock()
	defer fake.sendSMSMutex.RUnlock()
	return len(fake.sendSMSArgsForCall)
}

func (fake *FakeAPI) SendSMSCalls(stub func(string) error) {
	fake.sendSMSMutex.Lock()
	defer fake.sendSMSMutex.Unlock()
	fake.SendSMSStub = stub
}

func (fake *FakeAPI) SendSMSArgsForCall(i int) string {
	fake.sendSMSMutex.RLock()
	defer fake.sendSMSMutex.RUnlock()
	argsForCall := fake.sendSMSArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeAPI) SendSMSReturns(result1 error) {
	fake.sendSMSMutex.Lock()
	defer fake.sendSMSMutex.Unlock()
	fake.SendSMSStub = nil
	fake.sendSMSReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAPI) SendSMSReturnsOnCall(i int, result1 error) {
	fake.sendSMSMutex.Lock()
	defer fake.sendSMSMutex.Unlock()
	fake.SendSMSStub = nil
	if fake.sendSMSReturnsOnCall == nil {
		fake.sendSMSReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.sendSMSReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeAPI) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.sendSMSMutex.RLock()
	defer fake.sendSMSMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeAPI) recordInvocation(key string, args []interface{}) {
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

var _ alerting.API = new(FakeAPI)
