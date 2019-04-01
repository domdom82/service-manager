// Code generated by counterfeiter. DO NOT EDIT.
package extensionfakes

import (
	"sync"

	"github.com/Peripli/service-manager/pkg/extension"
)

type FakeDeleteInterceptor struct {
	OnAPIDeleteStub        func(h extension.InterceptDeleteOnAPI) extension.InterceptDeleteOnAPI
	onAPIDeleteMutex       sync.RWMutex
	onAPIDeleteArgsForCall []struct {
		h extension.InterceptDeleteOnAPI
	}
	onAPIDeleteReturns struct {
		result1 extension.InterceptDeleteOnAPI
	}
	onAPIDeleteReturnsOnCall map[int]struct {
		result1 extension.InterceptDeleteOnAPI
	}
	OnTransactionDeleteStub        func(f extension.InterceptDeleteOnTransaction) extension.InterceptDeleteOnTransaction
	onTransactionDeleteMutex       sync.RWMutex
	onTransactionDeleteArgsForCall []struct {
		f extension.InterceptDeleteOnTransaction
	}
	onTransactionDeleteReturns struct {
		result1 extension.InterceptDeleteOnTransaction
	}
	onTransactionDeleteReturnsOnCall map[int]struct {
		result1 extension.InterceptDeleteOnTransaction
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDeleteInterceptor) OnAPIDelete(h extension.InterceptDeleteOnAPI) extension.InterceptDeleteOnAPI {
	fake.onAPIDeleteMutex.Lock()
	ret, specificReturn := fake.onAPIDeleteReturnsOnCall[len(fake.onAPIDeleteArgsForCall)]
	fake.onAPIDeleteArgsForCall = append(fake.onAPIDeleteArgsForCall, struct {
		h extension.InterceptDeleteOnAPI
	}{h})
	fake.recordInvocation("OnAPIDelete", []interface{}{h})
	fake.onAPIDeleteMutex.Unlock()
	if fake.OnAPIDeleteStub != nil {
		return fake.OnAPIDeleteStub(h)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.onAPIDeleteReturns.result1
}

func (fake *FakeDeleteInterceptor) OnAPIDeleteCallCount() int {
	fake.onAPIDeleteMutex.RLock()
	defer fake.onAPIDeleteMutex.RUnlock()
	return len(fake.onAPIDeleteArgsForCall)
}

func (fake *FakeDeleteInterceptor) OnAPIDeleteArgsForCall(i int) extension.InterceptDeleteOnAPI {
	fake.onAPIDeleteMutex.RLock()
	defer fake.onAPIDeleteMutex.RUnlock()
	return fake.onAPIDeleteArgsForCall[i].h
}

func (fake *FakeDeleteInterceptor) OnAPIDeleteReturns(result1 extension.InterceptDeleteOnAPI) {
	fake.OnAPIDeleteStub = nil
	fake.onAPIDeleteReturns = struct {
		result1 extension.InterceptDeleteOnAPI
	}{result1}
}

func (fake *FakeDeleteInterceptor) OnAPIDeleteReturnsOnCall(i int, result1 extension.InterceptDeleteOnAPI) {
	fake.OnAPIDeleteStub = nil
	if fake.onAPIDeleteReturnsOnCall == nil {
		fake.onAPIDeleteReturnsOnCall = make(map[int]struct {
			result1 extension.InterceptDeleteOnAPI
		})
	}
	fake.onAPIDeleteReturnsOnCall[i] = struct {
		result1 extension.InterceptDeleteOnAPI
	}{result1}
}

func (fake *FakeDeleteInterceptor) OnTransactionDelete(f extension.InterceptDeleteOnTransaction) extension.InterceptDeleteOnTransaction {
	fake.onTransactionDeleteMutex.Lock()
	ret, specificReturn := fake.onTransactionDeleteReturnsOnCall[len(fake.onTransactionDeleteArgsForCall)]
	fake.onTransactionDeleteArgsForCall = append(fake.onTransactionDeleteArgsForCall, struct {
		f extension.InterceptDeleteOnTransaction
	}{f})
	fake.recordInvocation("OnTransactionDelete", []interface{}{f})
	fake.onTransactionDeleteMutex.Unlock()
	if fake.OnTransactionDeleteStub != nil {
		return fake.OnTransactionDeleteStub(f)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.onTransactionDeleteReturns.result1
}

func (fake *FakeDeleteInterceptor) OnTransactionDeleteCallCount() int {
	fake.onTransactionDeleteMutex.RLock()
	defer fake.onTransactionDeleteMutex.RUnlock()
	return len(fake.onTransactionDeleteArgsForCall)
}

func (fake *FakeDeleteInterceptor) OnTransactionDeleteArgsForCall(i int) extension.InterceptDeleteOnTransaction {
	fake.onTransactionDeleteMutex.RLock()
	defer fake.onTransactionDeleteMutex.RUnlock()
	return fake.onTransactionDeleteArgsForCall[i].f
}

func (fake *FakeDeleteInterceptor) OnTransactionDeleteReturns(result1 extension.InterceptDeleteOnTransaction) {
	fake.OnTransactionDeleteStub = nil
	fake.onTransactionDeleteReturns = struct {
		result1 extension.InterceptDeleteOnTransaction
	}{result1}
}

func (fake *FakeDeleteInterceptor) OnTransactionDeleteReturnsOnCall(i int, result1 extension.InterceptDeleteOnTransaction) {
	fake.OnTransactionDeleteStub = nil
	if fake.onTransactionDeleteReturnsOnCall == nil {
		fake.onTransactionDeleteReturnsOnCall = make(map[int]struct {
			result1 extension.InterceptDeleteOnTransaction
		})
	}
	fake.onTransactionDeleteReturnsOnCall[i] = struct {
		result1 extension.InterceptDeleteOnTransaction
	}{result1}
}

func (fake *FakeDeleteInterceptor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.onAPIDeleteMutex.RLock()
	defer fake.onAPIDeleteMutex.RUnlock()
	fake.onTransactionDeleteMutex.RLock()
	defer fake.onTransactionDeleteMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeDeleteInterceptor) recordInvocation(key string, args []interface{}) {
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

var _ extension.DeleteInterceptor = new(FakeDeleteInterceptor)
