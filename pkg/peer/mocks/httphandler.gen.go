// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"sync"

	"github.com/trustbloc/sidetree-core-go/pkg/restapi/common"
)

type HTTPHandler struct {
	PathStub        func() string
	pathMutex       sync.RWMutex
	pathArgsForCall []struct{}
	pathReturns     struct {
		result1 string
	}
	pathReturnsOnCall map[int]struct {
		result1 string
	}
	MethodStub        func() string
	methodMutex       sync.RWMutex
	methodArgsForCall []struct{}
	methodReturns     struct {
		result1 string
	}
	methodReturnsOnCall map[int]struct {
		result1 string
	}
	HandlerStub        func() common.HTTPRequestHandler
	handlerMutex       sync.RWMutex
	handlerArgsForCall []struct{}
	handlerReturns     struct {
		result1 common.HTTPRequestHandler
	}
	handlerReturnsOnCall map[int]struct {
		result1 common.HTTPRequestHandler
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *HTTPHandler) Path() string {
	fake.pathMutex.Lock()
	ret, specificReturn := fake.pathReturnsOnCall[len(fake.pathArgsForCall)]
	fake.pathArgsForCall = append(fake.pathArgsForCall, struct{}{})
	fake.recordInvocation("Path", []interface{}{})
	fake.pathMutex.Unlock()
	if fake.PathStub != nil {
		return fake.PathStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.pathReturns.result1
}

func (fake *HTTPHandler) PathCallCount() int {
	fake.pathMutex.RLock()
	defer fake.pathMutex.RUnlock()
	return len(fake.pathArgsForCall)
}

func (fake *HTTPHandler) PathReturns(result1 string) {
	fake.PathStub = nil
	fake.pathReturns = struct {
		result1 string
	}{result1}
}

func (fake *HTTPHandler) PathReturnsOnCall(i int, result1 string) {
	fake.PathStub = nil
	if fake.pathReturnsOnCall == nil {
		fake.pathReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.pathReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *HTTPHandler) Method() string {
	fake.methodMutex.Lock()
	ret, specificReturn := fake.methodReturnsOnCall[len(fake.methodArgsForCall)]
	fake.methodArgsForCall = append(fake.methodArgsForCall, struct{}{})
	fake.recordInvocation("Method", []interface{}{})
	fake.methodMutex.Unlock()
	if fake.MethodStub != nil {
		return fake.MethodStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.methodReturns.result1
}

func (fake *HTTPHandler) MethodCallCount() int {
	fake.methodMutex.RLock()
	defer fake.methodMutex.RUnlock()
	return len(fake.methodArgsForCall)
}

func (fake *HTTPHandler) MethodReturns(result1 string) {
	fake.MethodStub = nil
	fake.methodReturns = struct {
		result1 string
	}{result1}
}

func (fake *HTTPHandler) MethodReturnsOnCall(i int, result1 string) {
	fake.MethodStub = nil
	if fake.methodReturnsOnCall == nil {
		fake.methodReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.methodReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *HTTPHandler) Handler() common.HTTPRequestHandler {
	fake.handlerMutex.Lock()
	ret, specificReturn := fake.handlerReturnsOnCall[len(fake.handlerArgsForCall)]
	fake.handlerArgsForCall = append(fake.handlerArgsForCall, struct{}{})
	fake.recordInvocation("Handler", []interface{}{})
	fake.handlerMutex.Unlock()
	if fake.HandlerStub != nil {
		return fake.HandlerStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.handlerReturns.result1
}

func (fake *HTTPHandler) HandlerCallCount() int {
	fake.handlerMutex.RLock()
	defer fake.handlerMutex.RUnlock()
	return len(fake.handlerArgsForCall)
}

func (fake *HTTPHandler) HandlerReturns(result1 common.HTTPRequestHandler) {
	fake.HandlerStub = nil
	fake.handlerReturns = struct {
		result1 common.HTTPRequestHandler
	}{result1}
}

func (fake *HTTPHandler) HandlerReturnsOnCall(i int, result1 common.HTTPRequestHandler) {
	fake.HandlerStub = nil
	if fake.handlerReturnsOnCall == nil {
		fake.handlerReturnsOnCall = make(map[int]struct {
			result1 common.HTTPRequestHandler
		})
	}
	fake.handlerReturnsOnCall[i] = struct {
		result1 common.HTTPRequestHandler
	}{result1}
}

func (fake *HTTPHandler) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.pathMutex.RLock()
	defer fake.pathMutex.RUnlock()
	fake.methodMutex.RLock()
	defer fake.methodMutex.RUnlock()
	fake.handlerMutex.RLock()
	defer fake.handlerMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *HTTPHandler) recordInvocation(key string, args []interface{}) {
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

var _ common.HTTPHandler = new(HTTPHandler)
