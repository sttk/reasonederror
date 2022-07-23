package reasonederror

import (
	"container/list"
	"github.com/stretchr/testify/assert"
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

	assert.Nil(t, syncHandlers.head)
	assert.Nil(t, syncHandlers.last)
}

func TestAddSyncHandler_oneHandler(t *testing.T) {
	clearHandlers()
	defer clearHandlers()

	AddSyncHandler(func(re ReasonedError, dttm time.Time) {})

	assert.NotNil(t, syncHandlers.head)
	assert.Equal(t, syncHandlers.head, syncHandlers.last)
	assert.NotNil(t, syncHandlers.head.handler)

	typ := reflect.TypeOf(syncHandlers.head.handler)
	assert.Equal(t, typ.String(), "func(reasonederror.ReasonedError, time.Time)")

	assert.Nil(t, syncHandlers.last.next)
}

func TestAddSyncHandler_twoHandlers(t *testing.T) {
	clearHandlers()
	defer clearHandlers()

	AddSyncHandler(func(re ReasonedError, dttm time.Time) {})
	AddSyncHandler(func(re ReasonedError, dttm time.Time) {})

	assert.NotNil(t, syncHandlers.head)
	assert.NotEqual(t, syncHandlers.head, syncHandlers.last)
	assert.NotNil(t, syncHandlers.head.handler)

	typ := reflect.TypeOf(syncHandlers.head.handler)
	assert.Equal(t, typ.String(), "func(reasonederror.ReasonedError, time.Time)")

	assert.NotNil(t, syncHandlers.head.next)

	typ = reflect.TypeOf(syncHandlers.head.next.handler)
	assert.Equal(t, typ.String(), "func(reasonederror.ReasonedError, time.Time)")

	assert.Nil(t, syncHandlers.last.next)
}

func TestAddAsyncHandler_zeroHandler(t *testing.T) {
	clearHandlers()
	defer clearHandlers()

	assert.Nil(t, asyncHandlers.head)
	assert.Nil(t, asyncHandlers.last)
}

func TestAddAsyncHandler_oneHandler(t *testing.T) {
	clearHandlers()
	defer clearHandlers()

	AddAsyncHandler(func(re ReasonedError, dttm time.Time) {})

	assert.NotNil(t, asyncHandlers.head)
	assert.Equal(t, asyncHandlers.head, asyncHandlers.last)
	assert.NotNil(t, asyncHandlers.head.handler)

	typ := reflect.TypeOf(asyncHandlers.head.handler)
	assert.Equal(t, typ.String(), "func(reasonederror.ReasonedError, time.Time)")

	assert.Nil(t, asyncHandlers.last.next)
}

func TestAddAsyncHandler_twoHandlers(t *testing.T) {
	clearHandlers()
	defer clearHandlers()

	AddAsyncHandler(func(re ReasonedError, dttm time.Time) {})
	AddAsyncHandler(func(re ReasonedError, dttm time.Time) {})

	assert.NotNil(t, asyncHandlers.head)
	assert.NotEqual(t, asyncHandlers.head, asyncHandlers.last)
	assert.NotNil(t, asyncHandlers.head.handler)

	typ := reflect.TypeOf(asyncHandlers.head.handler)
	assert.Equal(t, typ.String(), "func(reasonederror.ReasonedError, time.Time)")

	assert.NotNil(t, asyncHandlers.head.next)

	typ = reflect.TypeOf(asyncHandlers.head.next.handler)
	assert.Equal(t, typ.String(), "func(reasonederror.ReasonedError, time.Time)")

	assert.Nil(t, asyncHandlers.last.next)
}

func TestFixConfiguration(t *testing.T) {
	clearHandlers()
	defer clearHandlers()

	AddSyncHandler(func(re ReasonedError, dttm time.Time) {})
	AddAsyncHandler(func(re ReasonedError, dttm time.Time) {})

	assert.NotNil(t, syncHandlers.head)
	assert.NotNil(t, syncHandlers.head, syncHandlers.last)
	assert.NotNil(t, syncHandlers.head.handler)

	typ := reflect.TypeOf(syncHandlers.head.handler)
	assert.Equal(t, typ.String(), "func(reasonederror.ReasonedError, time.Time)")

	assert.Nil(t, syncHandlers.last.next)

	assert.NotNil(t, asyncHandlers.head)
	assert.NotNil(t, asyncHandlers.head, asyncHandlers.last)
	assert.NotNil(t, asyncHandlers.head.handler)

	typ = reflect.TypeOf(asyncHandlers.head.handler)
	assert.Equal(t, typ.String(), "func(reasonederror.ReasonedError, time.Time)")

	assert.Nil(t, asyncHandlers.last.next)

	FixConfiguration()

	AddSyncHandler(func(re ReasonedError, dttm time.Time) {})
	AddAsyncHandler(func(re ReasonedError, dttm time.Time) {})

	assert.NotNil(t, syncHandlers.head)
	assert.NotNil(t, syncHandlers.head, asyncHandlers.last)
	assert.NotNil(t, syncHandlers.head.handler)

	typ = reflect.TypeOf(syncHandlers.head.handler)
	assert.Equal(t, typ.String(), "func(reasonederror.ReasonedError, time.Time)")

	assert.Nil(t, syncHandlers.last.next)

	assert.NotNil(t, asyncHandlers.head)
	assert.NotNil(t, asyncHandlers.head, asyncHandlers.last)
	assert.NotNil(t, asyncHandlers.head.handler)

	typ = reflect.TypeOf(asyncHandlers.head.handler)
	assert.Equal(t, typ.String(), "func(reasonederror.ReasonedError, time.Time)")

	assert.Nil(t, asyncHandlers.last.next)
}

func TestNotify_withNoHandler(t *testing.T) {
	clearHandlers()
	defer clearHandlers()

	By(FailToDoSomething{})

	assert.False(t, isFixed)

	FixConfiguration()

	By(FailToDoSomething{})

	assert.True(t, isFixed)
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

	assert.False(t, isFixed)
	assert.Equal(t, syncLogs.Len(), 0)
	assert.Equal(t, asyncLogs.Len(), 0)

	FixConfiguration()

	By(FailToDoSomething{})

	assert.True(t, isFixed)
	assert.Equal(t, syncLogs.Len(), 2)
	assert.Equal(t, syncLogs.Front().Value, "FailToDoSomething-1")
	assert.Equal(t, syncLogs.Front().Next().Value, "FailToDoSomething-2")

	time.Sleep(100 * time.Millisecond)

	assert.Equal(t, asyncLogs.Len(), 1)
	assert.Equal(t, asyncLogs.Front().Value, "FailToDoSomething-3")
}
