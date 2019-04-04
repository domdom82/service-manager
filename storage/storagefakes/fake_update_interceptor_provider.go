// Code generated by counterfeiter. DO NOT EDIT.
package storagefakes

import (
	sync "sync"

	storage "github.com/Peripli/service-manager/storage"
)

type FakeUpdateInterceptorProvider struct {
	ProvideStub        func() storage.UpdateInterceptor
	provideMutex       sync.RWMutex
	provideArgsForCall []struct {
	}
	provideReturns struct {
		result1 storage.UpdateInterceptor
	}
	provideReturnsOnCall map[int]struct {
		result1 storage.UpdateInterceptor
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeUpdateInterceptorProvider) Provide() storage.UpdateInterceptor {
	fake.provideMutex.Lock()
	ret, specificReturn := fake.provideReturnsOnCall[len(fake.provideArgsForCall)]
	fake.provideArgsForCall = append(fake.provideArgsForCall, struct {
	}{})
	fake.recordInvocation("Provide", []interface{}{})
	fake.provideMutex.Unlock()
	if fake.ProvideStub != nil {
		return fake.ProvideStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.provideReturns
	return fakeReturns.result1
}

func (fake *FakeUpdateInterceptorProvider) ProvideCallCount() int {
	fake.provideMutex.RLock()
	defer fake.provideMutex.RUnlock()
	return len(fake.provideArgsForCall)
}

func (fake *FakeUpdateInterceptorProvider) ProvideCalls(stub func() storage.UpdateInterceptor) {
	fake.provideMutex.Lock()
	defer fake.provideMutex.Unlock()
	fake.ProvideStub = stub
}

func (fake *FakeUpdateInterceptorProvider) ProvideReturns(result1 storage.UpdateInterceptor) {
	fake.provideMutex.Lock()
	defer fake.provideMutex.Unlock()
	fake.ProvideStub = nil
	fake.provideReturns = struct {
		result1 storage.UpdateInterceptor
	}{result1}
}

func (fake *FakeUpdateInterceptorProvider) ProvideReturnsOnCall(i int, result1 storage.UpdateInterceptor) {
	fake.provideMutex.Lock()
	defer fake.provideMutex.Unlock()
	fake.ProvideStub = nil
	if fake.provideReturnsOnCall == nil {
		fake.provideReturnsOnCall = make(map[int]struct {
			result1 storage.UpdateInterceptor
		})
	}
	fake.provideReturnsOnCall[i] = struct {
		result1 storage.UpdateInterceptor
	}{result1}
}

func (fake *FakeUpdateInterceptorProvider) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.provideMutex.RLock()
	defer fake.provideMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeUpdateInterceptorProvider) recordInvocation(key string, args []interface{}) {
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

var _ storage.UpdateInterceptorProvider = new(FakeUpdateInterceptorProvider)