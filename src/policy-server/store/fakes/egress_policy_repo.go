// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"policy-server/db"
	"policy-server/store"
	"sync"
)

type EgressPolicyRepo struct {
	CreateAppStub        func(tx db.Transaction, sourceTerminalGUID string, appGUID string) (int64, error)
	createAppMutex       sync.RWMutex
	createAppArgsForCall []struct {
		tx                 db.Transaction
		sourceTerminalGUID string
		appGUID            string
	}
	createAppReturns struct {
		result1 int64
		result2 error
	}
	createAppReturnsOnCall map[int]struct {
		result1 int64
		result2 error
	}
	CreateIPRangeStub        func(tx db.Transaction, destinationTerminalGUID string, startIP, endIP, protocol string, startPort, endPort, icmpType, icmpCode int64) (int64, error)
	createIPRangeMutex       sync.RWMutex
	createIPRangeArgsForCall []struct {
		tx                      db.Transaction
		destinationTerminalGUID string
		startIP                 string
		endIP                   string
		protocol                string
		startPort               int64
		endPort                 int64
		icmpType                int64
		icmpCode                int64
	}
	createIPRangeReturns struct {
		result1 int64
		result2 error
	}
	createIPRangeReturnsOnCall map[int]struct {
		result1 int64
		result2 error
	}
	CreateEgressPolicyStub        func(tx db.Transaction, sourceTerminalGUID, destinationTerminalGUID string) (string, error)
	createEgressPolicyMutex       sync.RWMutex
	createEgressPolicyArgsForCall []struct {
		tx                      db.Transaction
		sourceTerminalGUID      string
		destinationTerminalGUID string
	}
	createEgressPolicyReturns struct {
		result1 string
		result2 error
	}
	createEgressPolicyReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	CreateSpaceStub        func(tx db.Transaction, sourceTerminalGUID string, spaceGUID string) (int64, error)
	createSpaceMutex       sync.RWMutex
	createSpaceArgsForCall []struct {
		tx                 db.Transaction
		sourceTerminalGUID string
		spaceGUID          string
	}
	createSpaceReturns struct {
		result1 int64
		result2 error
	}
	createSpaceReturnsOnCall map[int]struct {
		result1 int64
		result2 error
	}
	GetTerminalByAppGUIDStub        func(tx db.Transaction, appGUID string) (string, error)
	getTerminalByAppGUIDMutex       sync.RWMutex
	getTerminalByAppGUIDArgsForCall []struct {
		tx      db.Transaction
		appGUID string
	}
	getTerminalByAppGUIDReturns struct {
		result1 string
		result2 error
	}
	getTerminalByAppGUIDReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	GetTerminalBySpaceGUIDStub        func(tx db.Transaction, appGUID string) (string, error)
	getTerminalBySpaceGUIDMutex       sync.RWMutex
	getTerminalBySpaceGUIDArgsForCall []struct {
		tx      db.Transaction
		appGUID string
	}
	getTerminalBySpaceGUIDReturns struct {
		result1 string
		result2 error
	}
	getTerminalBySpaceGUIDReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	GetAllPoliciesStub        func() ([]store.EgressPolicy, error)
	getAllPoliciesMutex       sync.RWMutex
	getAllPoliciesArgsForCall []struct{}
	getAllPoliciesReturns     struct {
		result1 []store.EgressPolicy
		result2 error
	}
	getAllPoliciesReturnsOnCall map[int]struct {
		result1 []store.EgressPolicy
		result2 error
	}
	GetByGuidsStub        func(ids []string) ([]store.EgressPolicy, error)
	getByGuidsMutex       sync.RWMutex
	getByGuidsArgsForCall []struct {
		ids []string
	}
	getByGuidsReturns struct {
		result1 []store.EgressPolicy
		result2 error
	}
	getByGuidsReturnsOnCall map[int]struct {
		result1 []store.EgressPolicy
		result2 error
	}
	GetIDCollectionsByEgressPolicyStub        func(tx db.Transaction, egressPolicy store.EgressPolicy) ([]store.EgressPolicyIDCollection, error)
	getIDCollectionsByEgressPolicyMutex       sync.RWMutex
	getIDCollectionsByEgressPolicyArgsForCall []struct {
		tx           db.Transaction
		egressPolicy store.EgressPolicy
	}
	getIDCollectionsByEgressPolicyReturns struct {
		result1 []store.EgressPolicyIDCollection
		result2 error
	}
	getIDCollectionsByEgressPolicyReturnsOnCall map[int]struct {
		result1 []store.EgressPolicyIDCollection
		result2 error
	}
	DeleteEgressPolicyStub        func(tx db.Transaction, egressPolicyGUID string) error
	deleteEgressPolicyMutex       sync.RWMutex
	deleteEgressPolicyArgsForCall []struct {
		tx               db.Transaction
		egressPolicyGUID string
	}
	deleteEgressPolicyReturns struct {
		result1 error
	}
	deleteEgressPolicyReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteIPRangeStub        func(tx db.Transaction, ipRangeID int64) error
	deleteIPRangeMutex       sync.RWMutex
	deleteIPRangeArgsForCall []struct {
		tx        db.Transaction
		ipRangeID int64
	}
	deleteIPRangeReturns struct {
		result1 error
	}
	deleteIPRangeReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteAppStub        func(tx db.Transaction, appID int64) error
	deleteAppMutex       sync.RWMutex
	deleteAppArgsForCall []struct {
		tx    db.Transaction
		appID int64
	}
	deleteAppReturns struct {
		result1 error
	}
	deleteAppReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteSpaceStub        func(tx db.Transaction, spaceID int64) error
	deleteSpaceMutex       sync.RWMutex
	deleteSpaceArgsForCall []struct {
		tx      db.Transaction
		spaceID int64
	}
	deleteSpaceReturns struct {
		result1 error
	}
	deleteSpaceReturnsOnCall map[int]struct {
		result1 error
	}
	IsTerminalInUseStub        func(tx db.Transaction, terminalGUID string) (bool, error)
	isTerminalInUseMutex       sync.RWMutex
	isTerminalInUseArgsForCall []struct {
		tx           db.Transaction
		terminalGUID string
	}
	isTerminalInUseReturns struct {
		result1 bool
		result2 error
	}
	isTerminalInUseReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *EgressPolicyRepo) CreateApp(tx db.Transaction, sourceTerminalGUID string, appGUID string) (int64, error) {
	fake.createAppMutex.Lock()
	ret, specificReturn := fake.createAppReturnsOnCall[len(fake.createAppArgsForCall)]
	fake.createAppArgsForCall = append(fake.createAppArgsForCall, struct {
		tx                 db.Transaction
		sourceTerminalGUID string
		appGUID            string
	}{tx, sourceTerminalGUID, appGUID})
	fake.recordInvocation("CreateApp", []interface{}{tx, sourceTerminalGUID, appGUID})
	fake.createAppMutex.Unlock()
	if fake.CreateAppStub != nil {
		return fake.CreateAppStub(tx, sourceTerminalGUID, appGUID)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createAppReturns.result1, fake.createAppReturns.result2
}

func (fake *EgressPolicyRepo) CreateAppCallCount() int {
	fake.createAppMutex.RLock()
	defer fake.createAppMutex.RUnlock()
	return len(fake.createAppArgsForCall)
}

func (fake *EgressPolicyRepo) CreateAppArgsForCall(i int) (db.Transaction, string, string) {
	fake.createAppMutex.RLock()
	defer fake.createAppMutex.RUnlock()
	return fake.createAppArgsForCall[i].tx, fake.createAppArgsForCall[i].sourceTerminalGUID, fake.createAppArgsForCall[i].appGUID
}

func (fake *EgressPolicyRepo) CreateAppReturns(result1 int64, result2 error) {
	fake.CreateAppStub = nil
	fake.createAppReturns = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyRepo) CreateAppReturnsOnCall(i int, result1 int64, result2 error) {
	fake.CreateAppStub = nil
	if fake.createAppReturnsOnCall == nil {
		fake.createAppReturnsOnCall = make(map[int]struct {
			result1 int64
			result2 error
		})
	}
	fake.createAppReturnsOnCall[i] = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyRepo) CreateIPRange(tx db.Transaction, destinationTerminalGUID string, startIP string, endIP string, protocol string, startPort int64, endPort int64, icmpType int64, icmpCode int64) (int64, error) {
	fake.createIPRangeMutex.Lock()
	ret, specificReturn := fake.createIPRangeReturnsOnCall[len(fake.createIPRangeArgsForCall)]
	fake.createIPRangeArgsForCall = append(fake.createIPRangeArgsForCall, struct {
		tx                      db.Transaction
		destinationTerminalGUID string
		startIP                 string
		endIP                   string
		protocol                string
		startPort               int64
		endPort                 int64
		icmpType                int64
		icmpCode                int64
	}{tx, destinationTerminalGUID, startIP, endIP, protocol, startPort, endPort, icmpType, icmpCode})
	fake.recordInvocation("CreateIPRange", []interface{}{tx, destinationTerminalGUID, startIP, endIP, protocol, startPort, endPort, icmpType, icmpCode})
	fake.createIPRangeMutex.Unlock()
	if fake.CreateIPRangeStub != nil {
		return fake.CreateIPRangeStub(tx, destinationTerminalGUID, startIP, endIP, protocol, startPort, endPort, icmpType, icmpCode)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createIPRangeReturns.result1, fake.createIPRangeReturns.result2
}

func (fake *EgressPolicyRepo) CreateIPRangeCallCount() int {
	fake.createIPRangeMutex.RLock()
	defer fake.createIPRangeMutex.RUnlock()
	return len(fake.createIPRangeArgsForCall)
}

func (fake *EgressPolicyRepo) CreateIPRangeArgsForCall(i int) (db.Transaction, string, string, string, string, int64, int64, int64, int64) {
	fake.createIPRangeMutex.RLock()
	defer fake.createIPRangeMutex.RUnlock()
	return fake.createIPRangeArgsForCall[i].tx, fake.createIPRangeArgsForCall[i].destinationTerminalGUID, fake.createIPRangeArgsForCall[i].startIP, fake.createIPRangeArgsForCall[i].endIP, fake.createIPRangeArgsForCall[i].protocol, fake.createIPRangeArgsForCall[i].startPort, fake.createIPRangeArgsForCall[i].endPort, fake.createIPRangeArgsForCall[i].icmpType, fake.createIPRangeArgsForCall[i].icmpCode
}

func (fake *EgressPolicyRepo) CreateIPRangeReturns(result1 int64, result2 error) {
	fake.CreateIPRangeStub = nil
	fake.createIPRangeReturns = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyRepo) CreateIPRangeReturnsOnCall(i int, result1 int64, result2 error) {
	fake.CreateIPRangeStub = nil
	if fake.createIPRangeReturnsOnCall == nil {
		fake.createIPRangeReturnsOnCall = make(map[int]struct {
			result1 int64
			result2 error
		})
	}
	fake.createIPRangeReturnsOnCall[i] = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyRepo) CreateEgressPolicy(tx db.Transaction, sourceTerminalGUID string, destinationTerminalGUID string) (string, error) {
	fake.createEgressPolicyMutex.Lock()
	ret, specificReturn := fake.createEgressPolicyReturnsOnCall[len(fake.createEgressPolicyArgsForCall)]
	fake.createEgressPolicyArgsForCall = append(fake.createEgressPolicyArgsForCall, struct {
		tx                      db.Transaction
		sourceTerminalGUID      string
		destinationTerminalGUID string
	}{tx, sourceTerminalGUID, destinationTerminalGUID})
	fake.recordInvocation("CreateEgressPolicy", []interface{}{tx, sourceTerminalGUID, destinationTerminalGUID})
	fake.createEgressPolicyMutex.Unlock()
	if fake.CreateEgressPolicyStub != nil {
		return fake.CreateEgressPolicyStub(tx, sourceTerminalGUID, destinationTerminalGUID)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createEgressPolicyReturns.result1, fake.createEgressPolicyReturns.result2
}

func (fake *EgressPolicyRepo) CreateEgressPolicyCallCount() int {
	fake.createEgressPolicyMutex.RLock()
	defer fake.createEgressPolicyMutex.RUnlock()
	return len(fake.createEgressPolicyArgsForCall)
}

func (fake *EgressPolicyRepo) CreateEgressPolicyArgsForCall(i int) (db.Transaction, string, string) {
	fake.createEgressPolicyMutex.RLock()
	defer fake.createEgressPolicyMutex.RUnlock()
	return fake.createEgressPolicyArgsForCall[i].tx, fake.createEgressPolicyArgsForCall[i].sourceTerminalGUID, fake.createEgressPolicyArgsForCall[i].destinationTerminalGUID
}

func (fake *EgressPolicyRepo) CreateEgressPolicyReturns(result1 string, result2 error) {
	fake.CreateEgressPolicyStub = nil
	fake.createEgressPolicyReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyRepo) CreateEgressPolicyReturnsOnCall(i int, result1 string, result2 error) {
	fake.CreateEgressPolicyStub = nil
	if fake.createEgressPolicyReturnsOnCall == nil {
		fake.createEgressPolicyReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.createEgressPolicyReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyRepo) CreateSpace(tx db.Transaction, sourceTerminalGUID string, spaceGUID string) (int64, error) {
	fake.createSpaceMutex.Lock()
	ret, specificReturn := fake.createSpaceReturnsOnCall[len(fake.createSpaceArgsForCall)]
	fake.createSpaceArgsForCall = append(fake.createSpaceArgsForCall, struct {
		tx                 db.Transaction
		sourceTerminalGUID string
		spaceGUID          string
	}{tx, sourceTerminalGUID, spaceGUID})
	fake.recordInvocation("CreateSpace", []interface{}{tx, sourceTerminalGUID, spaceGUID})
	fake.createSpaceMutex.Unlock()
	if fake.CreateSpaceStub != nil {
		return fake.CreateSpaceStub(tx, sourceTerminalGUID, spaceGUID)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createSpaceReturns.result1, fake.createSpaceReturns.result2
}

func (fake *EgressPolicyRepo) CreateSpaceCallCount() int {
	fake.createSpaceMutex.RLock()
	defer fake.createSpaceMutex.RUnlock()
	return len(fake.createSpaceArgsForCall)
}

func (fake *EgressPolicyRepo) CreateSpaceArgsForCall(i int) (db.Transaction, string, string) {
	fake.createSpaceMutex.RLock()
	defer fake.createSpaceMutex.RUnlock()
	return fake.createSpaceArgsForCall[i].tx, fake.createSpaceArgsForCall[i].sourceTerminalGUID, fake.createSpaceArgsForCall[i].spaceGUID
}

func (fake *EgressPolicyRepo) CreateSpaceReturns(result1 int64, result2 error) {
	fake.CreateSpaceStub = nil
	fake.createSpaceReturns = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyRepo) CreateSpaceReturnsOnCall(i int, result1 int64, result2 error) {
	fake.CreateSpaceStub = nil
	if fake.createSpaceReturnsOnCall == nil {
		fake.createSpaceReturnsOnCall = make(map[int]struct {
			result1 int64
			result2 error
		})
	}
	fake.createSpaceReturnsOnCall[i] = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyRepo) GetTerminalByAppGUID(tx db.Transaction, appGUID string) (string, error) {
	fake.getTerminalByAppGUIDMutex.Lock()
	ret, specificReturn := fake.getTerminalByAppGUIDReturnsOnCall[len(fake.getTerminalByAppGUIDArgsForCall)]
	fake.getTerminalByAppGUIDArgsForCall = append(fake.getTerminalByAppGUIDArgsForCall, struct {
		tx      db.Transaction
		appGUID string
	}{tx, appGUID})
	fake.recordInvocation("GetTerminalByAppGUID", []interface{}{tx, appGUID})
	fake.getTerminalByAppGUIDMutex.Unlock()
	if fake.GetTerminalByAppGUIDStub != nil {
		return fake.GetTerminalByAppGUIDStub(tx, appGUID)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getTerminalByAppGUIDReturns.result1, fake.getTerminalByAppGUIDReturns.result2
}

func (fake *EgressPolicyRepo) GetTerminalByAppGUIDCallCount() int {
	fake.getTerminalByAppGUIDMutex.RLock()
	defer fake.getTerminalByAppGUIDMutex.RUnlock()
	return len(fake.getTerminalByAppGUIDArgsForCall)
}

func (fake *EgressPolicyRepo) GetTerminalByAppGUIDArgsForCall(i int) (db.Transaction, string) {
	fake.getTerminalByAppGUIDMutex.RLock()
	defer fake.getTerminalByAppGUIDMutex.RUnlock()
	return fake.getTerminalByAppGUIDArgsForCall[i].tx, fake.getTerminalByAppGUIDArgsForCall[i].appGUID
}

func (fake *EgressPolicyRepo) GetTerminalByAppGUIDReturns(result1 string, result2 error) {
	fake.GetTerminalByAppGUIDStub = nil
	fake.getTerminalByAppGUIDReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyRepo) GetTerminalByAppGUIDReturnsOnCall(i int, result1 string, result2 error) {
	fake.GetTerminalByAppGUIDStub = nil
	if fake.getTerminalByAppGUIDReturnsOnCall == nil {
		fake.getTerminalByAppGUIDReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.getTerminalByAppGUIDReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyRepo) GetTerminalBySpaceGUID(tx db.Transaction, appGUID string) (string, error) {
	fake.getTerminalBySpaceGUIDMutex.Lock()
	ret, specificReturn := fake.getTerminalBySpaceGUIDReturnsOnCall[len(fake.getTerminalBySpaceGUIDArgsForCall)]
	fake.getTerminalBySpaceGUIDArgsForCall = append(fake.getTerminalBySpaceGUIDArgsForCall, struct {
		tx      db.Transaction
		appGUID string
	}{tx, appGUID})
	fake.recordInvocation("GetTerminalBySpaceGUID", []interface{}{tx, appGUID})
	fake.getTerminalBySpaceGUIDMutex.Unlock()
	if fake.GetTerminalBySpaceGUIDStub != nil {
		return fake.GetTerminalBySpaceGUIDStub(tx, appGUID)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getTerminalBySpaceGUIDReturns.result1, fake.getTerminalBySpaceGUIDReturns.result2
}

func (fake *EgressPolicyRepo) GetTerminalBySpaceGUIDCallCount() int {
	fake.getTerminalBySpaceGUIDMutex.RLock()
	defer fake.getTerminalBySpaceGUIDMutex.RUnlock()
	return len(fake.getTerminalBySpaceGUIDArgsForCall)
}

func (fake *EgressPolicyRepo) GetTerminalBySpaceGUIDArgsForCall(i int) (db.Transaction, string) {
	fake.getTerminalBySpaceGUIDMutex.RLock()
	defer fake.getTerminalBySpaceGUIDMutex.RUnlock()
	return fake.getTerminalBySpaceGUIDArgsForCall[i].tx, fake.getTerminalBySpaceGUIDArgsForCall[i].appGUID
}

func (fake *EgressPolicyRepo) GetTerminalBySpaceGUIDReturns(result1 string, result2 error) {
	fake.GetTerminalBySpaceGUIDStub = nil
	fake.getTerminalBySpaceGUIDReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyRepo) GetTerminalBySpaceGUIDReturnsOnCall(i int, result1 string, result2 error) {
	fake.GetTerminalBySpaceGUIDStub = nil
	if fake.getTerminalBySpaceGUIDReturnsOnCall == nil {
		fake.getTerminalBySpaceGUIDReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.getTerminalBySpaceGUIDReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyRepo) GetAllPolicies() ([]store.EgressPolicy, error) {
	fake.getAllPoliciesMutex.Lock()
	ret, specificReturn := fake.getAllPoliciesReturnsOnCall[len(fake.getAllPoliciesArgsForCall)]
	fake.getAllPoliciesArgsForCall = append(fake.getAllPoliciesArgsForCall, struct{}{})
	fake.recordInvocation("GetAllPolicies", []interface{}{})
	fake.getAllPoliciesMutex.Unlock()
	if fake.GetAllPoliciesStub != nil {
		return fake.GetAllPoliciesStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getAllPoliciesReturns.result1, fake.getAllPoliciesReturns.result2
}

func (fake *EgressPolicyRepo) GetAllPoliciesCallCount() int {
	fake.getAllPoliciesMutex.RLock()
	defer fake.getAllPoliciesMutex.RUnlock()
	return len(fake.getAllPoliciesArgsForCall)
}

func (fake *EgressPolicyRepo) GetAllPoliciesReturns(result1 []store.EgressPolicy, result2 error) {
	fake.GetAllPoliciesStub = nil
	fake.getAllPoliciesReturns = struct {
		result1 []store.EgressPolicy
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyRepo) GetAllPoliciesReturnsOnCall(i int, result1 []store.EgressPolicy, result2 error) {
	fake.GetAllPoliciesStub = nil
	if fake.getAllPoliciesReturnsOnCall == nil {
		fake.getAllPoliciesReturnsOnCall = make(map[int]struct {
			result1 []store.EgressPolicy
			result2 error
		})
	}
	fake.getAllPoliciesReturnsOnCall[i] = struct {
		result1 []store.EgressPolicy
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyRepo) GetByGuids(ids []string) ([]store.EgressPolicy, error) {
	var idsCopy []string
	if ids != nil {
		idsCopy = make([]string, len(ids))
		copy(idsCopy, ids)
	}
	fake.getByGuidsMutex.Lock()
	ret, specificReturn := fake.getByGuidsReturnsOnCall[len(fake.getByGuidsArgsForCall)]
	fake.getByGuidsArgsForCall = append(fake.getByGuidsArgsForCall, struct {
		ids []string
	}{idsCopy})
	fake.recordInvocation("GetByGuids", []interface{}{idsCopy})
	fake.getByGuidsMutex.Unlock()
	if fake.GetByGuidsStub != nil {
		return fake.GetByGuidsStub(ids)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getByGuidsReturns.result1, fake.getByGuidsReturns.result2
}

func (fake *EgressPolicyRepo) GetByGuidsCallCount() int {
	fake.getByGuidsMutex.RLock()
	defer fake.getByGuidsMutex.RUnlock()
	return len(fake.getByGuidsArgsForCall)
}

func (fake *EgressPolicyRepo) GetByGuidsArgsForCall(i int) []string {
	fake.getByGuidsMutex.RLock()
	defer fake.getByGuidsMutex.RUnlock()
	return fake.getByGuidsArgsForCall[i].ids
}

func (fake *EgressPolicyRepo) GetByGuidsReturns(result1 []store.EgressPolicy, result2 error) {
	fake.GetByGuidsStub = nil
	fake.getByGuidsReturns = struct {
		result1 []store.EgressPolicy
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyRepo) GetByGuidsReturnsOnCall(i int, result1 []store.EgressPolicy, result2 error) {
	fake.GetByGuidsStub = nil
	if fake.getByGuidsReturnsOnCall == nil {
		fake.getByGuidsReturnsOnCall = make(map[int]struct {
			result1 []store.EgressPolicy
			result2 error
		})
	}
	fake.getByGuidsReturnsOnCall[i] = struct {
		result1 []store.EgressPolicy
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyRepo) GetIDCollectionsByEgressPolicy(tx db.Transaction, egressPolicy store.EgressPolicy) ([]store.EgressPolicyIDCollection, error) {
	fake.getIDCollectionsByEgressPolicyMutex.Lock()
	ret, specificReturn := fake.getIDCollectionsByEgressPolicyReturnsOnCall[len(fake.getIDCollectionsByEgressPolicyArgsForCall)]
	fake.getIDCollectionsByEgressPolicyArgsForCall = append(fake.getIDCollectionsByEgressPolicyArgsForCall, struct {
		tx           db.Transaction
		egressPolicy store.EgressPolicy
	}{tx, egressPolicy})
	fake.recordInvocation("GetIDCollectionsByEgressPolicy", []interface{}{tx, egressPolicy})
	fake.getIDCollectionsByEgressPolicyMutex.Unlock()
	if fake.GetIDCollectionsByEgressPolicyStub != nil {
		return fake.GetIDCollectionsByEgressPolicyStub(tx, egressPolicy)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getIDCollectionsByEgressPolicyReturns.result1, fake.getIDCollectionsByEgressPolicyReturns.result2
}

func (fake *EgressPolicyRepo) GetIDCollectionsByEgressPolicyCallCount() int {
	fake.getIDCollectionsByEgressPolicyMutex.RLock()
	defer fake.getIDCollectionsByEgressPolicyMutex.RUnlock()
	return len(fake.getIDCollectionsByEgressPolicyArgsForCall)
}

func (fake *EgressPolicyRepo) GetIDCollectionsByEgressPolicyArgsForCall(i int) (db.Transaction, store.EgressPolicy) {
	fake.getIDCollectionsByEgressPolicyMutex.RLock()
	defer fake.getIDCollectionsByEgressPolicyMutex.RUnlock()
	return fake.getIDCollectionsByEgressPolicyArgsForCall[i].tx, fake.getIDCollectionsByEgressPolicyArgsForCall[i].egressPolicy
}

func (fake *EgressPolicyRepo) GetIDCollectionsByEgressPolicyReturns(result1 []store.EgressPolicyIDCollection, result2 error) {
	fake.GetIDCollectionsByEgressPolicyStub = nil
	fake.getIDCollectionsByEgressPolicyReturns = struct {
		result1 []store.EgressPolicyIDCollection
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyRepo) GetIDCollectionsByEgressPolicyReturnsOnCall(i int, result1 []store.EgressPolicyIDCollection, result2 error) {
	fake.GetIDCollectionsByEgressPolicyStub = nil
	if fake.getIDCollectionsByEgressPolicyReturnsOnCall == nil {
		fake.getIDCollectionsByEgressPolicyReturnsOnCall = make(map[int]struct {
			result1 []store.EgressPolicyIDCollection
			result2 error
		})
	}
	fake.getIDCollectionsByEgressPolicyReturnsOnCall[i] = struct {
		result1 []store.EgressPolicyIDCollection
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyRepo) DeleteEgressPolicy(tx db.Transaction, egressPolicyGUID string) error {
	fake.deleteEgressPolicyMutex.Lock()
	ret, specificReturn := fake.deleteEgressPolicyReturnsOnCall[len(fake.deleteEgressPolicyArgsForCall)]
	fake.deleteEgressPolicyArgsForCall = append(fake.deleteEgressPolicyArgsForCall, struct {
		tx               db.Transaction
		egressPolicyGUID string
	}{tx, egressPolicyGUID})
	fake.recordInvocation("DeleteEgressPolicy", []interface{}{tx, egressPolicyGUID})
	fake.deleteEgressPolicyMutex.Unlock()
	if fake.DeleteEgressPolicyStub != nil {
		return fake.DeleteEgressPolicyStub(tx, egressPolicyGUID)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.deleteEgressPolicyReturns.result1
}

func (fake *EgressPolicyRepo) DeleteEgressPolicyCallCount() int {
	fake.deleteEgressPolicyMutex.RLock()
	defer fake.deleteEgressPolicyMutex.RUnlock()
	return len(fake.deleteEgressPolicyArgsForCall)
}

func (fake *EgressPolicyRepo) DeleteEgressPolicyArgsForCall(i int) (db.Transaction, string) {
	fake.deleteEgressPolicyMutex.RLock()
	defer fake.deleteEgressPolicyMutex.RUnlock()
	return fake.deleteEgressPolicyArgsForCall[i].tx, fake.deleteEgressPolicyArgsForCall[i].egressPolicyGUID
}

func (fake *EgressPolicyRepo) DeleteEgressPolicyReturns(result1 error) {
	fake.DeleteEgressPolicyStub = nil
	fake.deleteEgressPolicyReturns = struct {
		result1 error
	}{result1}
}

func (fake *EgressPolicyRepo) DeleteEgressPolicyReturnsOnCall(i int, result1 error) {
	fake.DeleteEgressPolicyStub = nil
	if fake.deleteEgressPolicyReturnsOnCall == nil {
		fake.deleteEgressPolicyReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteEgressPolicyReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *EgressPolicyRepo) DeleteIPRange(tx db.Transaction, ipRangeID int64) error {
	fake.deleteIPRangeMutex.Lock()
	ret, specificReturn := fake.deleteIPRangeReturnsOnCall[len(fake.deleteIPRangeArgsForCall)]
	fake.deleteIPRangeArgsForCall = append(fake.deleteIPRangeArgsForCall, struct {
		tx        db.Transaction
		ipRangeID int64
	}{tx, ipRangeID})
	fake.recordInvocation("DeleteIPRange", []interface{}{tx, ipRangeID})
	fake.deleteIPRangeMutex.Unlock()
	if fake.DeleteIPRangeStub != nil {
		return fake.DeleteIPRangeStub(tx, ipRangeID)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.deleteIPRangeReturns.result1
}

func (fake *EgressPolicyRepo) DeleteIPRangeCallCount() int {
	fake.deleteIPRangeMutex.RLock()
	defer fake.deleteIPRangeMutex.RUnlock()
	return len(fake.deleteIPRangeArgsForCall)
}

func (fake *EgressPolicyRepo) DeleteIPRangeArgsForCall(i int) (db.Transaction, int64) {
	fake.deleteIPRangeMutex.RLock()
	defer fake.deleteIPRangeMutex.RUnlock()
	return fake.deleteIPRangeArgsForCall[i].tx, fake.deleteIPRangeArgsForCall[i].ipRangeID
}

func (fake *EgressPolicyRepo) DeleteIPRangeReturns(result1 error) {
	fake.DeleteIPRangeStub = nil
	fake.deleteIPRangeReturns = struct {
		result1 error
	}{result1}
}

func (fake *EgressPolicyRepo) DeleteIPRangeReturnsOnCall(i int, result1 error) {
	fake.DeleteIPRangeStub = nil
	if fake.deleteIPRangeReturnsOnCall == nil {
		fake.deleteIPRangeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteIPRangeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *EgressPolicyRepo) DeleteApp(tx db.Transaction, appID int64) error {
	fake.deleteAppMutex.Lock()
	ret, specificReturn := fake.deleteAppReturnsOnCall[len(fake.deleteAppArgsForCall)]
	fake.deleteAppArgsForCall = append(fake.deleteAppArgsForCall, struct {
		tx    db.Transaction
		appID int64
	}{tx, appID})
	fake.recordInvocation("DeleteApp", []interface{}{tx, appID})
	fake.deleteAppMutex.Unlock()
	if fake.DeleteAppStub != nil {
		return fake.DeleteAppStub(tx, appID)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.deleteAppReturns.result1
}

func (fake *EgressPolicyRepo) DeleteAppCallCount() int {
	fake.deleteAppMutex.RLock()
	defer fake.deleteAppMutex.RUnlock()
	return len(fake.deleteAppArgsForCall)
}

func (fake *EgressPolicyRepo) DeleteAppArgsForCall(i int) (db.Transaction, int64) {
	fake.deleteAppMutex.RLock()
	defer fake.deleteAppMutex.RUnlock()
	return fake.deleteAppArgsForCall[i].tx, fake.deleteAppArgsForCall[i].appID
}

func (fake *EgressPolicyRepo) DeleteAppReturns(result1 error) {
	fake.DeleteAppStub = nil
	fake.deleteAppReturns = struct {
		result1 error
	}{result1}
}

func (fake *EgressPolicyRepo) DeleteAppReturnsOnCall(i int, result1 error) {
	fake.DeleteAppStub = nil
	if fake.deleteAppReturnsOnCall == nil {
		fake.deleteAppReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteAppReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *EgressPolicyRepo) DeleteSpace(tx db.Transaction, spaceID int64) error {
	fake.deleteSpaceMutex.Lock()
	ret, specificReturn := fake.deleteSpaceReturnsOnCall[len(fake.deleteSpaceArgsForCall)]
	fake.deleteSpaceArgsForCall = append(fake.deleteSpaceArgsForCall, struct {
		tx      db.Transaction
		spaceID int64
	}{tx, spaceID})
	fake.recordInvocation("DeleteSpace", []interface{}{tx, spaceID})
	fake.deleteSpaceMutex.Unlock()
	if fake.DeleteSpaceStub != nil {
		return fake.DeleteSpaceStub(tx, spaceID)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.deleteSpaceReturns.result1
}

func (fake *EgressPolicyRepo) DeleteSpaceCallCount() int {
	fake.deleteSpaceMutex.RLock()
	defer fake.deleteSpaceMutex.RUnlock()
	return len(fake.deleteSpaceArgsForCall)
}

func (fake *EgressPolicyRepo) DeleteSpaceArgsForCall(i int) (db.Transaction, int64) {
	fake.deleteSpaceMutex.RLock()
	defer fake.deleteSpaceMutex.RUnlock()
	return fake.deleteSpaceArgsForCall[i].tx, fake.deleteSpaceArgsForCall[i].spaceID
}

func (fake *EgressPolicyRepo) DeleteSpaceReturns(result1 error) {
	fake.DeleteSpaceStub = nil
	fake.deleteSpaceReturns = struct {
		result1 error
	}{result1}
}

func (fake *EgressPolicyRepo) DeleteSpaceReturnsOnCall(i int, result1 error) {
	fake.DeleteSpaceStub = nil
	if fake.deleteSpaceReturnsOnCall == nil {
		fake.deleteSpaceReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteSpaceReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *EgressPolicyRepo) IsTerminalInUse(tx db.Transaction, terminalGUID string) (bool, error) {
	fake.isTerminalInUseMutex.Lock()
	ret, specificReturn := fake.isTerminalInUseReturnsOnCall[len(fake.isTerminalInUseArgsForCall)]
	fake.isTerminalInUseArgsForCall = append(fake.isTerminalInUseArgsForCall, struct {
		tx           db.Transaction
		terminalGUID string
	}{tx, terminalGUID})
	fake.recordInvocation("IsTerminalInUse", []interface{}{tx, terminalGUID})
	fake.isTerminalInUseMutex.Unlock()
	if fake.IsTerminalInUseStub != nil {
		return fake.IsTerminalInUseStub(tx, terminalGUID)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.isTerminalInUseReturns.result1, fake.isTerminalInUseReturns.result2
}

func (fake *EgressPolicyRepo) IsTerminalInUseCallCount() int {
	fake.isTerminalInUseMutex.RLock()
	defer fake.isTerminalInUseMutex.RUnlock()
	return len(fake.isTerminalInUseArgsForCall)
}

func (fake *EgressPolicyRepo) IsTerminalInUseArgsForCall(i int) (db.Transaction, string) {
	fake.isTerminalInUseMutex.RLock()
	defer fake.isTerminalInUseMutex.RUnlock()
	return fake.isTerminalInUseArgsForCall[i].tx, fake.isTerminalInUseArgsForCall[i].terminalGUID
}

func (fake *EgressPolicyRepo) IsTerminalInUseReturns(result1 bool, result2 error) {
	fake.IsTerminalInUseStub = nil
	fake.isTerminalInUseReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyRepo) IsTerminalInUseReturnsOnCall(i int, result1 bool, result2 error) {
	fake.IsTerminalInUseStub = nil
	if fake.isTerminalInUseReturnsOnCall == nil {
		fake.isTerminalInUseReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.isTerminalInUseReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyRepo) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createAppMutex.RLock()
	defer fake.createAppMutex.RUnlock()
	fake.createIPRangeMutex.RLock()
	defer fake.createIPRangeMutex.RUnlock()
	fake.createEgressPolicyMutex.RLock()
	defer fake.createEgressPolicyMutex.RUnlock()
	fake.createSpaceMutex.RLock()
	defer fake.createSpaceMutex.RUnlock()
	fake.getTerminalByAppGUIDMutex.RLock()
	defer fake.getTerminalByAppGUIDMutex.RUnlock()
	fake.getTerminalBySpaceGUIDMutex.RLock()
	defer fake.getTerminalBySpaceGUIDMutex.RUnlock()
	fake.getAllPoliciesMutex.RLock()
	defer fake.getAllPoliciesMutex.RUnlock()
	fake.getByGuidsMutex.RLock()
	defer fake.getByGuidsMutex.RUnlock()
	fake.getIDCollectionsByEgressPolicyMutex.RLock()
	defer fake.getIDCollectionsByEgressPolicyMutex.RUnlock()
	fake.deleteEgressPolicyMutex.RLock()
	defer fake.deleteEgressPolicyMutex.RUnlock()
	fake.deleteIPRangeMutex.RLock()
	defer fake.deleteIPRangeMutex.RUnlock()
	fake.deleteAppMutex.RLock()
	defer fake.deleteAppMutex.RUnlock()
	fake.deleteSpaceMutex.RLock()
	defer fake.deleteSpaceMutex.RUnlock()
	fake.isTerminalInUseMutex.RLock()
	defer fake.isTerminalInUseMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *EgressPolicyRepo) recordInvocation(key string, args []interface{}) {
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
