// Code generated by counterfeiter. DO NOT EDIT.
package monitorfakes

import (
	"dnsmonitor/config"
	"dnsmonitor/pkg/model"
	"dnsmonitor/pkg/monitor"
	"sync"
)

type FakeMonitor struct {
	CheckStub        func() model.Record
	checkMutex       sync.RWMutex
	checkArgsForCall []struct {
	}
	checkReturns struct {
		result1 model.Record
	}
	checkReturnsOnCall map[int]struct {
		result1 model.Record
	}
	ConfigStub        func() config.Config
	configMutex       sync.RWMutex
	configArgsForCall []struct {
	}
	configReturns struct {
		result1 config.Config
	}
	configReturnsOnCall map[int]struct {
		result1 config.Config
	}
	DomainStub        func() *model.Domain
	domainMutex       sync.RWMutex
	domainArgsForCall []struct {
	}
	domainReturns struct {
		result1 *model.Domain
	}
	domainReturnsOnCall map[int]struct {
		result1 *model.Domain
	}
	ObserveStub        func() model.Record
	observeMutex       sync.RWMutex
	observeArgsForCall []struct {
	}
	observeReturns struct {
		result1 model.Record
	}
	observeReturnsOnCall map[int]struct {
		result1 model.Record
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeMonitor) Check() model.Record {
	fake.checkMutex.Lock()
	ret, specificReturn := fake.checkReturnsOnCall[len(fake.checkArgsForCall)]
	fake.checkArgsForCall = append(fake.checkArgsForCall, struct {
	}{})
	fake.recordInvocation("Check", []interface{}{})
	fake.checkMutex.Unlock()
	if fake.CheckStub != nil {
		return fake.CheckStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.checkReturns
	return fakeReturns.result1
}

func (fake *FakeMonitor) CheckCallCount() int {
	fake.checkMutex.RLock()
	defer fake.checkMutex.RUnlock()
	return len(fake.checkArgsForCall)
}

func (fake *FakeMonitor) CheckCalls(stub func() model.Record) {
	fake.checkMutex.Lock()
	defer fake.checkMutex.Unlock()
	fake.CheckStub = stub
}

func (fake *FakeMonitor) CheckReturns(result1 model.Record) {
	fake.checkMutex.Lock()
	defer fake.checkMutex.Unlock()
	fake.CheckStub = nil
	fake.checkReturns = struct {
		result1 model.Record
	}{result1}
}

func (fake *FakeMonitor) CheckReturnsOnCall(i int, result1 model.Record) {
	fake.checkMutex.Lock()
	defer fake.checkMutex.Unlock()
	fake.CheckStub = nil
	if fake.checkReturnsOnCall == nil {
		fake.checkReturnsOnCall = make(map[int]struct {
			result1 model.Record
		})
	}
	fake.checkReturnsOnCall[i] = struct {
		result1 model.Record
	}{result1}
}

func (fake *FakeMonitor) Config() config.Config {
	fake.configMutex.Lock()
	ret, specificReturn := fake.configReturnsOnCall[len(fake.configArgsForCall)]
	fake.configArgsForCall = append(fake.configArgsForCall, struct {
	}{})
	fake.recordInvocation("Config", []interface{}{})
	fake.configMutex.Unlock()
	if fake.ConfigStub != nil {
		return fake.ConfigStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.configReturns
	return fakeReturns.result1
}

func (fake *FakeMonitor) ConfigCallCount() int {
	fake.configMutex.RLock()
	defer fake.configMutex.RUnlock()
	return len(fake.configArgsForCall)
}

func (fake *FakeMonitor) ConfigCalls(stub func() config.Config) {
	fake.configMutex.Lock()
	defer fake.configMutex.Unlock()
	fake.ConfigStub = stub
}

func (fake *FakeMonitor) ConfigReturns(result1 config.Config) {
	fake.configMutex.Lock()
	defer fake.configMutex.Unlock()
	fake.ConfigStub = nil
	fake.configReturns = struct {
		result1 config.Config
	}{result1}
}

func (fake *FakeMonitor) ConfigReturnsOnCall(i int, result1 config.Config) {
	fake.configMutex.Lock()
	defer fake.configMutex.Unlock()
	fake.ConfigStub = nil
	if fake.configReturnsOnCall == nil {
		fake.configReturnsOnCall = make(map[int]struct {
			result1 config.Config
		})
	}
	fake.configReturnsOnCall[i] = struct {
		result1 config.Config
	}{result1}
}

func (fake *FakeMonitor) Domain() *model.Domain {
	fake.domainMutex.Lock()
	ret, specificReturn := fake.domainReturnsOnCall[len(fake.domainArgsForCall)]
	fake.domainArgsForCall = append(fake.domainArgsForCall, struct {
	}{})
	fake.recordInvocation("Domain", []interface{}{})
	fake.domainMutex.Unlock()
	if fake.DomainStub != nil {
		return fake.DomainStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.domainReturns
	return fakeReturns.result1
}

func (fake *FakeMonitor) DomainCallCount() int {
	fake.domainMutex.RLock()
	defer fake.domainMutex.RUnlock()
	return len(fake.domainArgsForCall)
}

func (fake *FakeMonitor) DomainCalls(stub func() *model.Domain) {
	fake.domainMutex.Lock()
	defer fake.domainMutex.Unlock()
	fake.DomainStub = stub
}

func (fake *FakeMonitor) DomainReturns(result1 *model.Domain) {
	fake.domainMutex.Lock()
	defer fake.domainMutex.Unlock()
	fake.DomainStub = nil
	fake.domainReturns = struct {
		result1 *model.Domain
	}{result1}
}

func (fake *FakeMonitor) DomainReturnsOnCall(i int, result1 *model.Domain) {
	fake.domainMutex.Lock()
	defer fake.domainMutex.Unlock()
	fake.DomainStub = nil
	if fake.domainReturnsOnCall == nil {
		fake.domainReturnsOnCall = make(map[int]struct {
			result1 *model.Domain
		})
	}
	fake.domainReturnsOnCall[i] = struct {
		result1 *model.Domain
	}{result1}
}

func (fake *FakeMonitor) Observe() model.Record {
	fake.observeMutex.Lock()
	ret, specificReturn := fake.observeReturnsOnCall[len(fake.observeArgsForCall)]
	fake.observeArgsForCall = append(fake.observeArgsForCall, struct {
	}{})
	fake.recordInvocation("Observe", []interface{}{})
	fake.observeMutex.Unlock()
	if fake.ObserveStub != nil {
		return fake.ObserveStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.observeReturns
	return fakeReturns.result1
}

func (fake *FakeMonitor) ObserveCallCount() int {
	fake.observeMutex.RLock()
	defer fake.observeMutex.RUnlock()
	return len(fake.observeArgsForCall)
}

func (fake *FakeMonitor) ObserveCalls(stub func() model.Record) {
	fake.observeMutex.Lock()
	defer fake.observeMutex.Unlock()
	fake.ObserveStub = stub
}

func (fake *FakeMonitor) ObserveReturns(result1 model.Record) {
	fake.observeMutex.Lock()
	defer fake.observeMutex.Unlock()
	fake.ObserveStub = nil
	fake.observeReturns = struct {
		result1 model.Record
	}{result1}
}

func (fake *FakeMonitor) ObserveReturnsOnCall(i int, result1 model.Record) {
	fake.observeMutex.Lock()
	defer fake.observeMutex.Unlock()
	fake.ObserveStub = nil
	if fake.observeReturnsOnCall == nil {
		fake.observeReturnsOnCall = make(map[int]struct {
			result1 model.Record
		})
	}
	fake.observeReturnsOnCall[i] = struct {
		result1 model.Record
	}{result1}
}

func (fake *FakeMonitor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.checkMutex.RLock()
	defer fake.checkMutex.RUnlock()
	fake.configMutex.RLock()
	defer fake.configMutex.RUnlock()
	fake.domainMutex.RLock()
	defer fake.domainMutex.RUnlock()
	fake.observeMutex.RLock()
	defer fake.observeMutex.RUnlock()
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
