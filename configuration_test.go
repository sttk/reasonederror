package reasonederror

import (
	"container/list"
	"reflect"
	"testing"
	"time"
)

type FailToDoSomething struct{}

func clearHandlers() {
	syncHandlers.head = nil
	syncHandlers.last = nil
	asyncHandlers.head = nil
	asyncHandlers.last = nil
	isFixed = false
}

func TestAddSyncHandler_zeroHandler(t *testing.T) {
	clearHandlers()
	defer clearHandlers()

	if syncHandlers.head != nil {
		t.Errorf("syncHandlers.head should be nil but: %v.\n", syncHandlers.head)
	}
	if syncHandlers.last != nil {
		t.Errorf("syncHandlers.last should be nil but: %v.\n", syncHandlers.last)
	}
}

func TestAddSyncHandler_oneHandler(t *testing.T) {
	clearHandlers()
	defer clearHandlers()

	AddSyncHandler(func(re ReasonedError, dttm time.Time) {})

	if syncHandlers.head == nil {
		t.Errorf("syncHandlers.head should not be nil.\n")
	}
	if syncHandlers.head != syncHandlers.last {
		t.Errorf("syncHandlers.head should be equal to .last.\n")
	}
	if syncHandlers.head.handler == nil {
		t.Errorf("syncHandlers.head.handler should not be nil.\n")
	}

	typ := reflect.TypeOf(syncHandlers.head.handler)
	if typ.String() != "func(reasonederror.ReasonedError, time.Time)" {
		t.Errorf("Type of syncHandlers.head.handler: %v.\n", typ)
	}
	if syncHandlers.last.next != nil {
		t.Errorf("syncHandlers.last.next should be nil: %v.\n", syncHandlers.head.next)
	}
}

func TestAddSyncHandler_twoHandlers(t *testing.T) {
	clearHandlers()
	defer clearHandlers()

	AddSyncHandler(func(re ReasonedError, dttm time.Time) {})
	AddSyncHandler(func(re ReasonedError, dttm time.Time) {})

	if syncHandlers.head == nil {
		t.Errorf("syncHandlers.head should not be nil.\n")
	}
	if syncHandlers.head.next != syncHandlers.last {
		t.Errorf("syncHandlers.head.next should be equal to .last.\n")
	}
	if syncHandlers.head.handler == nil {
		t.Errorf("syncHandlers.head.handler should not be nil.\n")
	}

	typ := reflect.TypeOf(syncHandlers.head.handler)
	if typ.String() != "func(reasonederror.ReasonedError, time.Time)" {
		t.Errorf("Type of syncHandlers.head.handler: %v.\n", typ)
	}

	if syncHandlers.head.next.handler == nil {
		t.Errorf("syncHandlers.head.next.handler should not be nil.\n")
	}

	typ = reflect.TypeOf(syncHandlers.head.next.handler)
	if typ.String() != "func(reasonederror.ReasonedError, time.Time)" {
		t.Errorf("Type of syncHandlers.head.next.handler: %v.\n", typ)
	}

	if syncHandlers.last.next != nil {
		t.Errorf("syncHandlers.last.next should be nil: %v.\n", syncHandlers.head.next)
	}
}

func TestAddAsyncHandler_zeroHandler(t *testing.T) {
	clearHandlers()
	defer clearHandlers()

	if asyncHandlers.head != nil {
		t.Errorf("asyncHandlers.head should be nil but: %v.\n", asyncHandlers.head)
	}
	if asyncHandlers.last != nil {
		t.Errorf("asyncHandlers.last should be nil but: %v.\n", asyncHandlers.last)
	}
}

func TestAddAsyncHandler_oneHandler(t *testing.T) {
	clearHandlers()
	defer clearHandlers()

	AddAsyncHandler(func(re ReasonedError, dttm time.Time) {})

	if asyncHandlers.head == nil {
		t.Errorf("asyncHandlers.head should not be nil.\n")
	}
	if asyncHandlers.head != asyncHandlers.last {
		t.Errorf("asyncHandlers.head should be equal to .last.\n")
	}
	if asyncHandlers.head.handler == nil {
		t.Errorf("asyncHandlers.head.handler should not be nil.\n")
	}

	typ := reflect.TypeOf(asyncHandlers.head.handler)
	if typ.String() != "func(reasonederror.ReasonedError, time.Time)" {
		t.Errorf("Type of asyncHandlers.head.handler: %v.\n", typ)
	}

	if asyncHandlers.last.next != nil {
		t.Errorf("asyncHandlers.last.next should be nil, but: %v.\n", asyncHandlers.last.next)
	}
}

func TestAddAsyncHandler_twoHandlers(t *testing.T) {
	clearHandlers()
	defer clearHandlers()

	AddAsyncHandler(func(re ReasonedError, dttm time.Time) {})
	AddAsyncHandler(func(re ReasonedError, dttm time.Time) {})

	if asyncHandlers.head == nil {
		t.Errorf("asyncHandlers.head should not be nil.\n")
	}
	if asyncHandlers.head.next != asyncHandlers.last {
		t.Errorf("asyncHandlers.head.next should be equal to .last.\n")
	}

	if asyncHandlers.head.handler == nil {
		t.Errorf("asyncHandlers.head.handler should not be nil.\n")
	}

	typ := reflect.TypeOf(asyncHandlers.head.handler)
	if typ.String() != "func(reasonederror.ReasonedError, time.Time)" {
		t.Errorf("Type of asyncHandlers.head.handler: %v.\n", typ)
	}

	if asyncHandlers.head.next.handler == nil {
		t.Errorf("asyncHandlers.head.next.handler should not be nil.\n")
	}

	typ = reflect.TypeOf(asyncHandlers.head.next.handler)
	if typ.String() != "func(reasonederror.ReasonedError, time.Time)" {
		t.Errorf("Type of asyncHandlers.head.next.handler: %v.\n", typ)
	}

	if asyncHandlers.last.next != nil {
		t.Errorf("asyncHandlers.last.next should be nil: %v.\n", asyncHandlers.head.next)
	}
}

func TestFixConfiguration(t *testing.T) {
	clearHandlers()
	defer clearHandlers()

	AddSyncHandler(func(re ReasonedError, dttm time.Time) {})
	AddAsyncHandler(func(re ReasonedError, dttm time.Time) {})

	if syncHandlers.head == nil {
		t.Errorf("syncHandlers.head should not be nil.\n")
	}
	if syncHandlers.head != syncHandlers.last {
		t.Errorf("syncHandlers.head should be equal to .last.\n")
	}
	if syncHandlers.head.handler == nil {
		t.Errorf("syncHandlers.head.handler should not be nil.\n")
	}

	typ := reflect.TypeOf(syncHandlers.head.handler)
	if typ.String() != "func(reasonederror.ReasonedError, time.Time)" {
		t.Errorf("Type of syncHandlers.head.handler: %v.\n", typ)
	}

	if syncHandlers.last.next != nil {
		t.Errorf("syncHandlers.last.next should be nil: %v.\n", syncHandlers.head.next)
	}

	if asyncHandlers.head == nil {
		t.Errorf("asyncHandlers.head should not be nil.\n")
	}
	if asyncHandlers.head != asyncHandlers.last {
		t.Errorf("asyncHandlers.head should be equal to .last.\n")
	}
	if asyncHandlers.head.handler == nil {
		t.Errorf("asyncHandlers.head.handler should not be nil.\n")
	}

	typ = reflect.TypeOf(asyncHandlers.head.handler)
	if typ.String() != "func(reasonederror.ReasonedError, time.Time)" {
		t.Errorf("Type of asyncHandlers.head.handler: %v.\n", typ)
	}

	if asyncHandlers.last.next != nil {
		t.Errorf("asyncHandlers.last.next should be nil: %v.\n", asyncHandlers.head.next)
	}

	FixConfiguration()

	AddSyncHandler(func(re ReasonedError, dttm time.Time) {})
	AddAsyncHandler(func(re ReasonedError, dttm time.Time) {})

	if syncHandlers.head == nil {
		t.Errorf("syncHandlers.head should not be nil.\n")
	}
	if syncHandlers.head != syncHandlers.last {
		t.Errorf("syncHandlers.head should be equal to .last.\n")
	}
	if syncHandlers.head.handler == nil {
		t.Errorf("syncHandlers.head.handler should not be nil.\n")
	}

	typ = reflect.TypeOf(syncHandlers.head.handler)
	if typ.String() != "func(reasonederror.ReasonedError, time.Time)" {
		t.Errorf("Type of syncHandlers.head.handler: %v.\n", typ)
	}

	if syncHandlers.last.next != nil {
		t.Errorf("syncHandlers.last.next should be nil: %v.\n", syncHandlers.head.next)
	}

	if asyncHandlers.head == nil {
		t.Errorf("asyncHandlers.head should not be nil.\n")
	}
	if asyncHandlers.head != asyncHandlers.last {
		t.Errorf("asyncHandlers.head should be equal to .last.\n")
	}
	if asyncHandlers.head.handler == nil {
		t.Errorf("asyncHandlers.head.handler should not be nil.\n")
	}

	typ = reflect.TypeOf(asyncHandlers.head.handler)
	if typ.String() != "func(reasonederror.ReasonedError, time.Time)" {
		t.Errorf("Type of asyncHandlers.head.handler: %v.\n", typ)
	}

	if asyncHandlers.last.next != nil {
		t.Errorf("asyncHandlers.last.next should be nil: %v.\n", asyncHandlers.head.next)
	}
}

func TestNotify_withNoHandler(t *testing.T) {
	clearHandlers()
	defer clearHandlers()

	By(FailToDoSomething{})

	if isFixed {
		t.Errorf("isFixed should be false but: %v.\n", isFixed)
	}

	FixConfiguration()

	By(FailToDoSomething{})

	if !isFixed {
		t.Errorf("isFixed should be true but: %v.\n", isFixed)
	}
}

func TestNotify_withHandlers(t *testing.T) {
	clearHandlers()
	defer clearHandlers()

	syncLogs := list.New()
	asyncLogs := list.New()

	AddSyncHandler(func(re ReasonedError, dttm time.Time) {
		syncLogs.PushBack(re.ReasonName() + "-1")
	})
	AddSyncHandler(func(re ReasonedError, dttm time.Time) {
		syncLogs.PushBack(re.ReasonName() + "-2")
	})
	AddAsyncHandler(func(re ReasonedError, dttm time.Time) {
		asyncLogs.PushBack(re.ReasonName() + "-3")
	})

	By(FailToDoSomething{})

	if isFixed {
		t.Errorf("isFixed should be false but: %v.\n", isFixed)
	}
	if syncLogs.Len() != 0 {
		t.Errorf("The size of syncLogs should be zero because not calling FixConfiguration.\n")
	}
	if asyncLogs.Len() != 0 {
		t.Errorf("The size of asyncLogs should be zero because not calling FixConfiguration.\n")
	}

	FixConfiguration()

	By(FailToDoSomething{})

	if !isFixed {
		t.Errorf("isFixed should be true but: %v.\n", isFixed)
	}
	if syncLogs.Len() != 2 {
		t.Errorf("The size of syncLogs should be 2 but: %v.\n", syncLogs.Len())
	}
	if syncLogs.Front().Value != "FailToDoSomething-1" {
		t.Errorf("syncLogs[0]=%v.\n", syncLogs.Front().Value)
	}
	if syncLogs.Front().Next().Value != "FailToDoSomething-2" {
		t.Errorf("syncLogs[1]=%v.\n", syncLogs.Front().Next().Value)
	}

	time.Sleep(100 * time.Millisecond)

	if asyncLogs.Len() != 1 {
		t.Errorf("The size of asyncLogs should be 1 but: %v.\n", asyncLogs.Len())
	}
	if asyncLogs.Front().Value != "FailToDoSomething-3" {
		t.Errorf("asyncLogs[0]=%v.\n", asyncLogs.Front().Value)
	}
}
