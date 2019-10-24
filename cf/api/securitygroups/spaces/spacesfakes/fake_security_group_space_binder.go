// Code generated by counterfeiter. DO NOT EDIT.
package spacesfakes

import (
	"sync"

	"code.cloudfoundry.org/cli/cf/api/securitygroups/spaces"
)

type FakeSecurityGroupSpaceBinder struct {
	BindSpaceStub        func(string, string) error
	bindSpaceMutex       sync.RWMutex
	bindSpaceArgsForCall []struct {
		arg1 string
		arg2 string
	}
	bindSpaceReturns struct {
		result1 error
	}
	bindSpaceReturnsOnCall map[int]struct {
		result1 error
	}
	UnbindSpaceStub        func(string, string) error
	unbindSpaceMutex       sync.RWMutex
	unbindSpaceArgsForCall []struct {
		arg1 string
		arg2 string
	}
	unbindSpaceReturns struct {
		result1 error
	}
	unbindSpaceReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSecurityGroupSpaceBinder) BindSpace(arg1 string, arg2 string) error {
	fake.bindSpaceMutex.Lock()
	ret, specificReturn := fake.bindSpaceReturnsOnCall[len(fake.bindSpaceArgsForCall)]
	fake.bindSpaceArgsForCall = append(fake.bindSpaceArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("BindSpace", []interface{}{arg1, arg2})
	fake.bindSpaceMutex.Unlock()
	if fake.BindSpaceStub != nil {
		return fake.BindSpaceStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.bindSpaceReturns
	return fakeReturns.result1
}

func (fake *FakeSecurityGroupSpaceBinder) BindSpaceCallCount() int {
	fake.bindSpaceMutex.RLock()
	defer fake.bindSpaceMutex.RUnlock()
	return len(fake.bindSpaceArgsForCall)
}

func (fake *FakeSecurityGroupSpaceBinder) BindSpaceCalls(stub func(string, string) error) {
	fake.bindSpaceMutex.Lock()
	defer fake.bindSpaceMutex.Unlock()
	fake.BindSpaceStub = stub
}

func (fake *FakeSecurityGroupSpaceBinder) BindSpaceArgsForCall(i int) (string, string) {
	fake.bindSpaceMutex.RLock()
	defer fake.bindSpaceMutex.RUnlock()
	argsForCall := fake.bindSpaceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeSecurityGroupSpaceBinder) BindSpaceReturns(result1 error) {
	fake.bindSpaceMutex.Lock()
	defer fake.bindSpaceMutex.Unlock()
	fake.BindSpaceStub = nil
	fake.bindSpaceReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSecurityGroupSpaceBinder) BindSpaceReturnsOnCall(i int, result1 error) {
	fake.bindSpaceMutex.Lock()
	defer fake.bindSpaceMutex.Unlock()
	fake.BindSpaceStub = nil
	if fake.bindSpaceReturnsOnCall == nil {
		fake.bindSpaceReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.bindSpaceReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSecurityGroupSpaceBinder) UnbindSpace(arg1 string, arg2 string) error {
	fake.unbindSpaceMutex.Lock()
	ret, specificReturn := fake.unbindSpaceReturnsOnCall[len(fake.unbindSpaceArgsForCall)]
	fake.unbindSpaceArgsForCall = append(fake.unbindSpaceArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("UnbindSpace", []interface{}{arg1, arg2})
	fake.unbindSpaceMutex.Unlock()
	if fake.UnbindSpaceStub != nil {
		return fake.UnbindSpaceStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.unbindSpaceReturns
	return fakeReturns.result1
}

func (fake *FakeSecurityGroupSpaceBinder) UnbindSpaceCallCount() int {
	fake.unbindSpaceMutex.RLock()
	defer fake.unbindSpaceMutex.RUnlock()
	return len(fake.unbindSpaceArgsForCall)
}

func (fake *FakeSecurityGroupSpaceBinder) UnbindSpaceCalls(stub func(string, string) error) {
	fake.unbindSpaceMutex.Lock()
	defer fake.unbindSpaceMutex.Unlock()
	fake.UnbindSpaceStub = stub
}

func (fake *FakeSecurityGroupSpaceBinder) UnbindSpaceArgsForCall(i int) (string, string) {
	fake.unbindSpaceMutex.RLock()
	defer fake.unbindSpaceMutex.RUnlock()
	argsForCall := fake.unbindSpaceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeSecurityGroupSpaceBinder) UnbindSpaceReturns(result1 error) {
	fake.unbindSpaceMutex.Lock()
	defer fake.unbindSpaceMutex.Unlock()
	fake.UnbindSpaceStub = nil
	fake.unbindSpaceReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSecurityGroupSpaceBinder) UnbindSpaceReturnsOnCall(i int, result1 error) {
	fake.unbindSpaceMutex.Lock()
	defer fake.unbindSpaceMutex.Unlock()
	fake.UnbindSpaceStub = nil
	if fake.unbindSpaceReturnsOnCall == nil {
		fake.unbindSpaceReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.unbindSpaceReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSecurityGroupSpaceBinder) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.bindSpaceMutex.RLock()
	defer fake.bindSpaceMutex.RUnlock()
	fake.unbindSpaceMutex.RLock()
	defer fake.unbindSpaceMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSecurityGroupSpaceBinder) recordInvocation(key string, args []interface{}) {
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

var _ spaces.SecurityGroupSpaceBinder = new(FakeSecurityGroupSpaceBinder)