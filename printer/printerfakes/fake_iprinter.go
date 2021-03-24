// Code generated by counterfeiter. DO NOT EDIT.
package printerfakes

import (
	"sync"

	"github.com/batchcorp/plumber/printer"
)

type FakeIPrinter struct {
	ErrorStub        func(string)
	errorMutex       sync.RWMutex
	errorArgsForCall []struct {
		arg1 string
	}
	PrintStub        func(string)
	printMutex       sync.RWMutex
	printArgsForCall []struct {
		arg1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeIPrinter) Error(arg1 string) {
	fake.errorMutex.Lock()
	fake.errorArgsForCall = append(fake.errorArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.ErrorStub
	fake.recordInvocation("Error", []interface{}{arg1})
	fake.errorMutex.Unlock()
	if stub != nil {
		fake.ErrorStub(arg1)
	}
}

func (fake *FakeIPrinter) ErrorCallCount() int {
	fake.errorMutex.RLock()
	defer fake.errorMutex.RUnlock()
	return len(fake.errorArgsForCall)
}

func (fake *FakeIPrinter) ErrorCalls(stub func(string)) {
	fake.errorMutex.Lock()
	defer fake.errorMutex.Unlock()
	fake.ErrorStub = stub
}

func (fake *FakeIPrinter) ErrorArgsForCall(i int) string {
	fake.errorMutex.RLock()
	defer fake.errorMutex.RUnlock()
	argsForCall := fake.errorArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeIPrinter) Print(arg1 string) {
	fake.printMutex.Lock()
	fake.printArgsForCall = append(fake.printArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.PrintStub
	fake.recordInvocation("Print", []interface{}{arg1})
	fake.printMutex.Unlock()
	if stub != nil {
		fake.PrintStub(arg1)
	}
}

func (fake *FakeIPrinter) PrintCallCount() int {
	fake.printMutex.RLock()
	defer fake.printMutex.RUnlock()
	return len(fake.printArgsForCall)
}

func (fake *FakeIPrinter) PrintCalls(stub func(string)) {
	fake.printMutex.Lock()
	defer fake.printMutex.Unlock()
	fake.PrintStub = stub
}

func (fake *FakeIPrinter) PrintArgsForCall(i int) string {
	fake.printMutex.RLock()
	defer fake.printMutex.RUnlock()
	argsForCall := fake.printArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeIPrinter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.errorMutex.RLock()
	defer fake.errorMutex.RUnlock()
	fake.printMutex.RLock()
	defer fake.printMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeIPrinter) recordInvocation(key string, args []interface{}) {
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

var _ printer.IPrinter = new(FakeIPrinter)