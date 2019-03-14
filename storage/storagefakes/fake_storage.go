// Code generated by counterfeiter. DO NOT EDIT.
package storagefakes

import (
	"context"
	"sync"

	"github.com/Peripli/service-manager/pkg/query"
	"github.com/Peripli/service-manager/pkg/types"
	"github.com/Peripli/service-manager/storage"
)

type FakeStorage struct {
	OpenStub        func(options *storage.Settings, scheme *storage.Scheme) error
	openMutex       sync.RWMutex
	openArgsForCall []struct {
		options *storage.Settings
		scheme  *storage.Scheme
	}
	openReturns struct {
		result1 error
	}
	openReturnsOnCall map[int]struct {
		result1 error
	}
	CloseStub        func() error
	closeMutex       sync.RWMutex
	closeArgsForCall []struct{}
	closeReturns     struct {
		result1 error
	}
	closeReturnsOnCall map[int]struct {
		result1 error
	}
	PingStub        func() error
	pingMutex       sync.RWMutex
	pingArgsForCall []struct{}
	pingReturns     struct {
		result1 error
	}
	pingReturnsOnCall map[int]struct {
		result1 error
	}
	CreateStub        func(ctx context.Context, obj types.Object) (string, error)
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		ctx context.Context
		obj types.Object
	}
	createReturns struct {
		result1 string
		result2 error
	}
	createReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	GetStub        func(ctx context.Context, id string, objectType types.ObjectType) (types.Object, error)
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		ctx        context.Context
		id         string
		objectType types.ObjectType
	}
	getReturns struct {
		result1 types.Object
		result2 error
	}
	getReturnsOnCall map[int]struct {
		result1 types.Object
		result2 error
	}
	ListStub        func(ctx context.Context, objectType types.ObjectType, criteria ...query.Criterion) (types.ObjectList, error)
	listMutex       sync.RWMutex
	listArgsForCall []struct {
		ctx        context.Context
		objectType types.ObjectType
		criteria   []query.Criterion
	}
	listReturns struct {
		result1 types.ObjectList
		result2 error
	}
	listReturnsOnCall map[int]struct {
		result1 types.ObjectList
		result2 error
	}
	DeleteStub        func(ctx context.Context, objectType types.ObjectType, criteria ...query.Criterion) (types.ObjectList, error)
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		ctx        context.Context
		objectType types.ObjectType
		criteria   []query.Criterion
	}
	deleteReturns struct {
		result1 types.ObjectList
		result2 error
	}
	deleteReturnsOnCall map[int]struct {
		result1 types.ObjectList
		result2 error
	}
	UpdateStub        func(ctx context.Context, obj types.Object, labelChanges ...*query.LabelChange) (types.Object, error)
	updateMutex       sync.RWMutex
	updateArgsForCall []struct {
		ctx          context.Context
		obj          types.Object
		labelChanges []*query.LabelChange
	}
	updateReturns struct {
		result1 types.Object
		result2 error
	}
	updateReturnsOnCall map[int]struct {
		result1 types.Object
		result2 error
	}
	ServiceOfferingStub        func() storage.ServiceOffering
	serviceOfferingMutex       sync.RWMutex
	serviceOfferingArgsForCall []struct{}
	serviceOfferingReturns     struct {
		result1 storage.ServiceOffering
	}
	serviceOfferingReturnsOnCall map[int]struct {
		result1 storage.ServiceOffering
	}
	CredentialsStub        func() storage.Credentials
	credentialsMutex       sync.RWMutex
	credentialsArgsForCall []struct{}
	credentialsReturns     struct {
		result1 storage.Credentials
	}
	credentialsReturnsOnCall map[int]struct {
		result1 storage.Credentials
	}
	SecurityStub        func() storage.Security
	securityMutex       sync.RWMutex
	securityArgsForCall []struct{}
	securityReturns     struct {
		result1 storage.Security
	}
	securityReturnsOnCall map[int]struct {
		result1 storage.Security
	}
	InTransactionStub        func(ctx context.Context, f func(ctx context.Context, storage storage.Warehouse) error) error
	inTransactionMutex       sync.RWMutex
	inTransactionArgsForCall []struct {
		ctx context.Context
		f   func(ctx context.Context, storage storage.Warehouse) error
	}
	inTransactionReturns struct {
		result1 error
	}
	inTransactionReturnsOnCall map[int]struct {
		result1 error
	}
	IntroduceStub        func(entity storage.Entity)
	introduceMutex       sync.RWMutex
	introduceArgsForCall []struct {
		entity storage.Entity
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeStorage) Open(options *storage.Settings, scheme *storage.Scheme) error {
	fake.openMutex.Lock()
	ret, specificReturn := fake.openReturnsOnCall[len(fake.openArgsForCall)]
	fake.openArgsForCall = append(fake.openArgsForCall, struct {
		options *storage.Settings
		scheme  *storage.Scheme
	}{options, scheme})
	fake.recordInvocation("Open", []interface{}{options, scheme})
	fake.openMutex.Unlock()
	if fake.OpenStub != nil {
		return fake.OpenStub(options, scheme)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.openReturns.result1
}

func (fake *FakeStorage) OpenCallCount() int {
	fake.openMutex.RLock()
	defer fake.openMutex.RUnlock()
	return len(fake.openArgsForCall)
}

func (fake *FakeStorage) OpenArgsForCall(i int) (*storage.Settings, *storage.Scheme) {
	fake.openMutex.RLock()
	defer fake.openMutex.RUnlock()
	return fake.openArgsForCall[i].options, fake.openArgsForCall[i].scheme
}

func (fake *FakeStorage) OpenReturns(result1 error) {
	fake.OpenStub = nil
	fake.openReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStorage) OpenReturnsOnCall(i int, result1 error) {
	fake.OpenStub = nil
	if fake.openReturnsOnCall == nil {
		fake.openReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.openReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeStorage) Close() error {
	fake.closeMutex.Lock()
	ret, specificReturn := fake.closeReturnsOnCall[len(fake.closeArgsForCall)]
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct{}{})
	fake.recordInvocation("Close", []interface{}{})
	fake.closeMutex.Unlock()
	if fake.CloseStub != nil {
		return fake.CloseStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.closeReturns.result1
}

func (fake *FakeStorage) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakeStorage) CloseReturns(result1 error) {
	fake.CloseStub = nil
	fake.closeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStorage) CloseReturnsOnCall(i int, result1 error) {
	fake.CloseStub = nil
	if fake.closeReturnsOnCall == nil {
		fake.closeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.closeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeStorage) Ping() error {
	fake.pingMutex.Lock()
	ret, specificReturn := fake.pingReturnsOnCall[len(fake.pingArgsForCall)]
	fake.pingArgsForCall = append(fake.pingArgsForCall, struct{}{})
	fake.recordInvocation("Ping", []interface{}{})
	fake.pingMutex.Unlock()
	if fake.PingStub != nil {
		return fake.PingStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.pingReturns.result1
}

func (fake *FakeStorage) PingCallCount() int {
	fake.pingMutex.RLock()
	defer fake.pingMutex.RUnlock()
	return len(fake.pingArgsForCall)
}

func (fake *FakeStorage) PingReturns(result1 error) {
	fake.PingStub = nil
	fake.pingReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStorage) PingReturnsOnCall(i int, result1 error) {
	fake.PingStub = nil
	if fake.pingReturnsOnCall == nil {
		fake.pingReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.pingReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeStorage) Create(ctx context.Context, obj types.Object) (string, error) {
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		ctx context.Context
		obj types.Object
	}{ctx, obj})
	fake.recordInvocation("Create", []interface{}{ctx, obj})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(ctx, obj)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createReturns.result1, fake.createReturns.result2
}

func (fake *FakeStorage) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeStorage) CreateArgsForCall(i int) (context.Context, types.Object) {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.createArgsForCall[i].ctx, fake.createArgsForCall[i].obj
}

func (fake *FakeStorage) CreateReturns(result1 string, result2 error) {
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeStorage) CreateReturnsOnCall(i int, result1 string, result2 error) {
	fake.CreateStub = nil
	if fake.createReturnsOnCall == nil {
		fake.createReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.createReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeStorage) Get(ctx context.Context, objectType types.ObjectType, id string) (types.Object, error) {
	fake.getMutex.Lock()
	ret, specificReturn := fake.getReturnsOnCall[len(fake.getArgsForCall)]
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		ctx        context.Context
		id         string
		objectType types.ObjectType
	}{ctx, id, objectType})
	fake.recordInvocation("Get", []interface{}{ctx, id, objectType})
	fake.getMutex.Unlock()
	if fake.GetStub != nil {
		return fake.GetStub(ctx, id, objectType)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getReturns.result1, fake.getReturns.result2
}

func (fake *FakeStorage) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakeStorage) GetArgsForCall(i int) (context.Context, string, types.ObjectType) {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return fake.getArgsForCall[i].ctx, fake.getArgsForCall[i].id, fake.getArgsForCall[i].objectType
}

func (fake *FakeStorage) GetReturns(result1 types.Object, result2 error) {
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 types.Object
		result2 error
	}{result1, result2}
}

func (fake *FakeStorage) GetReturnsOnCall(i int, result1 types.Object, result2 error) {
	fake.GetStub = nil
	if fake.getReturnsOnCall == nil {
		fake.getReturnsOnCall = make(map[int]struct {
			result1 types.Object
			result2 error
		})
	}
	fake.getReturnsOnCall[i] = struct {
		result1 types.Object
		result2 error
	}{result1, result2}
}

func (fake *FakeStorage) List(ctx context.Context, objectType types.ObjectType, criteria ...query.Criterion) (types.ObjectList, error) {
	fake.listMutex.Lock()
	ret, specificReturn := fake.listReturnsOnCall[len(fake.listArgsForCall)]
	fake.listArgsForCall = append(fake.listArgsForCall, struct {
		ctx        context.Context
		objectType types.ObjectType
		criteria   []query.Criterion
	}{ctx, objectType, criteria})
	fake.recordInvocation("List", []interface{}{ctx, objectType, criteria})
	fake.listMutex.Unlock()
	if fake.ListStub != nil {
		return fake.ListStub(ctx, objectType, criteria...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.listReturns.result1, fake.listReturns.result2
}

func (fake *FakeStorage) ListCallCount() int {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	return len(fake.listArgsForCall)
}

func (fake *FakeStorage) ListArgsForCall(i int) (context.Context, types.ObjectType, []query.Criterion) {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	return fake.listArgsForCall[i].ctx, fake.listArgsForCall[i].objectType, fake.listArgsForCall[i].criteria
}

func (fake *FakeStorage) ListReturns(result1 types.ObjectList, result2 error) {
	fake.ListStub = nil
	fake.listReturns = struct {
		result1 types.ObjectList
		result2 error
	}{result1, result2}
}

func (fake *FakeStorage) ListReturnsOnCall(i int, result1 types.ObjectList, result2 error) {
	fake.ListStub = nil
	if fake.listReturnsOnCall == nil {
		fake.listReturnsOnCall = make(map[int]struct {
			result1 types.ObjectList
			result2 error
		})
	}
	fake.listReturnsOnCall[i] = struct {
		result1 types.ObjectList
		result2 error
	}{result1, result2}
}

func (fake *FakeStorage) Delete(ctx context.Context, objectType types.ObjectType, criteria ...query.Criterion) (types.ObjectList, error) {
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		ctx        context.Context
		objectType types.ObjectType
		criteria   []query.Criterion
	}{ctx, objectType, criteria})
	fake.recordInvocation("Delete", []interface{}{ctx, objectType, criteria})
	fake.deleteMutex.Unlock()
	if fake.DeleteStub != nil {
		return fake.DeleteStub(ctx, objectType, criteria...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.deleteReturns.result1, fake.deleteReturns.result2
}

func (fake *FakeStorage) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeStorage) DeleteArgsForCall(i int) (context.Context, types.ObjectType, []query.Criterion) {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return fake.deleteArgsForCall[i].ctx, fake.deleteArgsForCall[i].objectType, fake.deleteArgsForCall[i].criteria
}

func (fake *FakeStorage) DeleteReturns(result1 types.ObjectList, result2 error) {
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 types.ObjectList
		result2 error
	}{result1, result2}
}

func (fake *FakeStorage) DeleteReturnsOnCall(i int, result1 types.ObjectList, result2 error) {
	fake.DeleteStub = nil
	if fake.deleteReturnsOnCall == nil {
		fake.deleteReturnsOnCall = make(map[int]struct {
			result1 types.ObjectList
			result2 error
		})
	}
	fake.deleteReturnsOnCall[i] = struct {
		result1 types.ObjectList
		result2 error
	}{result1, result2}
}

func (fake *FakeStorage) Update(ctx context.Context, obj types.Object, labelChanges ...*query.LabelChange) (types.Object, error) {
	fake.updateMutex.Lock()
	ret, specificReturn := fake.updateReturnsOnCall[len(fake.updateArgsForCall)]
	fake.updateArgsForCall = append(fake.updateArgsForCall, struct {
		ctx          context.Context
		obj          types.Object
		labelChanges []*query.LabelChange
	}{ctx, obj, labelChanges})
	fake.recordInvocation("Update", []interface{}{ctx, obj, labelChanges})
	fake.updateMutex.Unlock()
	if fake.UpdateStub != nil {
		return fake.UpdateStub(ctx, obj, labelChanges...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.updateReturns.result1, fake.updateReturns.result2
}

func (fake *FakeStorage) UpdateCallCount() int {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return len(fake.updateArgsForCall)
}

func (fake *FakeStorage) UpdateArgsForCall(i int) (context.Context, types.Object, []*query.LabelChange) {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return fake.updateArgsForCall[i].ctx, fake.updateArgsForCall[i].obj, fake.updateArgsForCall[i].labelChanges
}

func (fake *FakeStorage) UpdateReturns(result1 types.Object, result2 error) {
	fake.UpdateStub = nil
	fake.updateReturns = struct {
		result1 types.Object
		result2 error
	}{result1, result2}
}

func (fake *FakeStorage) UpdateReturnsOnCall(i int, result1 types.Object, result2 error) {
	fake.UpdateStub = nil
	if fake.updateReturnsOnCall == nil {
		fake.updateReturnsOnCall = make(map[int]struct {
			result1 types.Object
			result2 error
		})
	}
	fake.updateReturnsOnCall[i] = struct {
		result1 types.Object
		result2 error
	}{result1, result2}
}

func (fake *FakeStorage) ServiceOffering() storage.ServiceOffering {
	fake.serviceOfferingMutex.Lock()
	ret, specificReturn := fake.serviceOfferingReturnsOnCall[len(fake.serviceOfferingArgsForCall)]
	fake.serviceOfferingArgsForCall = append(fake.serviceOfferingArgsForCall, struct{}{})
	fake.recordInvocation("ServiceOffering", []interface{}{})
	fake.serviceOfferingMutex.Unlock()
	if fake.ServiceOfferingStub != nil {
		return fake.ServiceOfferingStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.serviceOfferingReturns.result1
}

func (fake *FakeStorage) ServiceOfferingCallCount() int {
	fake.serviceOfferingMutex.RLock()
	defer fake.serviceOfferingMutex.RUnlock()
	return len(fake.serviceOfferingArgsForCall)
}

func (fake *FakeStorage) ServiceOfferingReturns(result1 storage.ServiceOffering) {
	fake.ServiceOfferingStub = nil
	fake.serviceOfferingReturns = struct {
		result1 storage.ServiceOffering
	}{result1}
}

func (fake *FakeStorage) ServiceOfferingReturnsOnCall(i int, result1 storage.ServiceOffering) {
	fake.ServiceOfferingStub = nil
	if fake.serviceOfferingReturnsOnCall == nil {
		fake.serviceOfferingReturnsOnCall = make(map[int]struct {
			result1 storage.ServiceOffering
		})
	}
	fake.serviceOfferingReturnsOnCall[i] = struct {
		result1 storage.ServiceOffering
	}{result1}
}

func (fake *FakeStorage) Credentials() storage.Credentials {
	fake.credentialsMutex.Lock()
	ret, specificReturn := fake.credentialsReturnsOnCall[len(fake.credentialsArgsForCall)]
	fake.credentialsArgsForCall = append(fake.credentialsArgsForCall, struct{}{})
	fake.recordInvocation("Credentials", []interface{}{})
	fake.credentialsMutex.Unlock()
	if fake.CredentialsStub != nil {
		return fake.CredentialsStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.credentialsReturns.result1
}

func (fake *FakeStorage) CredentialsCallCount() int {
	fake.credentialsMutex.RLock()
	defer fake.credentialsMutex.RUnlock()
	return len(fake.credentialsArgsForCall)
}

func (fake *FakeStorage) CredentialsReturns(result1 storage.Credentials) {
	fake.CredentialsStub = nil
	fake.credentialsReturns = struct {
		result1 storage.Credentials
	}{result1}
}

func (fake *FakeStorage) CredentialsReturnsOnCall(i int, result1 storage.Credentials) {
	fake.CredentialsStub = nil
	if fake.credentialsReturnsOnCall == nil {
		fake.credentialsReturnsOnCall = make(map[int]struct {
			result1 storage.Credentials
		})
	}
	fake.credentialsReturnsOnCall[i] = struct {
		result1 storage.Credentials
	}{result1}
}

func (fake *FakeStorage) Security() storage.Security {
	fake.securityMutex.Lock()
	ret, specificReturn := fake.securityReturnsOnCall[len(fake.securityArgsForCall)]
	fake.securityArgsForCall = append(fake.securityArgsForCall, struct{}{})
	fake.recordInvocation("Security", []interface{}{})
	fake.securityMutex.Unlock()
	if fake.SecurityStub != nil {
		return fake.SecurityStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.securityReturns.result1
}

func (fake *FakeStorage) SecurityCallCount() int {
	fake.securityMutex.RLock()
	defer fake.securityMutex.RUnlock()
	return len(fake.securityArgsForCall)
}

func (fake *FakeStorage) SecurityReturns(result1 storage.Security) {
	fake.SecurityStub = nil
	fake.securityReturns = struct {
		result1 storage.Security
	}{result1}
}

func (fake *FakeStorage) SecurityReturnsOnCall(i int, result1 storage.Security) {
	fake.SecurityStub = nil
	if fake.securityReturnsOnCall == nil {
		fake.securityReturnsOnCall = make(map[int]struct {
			result1 storage.Security
		})
	}
	fake.securityReturnsOnCall[i] = struct {
		result1 storage.Security
	}{result1}
}

func (fake *FakeStorage) InTransaction(ctx context.Context, f func(ctx context.Context, storage storage.Warehouse) error) error {
	fake.inTransactionMutex.Lock()
	ret, specificReturn := fake.inTransactionReturnsOnCall[len(fake.inTransactionArgsForCall)]
	fake.inTransactionArgsForCall = append(fake.inTransactionArgsForCall, struct {
		ctx context.Context
		f   func(ctx context.Context, storage storage.Warehouse) error
	}{ctx, f})
	fake.recordInvocation("InTransaction", []interface{}{ctx, f})
	fake.inTransactionMutex.Unlock()
	if fake.InTransactionStub != nil {
		return fake.InTransactionStub(ctx, f)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.inTransactionReturns.result1
}

func (fake *FakeStorage) InTransactionCallCount() int {
	fake.inTransactionMutex.RLock()
	defer fake.inTransactionMutex.RUnlock()
	return len(fake.inTransactionArgsForCall)
}

func (fake *FakeStorage) InTransactionArgsForCall(i int) (context.Context, func(ctx context.Context, storage storage.Warehouse) error) {
	fake.inTransactionMutex.RLock()
	defer fake.inTransactionMutex.RUnlock()
	return fake.inTransactionArgsForCall[i].ctx, fake.inTransactionArgsForCall[i].f
}

func (fake *FakeStorage) InTransactionReturns(result1 error) {
	fake.InTransactionStub = nil
	fake.inTransactionReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStorage) InTransactionReturnsOnCall(i int, result1 error) {
	fake.InTransactionStub = nil
	if fake.inTransactionReturnsOnCall == nil {
		fake.inTransactionReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.inTransactionReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeStorage) Introduce(entity storage.Entity) {
	fake.introduceMutex.Lock()
	fake.introduceArgsForCall = append(fake.introduceArgsForCall, struct {
		entity storage.Entity
	}{entity})
	fake.recordInvocation("Introduce", []interface{}{entity})
	fake.introduceMutex.Unlock()
	if fake.IntroduceStub != nil {
		fake.IntroduceStub(entity)
	}
}

func (fake *FakeStorage) IntroduceCallCount() int {
	fake.introduceMutex.RLock()
	defer fake.introduceMutex.RUnlock()
	return len(fake.introduceArgsForCall)
}

func (fake *FakeStorage) IntroduceArgsForCall(i int) storage.Entity {
	fake.introduceMutex.RLock()
	defer fake.introduceMutex.RUnlock()
	return fake.introduceArgsForCall[i].entity
}

func (fake *FakeStorage) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.openMutex.RLock()
	defer fake.openMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	fake.pingMutex.RLock()
	defer fake.pingMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	fake.serviceOfferingMutex.RLock()
	defer fake.serviceOfferingMutex.RUnlock()
	fake.credentialsMutex.RLock()
	defer fake.credentialsMutex.RUnlock()
	fake.securityMutex.RLock()
	defer fake.securityMutex.RUnlock()
	fake.inTransactionMutex.RLock()
	defer fake.inTransactionMutex.RUnlock()
	fake.introduceMutex.RLock()
	defer fake.introduceMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeStorage) recordInvocation(key string, args []interface{}) {
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

var _ storage.Storage = new(FakeStorage)
