// Copyright (C) 2021-2023 Takayuki Sato. All Rights Reserved.
// This program is free software under MIT License.
// See the file LICENSE in this distribution for more details.

/*
Package github.com/sttk-go/reasonederror is for error processes in Go
program.

This package defines Err structure type which is an error type and
takes a structure type value or pointer as a reason for an error.
The type of this reason is any, therefore this structure can have any fields
according to a situation when an error is caused.

# Error with a reason

Err in this package is an error type with a reason.
A reason is defined by a structure type.
The name of the structure type represents what the reason is.
Since a type is always unique in a Go program, a reason for an error can
be identified by the structure type.

By defining an structure type for a reason in a package which causes an
Err, the package can be identified with the structure type
because a package of a structure type can be solved with Go reflection.

In addition, since a reason can be any structure type, a reason have some
fields.
These fields would help to know a situation when an error was caused.
The values of these fields can be obtained with Situation or Get methods.

# How to create an Err

First, defines struct types which represent reasons of errors.
It is desirable that these structure types are defined in the package in which
errors for the reason are caused, because it makes possible to solve the
causing package from a reason structure type with Go reflection.

	// Defines error reasons by structure type
	type (
	    FailToDoSomething struct {}
	    FailToDoSomethingWithParams struct {
	        Param1 string
	        Param2 int
	    }
	    ...
	)

To create an Err, By function is used.
This function returns a value of an Err.
A way to create an Err is as follows:

	return reasonederror.NewErr(FailToDoSomething{})

If a structure type for a reason has fields, set the field values:

	return reasonederror.NewErr(FailToDoSomethingWithParams{
	    Param1: "abc",
	    Param2: 123,
	})

If there is a causal error, pass it to By function (Err supports
Unwrap method):

	err, _ := ...

	return reasonederror.NewErr(FailToDoSomethingWithParams{
	    Param1: "abc",
	    Param2: 123,
	}, err)

# How to evaluate the reason in an Err

Err has the method Reason() which return the reason structure value.
Therefore the reason can evaluate by using type switch.

	re := reasonederror.NewErr(FailToDoSomething{})

	switch re.(type) {
	case FailToDoSomething, *FailToDoSomething:
	    ...
	}

# Error notification

By registering handlers with AddSyncErrHandler or AddAsyncErrHandler, these
handlers are notified whenever Err(s) are created with By function.

AddSyncErrHandler registers a handler which is executed synchronously when an
Err is created, and AddAsyncErrHandler registers a handler which is
executed asynchronously.
These functions are effective only before calling FixErrCfgs function.
Therefore creation handlers should be registered in start-up process of an
application.

	reasonederror.AddSyncErrHandler(func(err reasonederror.Err, occ reasonederror.ErrOccasion) {
	    // (1)
	})
	reasonederror.AddAsyncErrHandler(func(err reasonederror.Err, occ reasonederror.ErrOccasion) {
	    // (2)
	})
	reasonederror.FixErrCfgs()  // fixes configuration to disable to add more handlers.

Whenever an Err is created with NewErr function, these registered
handlers are called.
The (1) handler is executed synchronously, and (2) is executed asynchronously
in another goroutine.
*/
package reasonederror
