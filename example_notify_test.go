package reasonederror_test

import (
	"fmt"
	"strconv"
	"time"

	"github.com/sttk/reasonederror"
)

func ExampleAddAsyncErrHandler() {
	reasonederror.AddAsyncErrHandler(func(err reasonederror.Err, occ reasonederror.ErrOccasion) {
		fmt.Println("Asynchronous error handling: " + err.Error())
	})
	reasonederror.FixErrCfgs()

	type FailToDoSomething struct{ Name string }

	reasonederror.NewErr(FailToDoSomething{Name: "abc"})

	// Output:
	// Asynchronous error handling: {reason=FailToDoSomething, Name=abc}

	time.Sleep(100 * time.Millisecond)
	reasonederror.ClearErrHandlers()
}

func ExampleAddSyncErrHandler() {
	reasonederror.AddSyncErrHandler(func(err reasonederror.Err, occ reasonederror.ErrOccasion) {
		fmt.Println("Synchronous error handling: " + err.Error())
	})
	reasonederror.FixErrCfgs()

	type FailToDoSomething struct{ Name string }

	reasonederror.NewErr(FailToDoSomething{Name: "abc"})

	// Output:
	// Synchronous error handling: {reason=FailToDoSomething, Name=abc}

	reasonederror.ClearErrHandlers()
}

func ExampleFixErrCfgs() {
	reasonederror.AddSyncErrHandler(func(err reasonederror.Err, occ reasonederror.ErrOccasion) {
		fmt.Println("This handler is registered at " + occ.File() + ":" +
			strconv.Itoa(occ.Line()))
	})

	reasonederror.FixErrCfgs()

	reasonederror.AddSyncErrHandler(func(err reasonederror.Err, occ reasonederror.ErrOccasion) {
		fmt.Println("This handler is not registered")
	})

	type FailToDoSomething struct{ Name string }

	reasonederror.NewErr(FailToDoSomething{Name: "abc"})

	// Output:
	// This handler is registered at example_notify_test.go:58

	reasonederror.ClearErrHandlers()
}
