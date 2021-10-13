// Copyright (C) 2021 Takayuki Sato. All Rights Reserved.
// This program is free software under MIT License.
// See the file LICENSE in this distribution for more details.

/*
Package github.com/sttk-go/reasonederror is for error processes in Go
program.

This package defines ReasonedError structure type which is an error type and
takes a structure type value or pointer as a reason for an error.
The type of this reason is any, therefore this structure can have any fields
according to a situation when an error is caused.

Error with a reason

ReasonedError in this package is an error type with a reason.
A reason is defined by a structure type.
The name of the structure type represents what the reason is.
Since a type is always unique in a Go program, a reason for an error can
be identified by the structure type.

By defining an structure type for a reason in a package which causes a
ReasonedError, the package can be identified with the structure type
because a package of a structure type can be solved with Go reflection.

In addition, since a reason can be any structure type, a reason have some
fields.
These fields would help to know a situation when an error was caused.
The values of these fields can be obtained with Situation or SituationValue
methods.

How to create a ReasonedError

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

To create a ReasonedError, By function is used.
This function returns a value of a ReasonedError.
A way to create ReasonedError is as follows:

    return reasonederror.By(FailToDoSomething{})

If a structure type for a reason has fields, set the field values:

    return reasonederror.By(FailToDoSomethingWithParams{
        Param1: "abc",
        Param2: 123,
    })

If there is a causal error, pass it to By function (ReasonedError supports
Unwrap method):

    err, _ := ...

    return reasonederror.By(FailToDoSomethingWithParams{
        Param1: "abc",
        Param2: 123,
    }, err)

How to evaluate the reason in a ReasonedError

ReasonedError has the method Reason() which return the reason structure value.
Therefore the reason can evaluate by using type switch.

    re := reasonederror.By(FailToDoSomething{})

    switch re.(type) {
    case FailToDoSomething, *FailToDoSomething:
        ...
    }

Creation notification

By registering creation handlers with AddSyncHandler or AddAsyncHandler, these
handlers are notified whenever ReasonedErrors are created with By function.

AddSyncHandler registers a handler which is executed synchronously when a
ReasonedError is created, and AddAsyncHandler registers a handler which is
executed asynchronously.
These functions are effective only before calling FixConfiguration function.
Therefore creation handlers should be registered in start-up process of an
application.

    reasonederror.AddSyncHandler(func(re reasonederror.ReasonedError, dttm time.Time) {
        // (1)
    })
    reasonederror.AddAsyncHandler(func(re reasonederror.ReasonedError, dttm time.Time) {
        // (2)
    })
    reasonederror.FixConfiguration()  // fixes configuration to disable to add more handlers.

Whenever a ReasonedError is created with By function, these registered
handlers are called.
The (1) handler is executed synchronously, and (2) is executed asynchronously
in another goroutine.
*/
package reasonederror
