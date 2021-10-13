// Copyright (C) 2021 Takayuki Sato. All Rights Reserved.
// This program is free software under MIT License.
// See the file LICENSE in this distribution for more details.

package reasonederror

import (
	"sync"
	"time"
)

type handlerListElem struct {
	handler func(ReasonedError, time.Time)
	next    *handlerListElem
}

type handlerList struct {
	head *handlerListElem
	last *handlerListElem
}

var (
	syncHandlers  = handlerList{nil, nil}
	asyncHandlers = handlerList{nil, nil}
	mutex         = sync.Mutex{}
	isFixed       = false
)

// Adds a creation handler which is executed synchronously when a ReasonedError
// is created.
// Handlers added with this method are executed in the order of addition.
func AddSyncHandler(handler func(ReasonedError, time.Time)) {
	mutex.Lock()
	defer mutex.Unlock()

	if isFixed {
		return
	}

	last := syncHandlers.last
	syncHandlers.last = &handlerListElem{handler, nil}

	if last != nil {
		last.next = syncHandlers.last
	}

	if syncHandlers.head == nil {
		syncHandlers.head = syncHandlers.last
	}
}

// Adds a creation handler which is executed asynchronously when a
// ReasonedError is created.
func AddAsyncHandler(handler func(ReasonedError, time.Time)) {
	mutex.Lock()
	defer mutex.Unlock()

	if isFixed {
		return
	}

	last := asyncHandlers.last
	asyncHandlers.last = &handlerListElem{handler, nil}

	if last != nil {
		last.next = asyncHandlers.last
	}

	if asyncHandlers.head == nil {
		asyncHandlers.head = asyncHandlers.last
	}
}

// Fixes configuration about behaviors of ReasonedError.
// After calling this function, creation handlers cannot be registered any
// more and the notification becomes effective.
func FixConfiguration() {
	mutex.Lock()
	defer mutex.Unlock()

	isFixed = true
}

func notify(re ReasonedError) {
	if !isFixed {
		return
	}

	if syncHandlers.head == nil && asyncHandlers.head == nil {
		return
	}

	dttm := time.Now()

	for e := syncHandlers.head; e != nil; e = e.next {
		e.handler(re, dttm)
	}

	if asyncHandlers.head != nil {
		go func() {
			for e := asyncHandlers.head; e != nil; e = e.next {
				go e.handler(re, dttm)
			}
		}()
	}
}
