// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"sync"

	"github.com/chrusty/go-grpc-service/pkg/interfaces"
)

type FakeLogger struct {
	DebugfStub        func(string, ...interface{})
	debugfMutex       sync.RWMutex
	debugfArgsForCall []struct {
		arg1 string
		arg2 []interface{}
	}
	ErrorfStub        func(string, ...interface{})
	errorfMutex       sync.RWMutex
	errorfArgsForCall []struct {
		arg1 string
		arg2 []interface{}
	}
	FatalfStub        func(string, ...interface{})
	fatalfMutex       sync.RWMutex
	fatalfArgsForCall []struct {
		arg1 string
		arg2 []interface{}
	}
	InfofStub        func(string, ...interface{})
	infofMutex       sync.RWMutex
	infofArgsForCall []struct {
		arg1 string
		arg2 []interface{}
	}
	TracefStub        func(string, ...interface{})
	tracefMutex       sync.RWMutex
	tracefArgsForCall []struct {
		arg1 string
		arg2 []interface{}
	}
	WarnfStub        func(string, ...interface{})
	warnfMutex       sync.RWMutex
	warnfArgsForCall []struct {
		arg1 string
		arg2 []interface{}
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeLogger) Debugf(arg1 string, arg2 ...interface{}) {
	fake.debugfMutex.Lock()
	fake.debugfArgsForCall = append(fake.debugfArgsForCall, struct {
		arg1 string
		arg2 []interface{}
	}{arg1, arg2})
	fake.recordInvocation("Debugf", []interface{}{arg1, arg2})
	fake.debugfMutex.Unlock()
	if fake.DebugfStub != nil {
		fake.DebugfStub(arg1, arg2...)
	}
}

func (fake *FakeLogger) DebugfCallCount() int {
	fake.debugfMutex.RLock()
	defer fake.debugfMutex.RUnlock()
	return len(fake.debugfArgsForCall)
}

func (fake *FakeLogger) DebugfCalls(stub func(string, ...interface{})) {
	fake.debugfMutex.Lock()
	defer fake.debugfMutex.Unlock()
	fake.DebugfStub = stub
}

func (fake *FakeLogger) DebugfArgsForCall(i int) (string, []interface{}) {
	fake.debugfMutex.RLock()
	defer fake.debugfMutex.RUnlock()
	argsForCall := fake.debugfArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeLogger) Errorf(arg1 string, arg2 ...interface{}) {
	fake.errorfMutex.Lock()
	fake.errorfArgsForCall = append(fake.errorfArgsForCall, struct {
		arg1 string
		arg2 []interface{}
	}{arg1, arg2})
	fake.recordInvocation("Errorf", []interface{}{arg1, arg2})
	fake.errorfMutex.Unlock()
	if fake.ErrorfStub != nil {
		fake.ErrorfStub(arg1, arg2...)
	}
}

func (fake *FakeLogger) ErrorfCallCount() int {
	fake.errorfMutex.RLock()
	defer fake.errorfMutex.RUnlock()
	return len(fake.errorfArgsForCall)
}

func (fake *FakeLogger) ErrorfCalls(stub func(string, ...interface{})) {
	fake.errorfMutex.Lock()
	defer fake.errorfMutex.Unlock()
	fake.ErrorfStub = stub
}

func (fake *FakeLogger) ErrorfArgsForCall(i int) (string, []interface{}) {
	fake.errorfMutex.RLock()
	defer fake.errorfMutex.RUnlock()
	argsForCall := fake.errorfArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeLogger) Fatalf(arg1 string, arg2 ...interface{}) {
	fake.fatalfMutex.Lock()
	fake.fatalfArgsForCall = append(fake.fatalfArgsForCall, struct {
		arg1 string
		arg2 []interface{}
	}{arg1, arg2})
	fake.recordInvocation("Fatalf", []interface{}{arg1, arg2})
	fake.fatalfMutex.Unlock()
	if fake.FatalfStub != nil {
		fake.FatalfStub(arg1, arg2...)
	}
}

func (fake *FakeLogger) FatalfCallCount() int {
	fake.fatalfMutex.RLock()
	defer fake.fatalfMutex.RUnlock()
	return len(fake.fatalfArgsForCall)
}

func (fake *FakeLogger) FatalfCalls(stub func(string, ...interface{})) {
	fake.fatalfMutex.Lock()
	defer fake.fatalfMutex.Unlock()
	fake.FatalfStub = stub
}

func (fake *FakeLogger) FatalfArgsForCall(i int) (string, []interface{}) {
	fake.fatalfMutex.RLock()
	defer fake.fatalfMutex.RUnlock()
	argsForCall := fake.fatalfArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeLogger) Infof(arg1 string, arg2 ...interface{}) {
	fake.infofMutex.Lock()
	fake.infofArgsForCall = append(fake.infofArgsForCall, struct {
		arg1 string
		arg2 []interface{}
	}{arg1, arg2})
	fake.recordInvocation("Infof", []interface{}{arg1, arg2})
	fake.infofMutex.Unlock()
	if fake.InfofStub != nil {
		fake.InfofStub(arg1, arg2...)
	}
}

func (fake *FakeLogger) InfofCallCount() int {
	fake.infofMutex.RLock()
	defer fake.infofMutex.RUnlock()
	return len(fake.infofArgsForCall)
}

func (fake *FakeLogger) InfofCalls(stub func(string, ...interface{})) {
	fake.infofMutex.Lock()
	defer fake.infofMutex.Unlock()
	fake.InfofStub = stub
}

func (fake *FakeLogger) InfofArgsForCall(i int) (string, []interface{}) {
	fake.infofMutex.RLock()
	defer fake.infofMutex.RUnlock()
	argsForCall := fake.infofArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeLogger) Tracef(arg1 string, arg2 ...interface{}) {
	fake.tracefMutex.Lock()
	fake.tracefArgsForCall = append(fake.tracefArgsForCall, struct {
		arg1 string
		arg2 []interface{}
	}{arg1, arg2})
	fake.recordInvocation("Tracef", []interface{}{arg1, arg2})
	fake.tracefMutex.Unlock()
	if fake.TracefStub != nil {
		fake.TracefStub(arg1, arg2...)
	}
}

func (fake *FakeLogger) TracefCallCount() int {
	fake.tracefMutex.RLock()
	defer fake.tracefMutex.RUnlock()
	return len(fake.tracefArgsForCall)
}

func (fake *FakeLogger) TracefCalls(stub func(string, ...interface{})) {
	fake.tracefMutex.Lock()
	defer fake.tracefMutex.Unlock()
	fake.TracefStub = stub
}

func (fake *FakeLogger) TracefArgsForCall(i int) (string, []interface{}) {
	fake.tracefMutex.RLock()
	defer fake.tracefMutex.RUnlock()
	argsForCall := fake.tracefArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeLogger) Warnf(arg1 string, arg2 ...interface{}) {
	fake.warnfMutex.Lock()
	fake.warnfArgsForCall = append(fake.warnfArgsForCall, struct {
		arg1 string
		arg2 []interface{}
	}{arg1, arg2})
	fake.recordInvocation("Warnf", []interface{}{arg1, arg2})
	fake.warnfMutex.Unlock()
	if fake.WarnfStub != nil {
		fake.WarnfStub(arg1, arg2...)
	}
}

func (fake *FakeLogger) WarnfCallCount() int {
	fake.warnfMutex.RLock()
	defer fake.warnfMutex.RUnlock()
	return len(fake.warnfArgsForCall)
}

func (fake *FakeLogger) WarnfCalls(stub func(string, ...interface{})) {
	fake.warnfMutex.Lock()
	defer fake.warnfMutex.Unlock()
	fake.WarnfStub = stub
}

func (fake *FakeLogger) WarnfArgsForCall(i int) (string, []interface{}) {
	fake.warnfMutex.RLock()
	defer fake.warnfMutex.RUnlock()
	argsForCall := fake.warnfArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeLogger) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.debugfMutex.RLock()
	defer fake.debugfMutex.RUnlock()
	fake.errorfMutex.RLock()
	defer fake.errorfMutex.RUnlock()
	fake.fatalfMutex.RLock()
	defer fake.fatalfMutex.RUnlock()
	fake.infofMutex.RLock()
	defer fake.infofMutex.RUnlock()
	fake.tracefMutex.RLock()
	defer fake.tracefMutex.RUnlock()
	fake.warnfMutex.RLock()
	defer fake.warnfMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeLogger) recordInvocation(key string, args []interface{}) {
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

var _ interfaces.Logger = new(FakeLogger)
