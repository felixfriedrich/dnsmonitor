// Code generated by counterfeiter. DO NOT EDIT.
package monitorfakes

import (
	"dnsmonitor/pkg/configuration"
	"dnsmonitor/pkg/model"
	"dnsmonitor/pkg/monitor"
	"sync"
)

type FakeMonitor struct {
	CheckStub        func()
	checkMutex       sync.RWMutex
	checkArgsForCall []struct {
	}
	ConfigStub        func() configuration.Monitor
	configMutex       sync.RWMutex
	configArgsForCall []struct {
	}
	configReturns struct {
		result1 configuration.Monitor
	}
	configReturnsOnCall map[int]struct {
		result1 configuration.Monitor
	}
	DomainsStub        func() []*model.Domain
	domainsMutex       sync.RWMutex
	domainsArgsForCall []struct {
	}
	domainsReturns struct {
		result1 []*model.Domain
	}
	domainsReturnsOnCall map[int]struct {
		result1 []*model.Domain
	}
	ObserveStub        func()
	observeMutex       sync.RWMutex
	observeArgsForCall []struct {
	}
	RunStub        func(int, bool)
	runMutex       sync.RWMutex
	runArgsForCall []struct {
		arg1 int
		arg2 bool
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeMonitor) Check() {
	fake.checkMutex.Lock()
	fake.checkArgsForCall = append(fake.checkArgsForCall, struct {
	}{})
	stub := fake.CheckStub
	fake.recordInvocation("Check", []interface{}{})
	fake.checkMutex.Unlock()
	if stub != nil {
		fake.CheckStub()
	}
}

func (fake *FakeMonitor) CheckCallCount() int {
	fake.checkMutex.RLock()
	defer fake.checkMutex.RUnlock()
	return len(fake.checkArgsForCall)
}

func (fake *FakeMonitor) CheckCalls(stub func()) {
	fake.checkMutex.Lock()
	defer fake.checkMutex.Unlock()
	fake.CheckStub = stub
}

func (fake *FakeMonitor) Config() configuration.Monitor {
	fake.configMutex.Lock()
	ret, specificReturn := fake.configReturnsOnCall[len(fake.configArgsForCall)]
	fake.configArgsForCall = append(fake.configArgsForCall, struct {
	}{})
	stub := fake.ConfigStub
	fakeReturns := fake.configReturns
	fake.recordInvocation("Config", []interface{}{})
	fake.configMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeMonitor) ConfigCallCount() int {
	fake.configMutex.RLock()
	defer fake.configMutex.RUnlock()
	return len(fake.configArgsForCall)
}

func (fake *FakeMonitor) ConfigCalls(stub func() configuration.Monitor) {
	fake.configMutex.Lock()
	defer fake.configMutex.Unlock()
	fake.ConfigStub = stub
}

func (fake *FakeMonitor) ConfigReturns(result1 configuration.Monitor) {
	fake.configMutex.Lock()
	defer fake.configMutex.Unlock()
	fake.ConfigStub = nil
	fake.configReturns = struct {
		result1 configuration.Monitor
	}{result1}
}

func (fake *FakeMonitor) ConfigReturnsOnCall(i int, result1 configuration.Monitor) {
	fake.configMutex.Lock()
	defer fake.configMutex.Unlock()
	fake.ConfigStub = nil
	if fake.configReturnsOnCall == nil {
		fake.configReturnsOnCall = make(map[int]struct {
			result1 configuration.Monitor
		})
	}
	fake.configReturnsOnCall[i] = struct {
		result1 configuration.Monitor
	}{result1}
}

func (fake *FakeMonitor) Domains() []*model.Domain {
	fake.domainsMutex.Lock()
	ret, specificReturn := fake.domainsReturnsOnCall[len(fake.domainsArgsForCall)]
	fake.domainsArgsForCall = append(fake.domainsArgsForCall, struct {
	}{})
	stub := fake.DomainsStub
	fakeReturns := fake.domainsReturns
	fake.recordInvocation("Domains", []interface{}{})
	fake.domainsMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeMonitor) DomainsCallCount() int {
	fake.domainsMutex.RLock()
	defer fake.domainsMutex.RUnlock()
	return len(fake.domainsArgsForCall)
}

func (fake *FakeMonitor) DomainsCalls(stub func() []*model.Domain) {
	fake.domainsMutex.Lock()
	defer fake.domainsMutex.Unlock()
	fake.DomainsStub = stub
}

func (fake *FakeMonitor) DomainsReturns(result1 []*model.Domain) {
	fake.domainsMutex.Lock()
	defer fake.domainsMutex.Unlock()
	fake.DomainsStub = nil
	fake.domainsReturns = struct {
		result1 []*model.Domain
	}{result1}
}

func (fake *FakeMonitor) DomainsReturnsOnCall(i int, result1 []*model.Domain) {
	fake.domainsMutex.Lock()
	defer fake.domainsMutex.Unlock()
	fake.DomainsStub = nil
	if fake.domainsReturnsOnCall == nil {
		fake.domainsReturnsOnCall = make(map[int]struct {
			result1 []*model.Domain
		})
	}
	fake.domainsReturnsOnCall[i] = struct {
		result1 []*model.Domain
	}{result1}
}

func (fake *FakeMonitor) Observe() {
	fake.observeMutex.Lock()
	fake.observeArgsForCall = append(fake.observeArgsForCall, struct {
	}{})
	stub := fake.ObserveStub
	fake.recordInvocation("Observe", []interface{}{})
	fake.observeMutex.Unlock()
	if stub != nil {
		fake.ObserveStub()
	}
}

func (fake *FakeMonitor) ObserveCallCount() int {
	fake.observeMutex.RLock()
	defer fake.observeMutex.RUnlock()
	return len(fake.observeArgsForCall)
}

func (fake *FakeMonitor) ObserveCalls(stub func()) {
	fake.observeMutex.Lock()
	defer fake.observeMutex.Unlock()
	fake.ObserveStub = stub
}

func (fake *FakeMonitor) Run(arg1 int, arg2 bool) {
	fake.runMutex.Lock()
	fake.runArgsForCall = append(fake.runArgsForCall, struct {
		arg1 int
		arg2 bool
	}{arg1, arg2})
	stub := fake.RunStub
	fake.recordInvocation("Run", []interface{}{arg1, arg2})
	fake.runMutex.Unlock()
	if stub != nil {
		fake.RunStub(arg1, arg2)
	}
}

func (fake *FakeMonitor) RunCallCount() int {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return len(fake.runArgsForCall)
}

func (fake *FakeMonitor) RunCalls(stub func(int, bool)) {
	fake.runMutex.Lock()
	defer fake.runMutex.Unlock()
	fake.RunStub = stub
}

func (fake *FakeMonitor) RunArgsForCall(i int) (int, bool) {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	argsForCall := fake.runArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeMonitor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.checkMutex.RLock()
	defer fake.checkMutex.RUnlock()
	fake.configMutex.RLock()
	defer fake.configMutex.RUnlock()
	fake.domainsMutex.RLock()
	defer fake.domainsMutex.RUnlock()
	fake.observeMutex.RLock()
	defer fake.observeMutex.RUnlock()
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeMonitor) recordInvocation(key string, args []interface{}) {
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

var _ monitor.Monitor = new(FakeMonitor)
