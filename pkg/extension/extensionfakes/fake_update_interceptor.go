// Code generated by counterfeiter. DO NOT EDIT.
package extensionfakes

import (
	"sync"

	"github.com/Peripli/service-manager/pkg/extension"
)

type FakeUpdateInterceptor struct {
	OnAPIUpdateStub        func(h extension.InterceptUpdateOnAPI) extension.InterceptUpdateOnAPI
	onAPIUpdateMutex       sync.RWMutex
	onAPIUpdateArgsForCall []struct {
		h extension.InterceptUpdateOnAPI
	}
	onAPIUpdateReturns struct {
		result1 extension.InterceptUpdateOnAPI
	}
	onAPIUpdateReturnsOnCall map[int]struct {
		result1 extension.InterceptUpdateOnAPI
	}
	OnTransactionUpdateStub        func(f extension.InterceptUpdateOnTx) extension.InterceptUpdateOnTx
	onTransactionUpdateMutex       sync.RWMutex
	onTransactionUpdateArgsForCall []struct {
		f extension.InterceptUpdateOnTx
	}
	onTransactionUpdateReturns struct {
		result1 extension.InterceptUpdateOnTx
	}
	onTransactionUpdateReturnsOnCall map[int]struct {
		result1 extension.InterceptUpdateOnTx
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeUpdateInterceptor) OnAPIUpdate(h extension.InterceptUpdateOnAPI) extension.InterceptUpdateOnAPI {
	fake.onAPIUpdateMutex.Lock()
	ret, specificReturn := fake.onAPIUpdateReturnsOnCall[len(fake.onAPIUpdateArgsForCall)]
	fake.onAPIUpdateArgsForCall = append(fake.onAPIUpdateArgsForCall, struct {
		h extension.InterceptUpdateOnAPI
	}{h})
	fake.recordInvocation("OnAPIUpdate", []interface{}{h})
	fake.onAPIUpdateMutex.Unlock()
	if fake.OnAPIUpdateStub != nil {
		return fake.OnAPIUpdateStub(h)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.onAPIUpdateReturns.result1
}

func (fake *FakeUpdateInterceptor) OnAPIUpdateCallCount() int {
	fake.onAPIUpdateMutex.RLock()
	defer fake.onAPIUpdateMutex.RUnlock()
	return len(fake.onAPIUpdateArgsForCall)
}

func (fake *FakeUpdateInterceptor) OnAPIUpdateArgsForCall(i int) extension.InterceptUpdateOnAPI {
	fake.onAPIUpdateMutex.RLock()
	defer fake.onAPIUpdateMutex.RUnlock()
	return fake.onAPIUpdateArgsForCall[i].h
}

func (fake *FakeUpdateInterceptor) OnAPIUpdateReturns(result1 extension.InterceptUpdateOnAPI) {
	fake.OnAPIUpdateStub = nil
	fake.onAPIUpdateReturns = struct {
		result1 extension.InterceptUpdateOnAPI
	}{result1}
}

func (fake *FakeUpdateInterceptor) OnAPIUpdateReturnsOnCall(i int, result1 extension.InterceptUpdateOnAPI) {
	fake.OnAPIUpdateStub = nil
	if fake.onAPIUpdateReturnsOnCall == nil {
		fake.onAPIUpdateReturnsOnCall = make(map[int]struct {
			result1 extension.InterceptUpdateOnAPI
		})
	}
	fake.onAPIUpdateReturnsOnCall[i] = struct {
		result1 extension.InterceptUpdateOnAPI
	}{result1}
}

func (fake *FakeUpdateInterceptor) OnTxUpdate(f extension.InterceptUpdateOnTx) extension.InterceptUpdateOnTx {
	fake.onTransactionUpdateMutex.Lock()
	ret, specificReturn := fake.onTransactionUpdateReturnsOnCall[len(fake.onTransactionUpdateArgsForCall)]
	fake.onTransactionUpdateArgsForCall = append(fake.onTransactionUpdateArgsForCall, struct {
		f extension.InterceptUpdateOnTx
	}{f})
	fake.recordInvocation("OnTransactionUpdate", []interface{}{f})
	fake.onTransactionUpdateMutex.Unlock()
	if fake.OnTransactionUpdateStub != nil {
		return fake.OnTransactionUpdateStub(f)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.onTransactionUpdateReturns.result1
}

func (fake *FakeUpdateInterceptor) OnTransactionUpdateCallCount() int {
	fake.onTransactionUpdateMutex.RLock()
	defer fake.onTransactionUpdateMutex.RUnlock()
	return len(fake.onTransactionUpdateArgsForCall)
}

func (fake *FakeUpdateInterceptor) OnTransactionUpdateArgsForCall(i int) extension.InterceptUpdateOnTx {
	fake.onTransactionUpdateMutex.RLock()
	defer fake.onTransactionUpdateMutex.RUnlock()
	return fake.onTransactionUpdateArgsForCall[i].f
}

func (fake *FakeUpdateInterceptor) OnTransactionUpdateReturns(result1 extension.InterceptUpdateOnTx) {
	fake.OnTransactionUpdateStub = nil
	fake.onTransactionUpdateReturns = struct {
		result1 extension.InterceptUpdateOnTx
	}{result1}
}

func (fake *FakeUpdateInterceptor) OnTransactionUpdateReturnsOnCall(i int, result1 extension.InterceptUpdateOnTx) {
	fake.OnTransactionUpdateStub = nil
	if fake.onTransactionUpdateReturnsOnCall == nil {
		fake.onTransactionUpdateReturnsOnCall = make(map[int]struct {
			result1 extension.InterceptUpdateOnTx
		})
	}
	fake.onTransactionUpdateReturnsOnCall[i] = struct {
		result1 extension.InterceptUpdateOnTx
	}{result1}
}

func (fake *FakeUpdateInterceptor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.onAPIUpdateMutex.RLock()
	defer fake.onAPIUpdateMutex.RUnlock()
	fake.onTransactionUpdateMutex.RLock()
	defer fake.onTransactionUpdateMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeUpdateInterceptor) recordInvocation(key string, args []interface{}) {
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

var _ extension.UpdateInterceptor = new(FakeUpdateInterceptor)
