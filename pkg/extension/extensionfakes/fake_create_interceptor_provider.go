// Code generated by counterfeiter. DO NOT EDIT.
package extensionfakes

import (
	"sync"

	"github.com/Peripli/service-manager/pkg/extension"
)

type FakeCreateInterceptorProvider struct {
	NameStub        func() string
	nameMutex       sync.RWMutex
	nameArgsForCall []struct{}
	nameReturns     struct {
		result1 string
	}
	nameReturnsOnCall map[int]struct {
		result1 string
	}
	ProvideStub        func() extension.CreateInterceptor
	provideMutex       sync.RWMutex
	provideArgsForCall []struct{}
	provideReturns     struct {
		result1 extension.CreateInterceptor
	}
	provideReturnsOnCall map[int]struct {
		result1 extension.CreateInterceptor
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCreateInterceptorProvider) Name() string {
	fake.nameMutex.Lock()
	ret, specificReturn := fake.nameReturnsOnCall[len(fake.nameArgsForCall)]
	fake.nameArgsForCall = append(fake.nameArgsForCall, struct{}{})
	fake.recordInvocation("Name", []interface{}{})
	fake.nameMutex.Unlock()
	if fake.NameStub != nil {
		return fake.NameStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.nameReturns.result1
}

func (fake *FakeCreateInterceptorProvider) NameCallCount() int {
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	return len(fake.nameArgsForCall)
}

func (fake *FakeCreateInterceptorProvider) NameReturns(result1 string) {
	fake.NameStub = nil
	fake.nameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeCreateInterceptorProvider) NameReturnsOnCall(i int, result1 string) {
	fake.NameStub = nil
	if fake.nameReturnsOnCall == nil {
		fake.nameReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.nameReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeCreateInterceptorProvider) Provide() extension.CreateInterceptor {
	fake.provideMutex.Lock()
	ret, specificReturn := fake.provideReturnsOnCall[len(fake.provideArgsForCall)]
	fake.provideArgsForCall = append(fake.provideArgsForCall, struct{}{})
	fake.recordInvocation("Provide", []interface{}{})
	fake.provideMutex.Unlock()
	if fake.ProvideStub != nil {
		return fake.ProvideStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.provideReturns.result1
}

func (fake *FakeCreateInterceptorProvider) ProvideCallCount() int {
	fake.provideMutex.RLock()
	defer fake.provideMutex.RUnlock()
	return len(fake.provideArgsForCall)
}

func (fake *FakeCreateInterceptorProvider) ProvideReturns(result1 extension.CreateInterceptor) {
	fake.ProvideStub = nil
	fake.provideReturns = struct {
		result1 extension.CreateInterceptor
	}{result1}
}

func (fake *FakeCreateInterceptorProvider) ProvideReturnsOnCall(i int, result1 extension.CreateInterceptor) {
	fake.ProvideStub = nil
	if fake.provideReturnsOnCall == nil {
		fake.provideReturnsOnCall = make(map[int]struct {
			result1 extension.CreateInterceptor
		})
	}
	fake.provideReturnsOnCall[i] = struct {
		result1 extension.CreateInterceptor
	}{result1}
}

func (fake *FakeCreateInterceptorProvider) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	fake.provideMutex.RLock()
	defer fake.provideMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCreateInterceptorProvider) recordInvocation(key string, args []interface{}) {
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

var _ extension.CreateInterceptorProvider = new(FakeCreateInterceptorProvider)
