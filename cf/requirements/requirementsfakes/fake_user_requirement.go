// Code generated by counterfeiter. DO NOT EDIT.
package requirementsfakes

import (
	"sync"

	"code.cloudfoundry.org/cli/cf/models"
	"code.cloudfoundry.org/cli/cf/requirements"
)

type FakeUserRequirement struct {
	ExecuteStub        func() error
	executeMutex       sync.RWMutex
	executeArgsForCall []struct {
	}
	executeReturns struct {
		result1 error
	}
	executeReturnsOnCall map[int]struct {
		result1 error
	}
	GetUserStub        func() models.UserFields
	getUserMutex       sync.RWMutex
	getUserArgsForCall []struct {
	}
	getUserReturns struct {
		result1 models.UserFields
	}
	getUserReturnsOnCall map[int]struct {
		result1 models.UserFields
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeUserRequirement) Execute() error {
	fake.executeMutex.Lock()
	ret, specificReturn := fake.executeReturnsOnCall[len(fake.executeArgsForCall)]
	fake.executeArgsForCall = append(fake.executeArgsForCall, struct {
	}{})
	fake.recordInvocation("Execute", []interface{}{})
	fake.executeMutex.Unlock()
	if fake.ExecuteStub != nil {
		return fake.ExecuteStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.executeReturns
	return fakeReturns.result1
}

func (fake *FakeUserRequirement) ExecuteCallCount() int {
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	return len(fake.executeArgsForCall)
}

func (fake *FakeUserRequirement) ExecuteCalls(stub func() error) {
	fake.executeMutex.Lock()
	defer fake.executeMutex.Unlock()
	fake.ExecuteStub = stub
}

func (fake *FakeUserRequirement) ExecuteReturns(result1 error) {
	fake.executeMutex.Lock()
	defer fake.executeMutex.Unlock()
	fake.ExecuteStub = nil
	fake.executeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeUserRequirement) ExecuteReturnsOnCall(i int, result1 error) {
	fake.executeMutex.Lock()
	defer fake.executeMutex.Unlock()
	fake.ExecuteStub = nil
	if fake.executeReturnsOnCall == nil {
		fake.executeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.executeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeUserRequirement) GetUser() models.UserFields {
	fake.getUserMutex.Lock()
	ret, specificReturn := fake.getUserReturnsOnCall[len(fake.getUserArgsForCall)]
	fake.getUserArgsForCall = append(fake.getUserArgsForCall, struct {
	}{})
	fake.recordInvocation("GetUser", []interface{}{})
	fake.getUserMutex.Unlock()
	if fake.GetUserStub != nil {
		return fake.GetUserStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.getUserReturns
	return fakeReturns.result1
}

func (fake *FakeUserRequirement) GetUserCallCount() int {
	fake.getUserMutex.RLock()
	defer fake.getUserMutex.RUnlock()
	return len(fake.getUserArgsForCall)
}

func (fake *FakeUserRequirement) GetUserCalls(stub func() models.UserFields) {
	fake.getUserMutex.Lock()
	defer fake.getUserMutex.Unlock()
	fake.GetUserStub = stub
}

func (fake *FakeUserRequirement) GetUserReturns(result1 models.UserFields) {
	fake.getUserMutex.Lock()
	defer fake.getUserMutex.Unlock()
	fake.GetUserStub = nil
	fake.getUserReturns = struct {
		result1 models.UserFields
	}{result1}
}

func (fake *FakeUserRequirement) GetUserReturnsOnCall(i int, result1 models.UserFields) {
	fake.getUserMutex.Lock()
	defer fake.getUserMutex.Unlock()
	fake.GetUserStub = nil
	if fake.getUserReturnsOnCall == nil {
		fake.getUserReturnsOnCall = make(map[int]struct {
			result1 models.UserFields
		})
	}
	fake.getUserReturnsOnCall[i] = struct {
		result1 models.UserFields
	}{result1}
}

func (fake *FakeUserRequirement) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	fake.getUserMutex.RLock()
	defer fake.getUserMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeUserRequirement) recordInvocation(key string, args []interface{}) {
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

var _ requirements.UserRequirement = new(FakeUserRequirement)