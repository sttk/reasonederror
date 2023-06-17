# [reasonederror][repo-url] [![Go Reference][pkg-dev-img]][pkg-dev-url] [![CI Status][ci-img]][ci-url] [![MIT License][mit-img]][mit-url]


The error processing library for Go.


- [What is this?](#what-is-this)
- [Features](#features)
- [Usage](#usage)
- [Supporting Go versions](#supporting-go-versions)
- [License](#license)


<a name="what-is-this"></a>
## What is this?

`reasonederror` is a library for error processes in Go program.

The main type of this library is `Err`.
This structure type takes any structure value/pointer as a reason of an error.
The type of this reason is any, therefore this can have any fields which represent a situation when an error is caused.


<a name="features"></a>
## Features

This library provides the following features:

- Error with a reason
- Error notification

### Error with a reason

`Err` in this library is an error type with a reason of an error.
A reason is defined by a structure type.
The name of this structure type represents what the reason is.
Since a structure type is always unique in a Go program, a reason by a structure value can identify an error.
This will free you from efforts to implement many error types and their fields and methods for various error situations.

By defining a structure type for a reason in a package which creates its `Err`, the package can be identified with the structure value because an package of an structure type can be solved with Go reflection.

`Err` is created with `NewErr` function.
This function can take a value or pointer of a structure type, and can also take parameters by using fields of the structure.
These parameters help to know a situation when an error is caused, and can get their values with `#Situation` or `#Get` method.

`reasonederror.Ok` is a global value of `Err`, which indicates no error.
Since `NewErr` function returns a `Err` value, not a pointer, it is needed the way to indicate and check no error with a `Err` value. `reasonederror.Ok` just indicates it, and the method `#IsOk` can check it.

### Error notification

`Err` has a function to notify the creation of itself when calling `NewErr` function.
The notifications can be done to synchronous handlers or asynchronous handlers.
A synchronous handler is registered with `AddSyncErrHandler` function and an asynchronous handler is registered with `AddAsyncErrHandler` function.
And then the notification is made possible by calling `FixErrCfgs` function.


<a name="usage"></a>
## Usage

This section explains the usage of functions, structure types, and methods in this library.

### Creates an `Err`

First, imports `reasonederror` package as follows:

```
import "github.com/sttk-go/reasonederror"
```

Next, defines structure types which represent reasons of errors.
It is desirable that these structure types are defined in the package in which errors are caused, because it makes possible to solve the package path of the structure value.

```
    // Defines error reasons by structure types.
    type (
      FailToDoSomething struct {}
      FailToDoSomethingWithParams struct {
        Param1 string,
        Param2 int,
      }
      ...
    )
```

A way to create a `Err` is as follows:

```
  return reasonederror.NewErr(FailToDoSomething{})
```

If a structure type for a reason has fields which help to know a situation when an error is caused, set the field values:

```
  return reasonederror.NewErr(FailToDoSomethingWithParams{
    Param1: "abc",
    Param2: 123,
  })
```

If there is a causal error, pass it to `By` function (`Err` supports `#Unwrap` method):

```
  err, result := ...

  return reasonederror.NewErr(FailToDoSomethingWithParams{
    Param1: "abc",
    Param2: 123,
  }, err)
```

To return `Err` value which indicates no error, `reasonederror.Ok` is used.

```
  return reasonederror.Ok
```

Then, there are two way to check whether a `Err` value indicates no error or not.
One way is as follows:

```
  var re readsonederror.Err
  re = ...
  if re.IsOk() {
    ...
  } else {
    ...
  }
```

And another way is as follows:

```
  var re readsonederror.Err
  re = ...
  switch re.Reason().(type) {
  case nil:
    ...
  default:
    ...
  }
```

### Registers error handlers

By registering error handlers with `AddSyncErrHandler` or `AddAsyncErrHandler`, these handlers are notified whenever `Err`s are created with `NewErr` function.

`AddSyncErrHandler` registers a handler which is executed synchronously when a `Err` is created, and `AddAsyncErrHandler` registers a handler which is executed asynchronously.
These functions are effective only before calling `FixErrCfgs` function.

Error handlers should be registered in start-up process of an application.

```
reasonederror.AddSyncErrHandler(func(err reasonederror.Err, occ reasonederror.ErrOccasion) {
    // (1)
})
reasonederror.AddAsyncErrHandler(func(err reasonederror.Err, occ reasonederror.ErrOccasion) {
    // (2)
})
reasonederror.FixErrCfgs() // fixes configuration to disable to add more handlers.
```

Whenever a `Err` is created with `NewErr` funciton, these registered handlers are called. The (1) handler is executed synchronously, and (2) is executed asynchronously in another goroutine.


<a name="supporting-go-versions"></a>
## Supporting Go versions

This library supports Go 1.13 or later.

### Actual test results for each Go version:

```
% gvm-fav
Now using version go1.13.15
go version go1.13.15 darwin/amd64
ok  	github.com/sttk-go/reasonederror	0.367s	coverage: 100.0% of statements

Now using version go1.14.15
go version go1.14.15 darwin/amd64
ok  	github.com/sttk-go/reasonederror	0.361s	coverage: 100.0% of statements

Now using version go1.15.15
go version go1.15.15 darwin/amd64
ok  	github.com/sttk-go/reasonederror	0.341s	coverage: 100.0% of statements

Now using version go1.16.15
go version go1.16.15 darwin/amd64
ok  	github.com/sttk-go/reasonederror	0.338s	coverage: 100.0% of statements

Now using version go1.17.13
go version go1.17.13 darwin/amd64
ok  	github.com/sttk-go/reasonederror	0.334s	coverage: 100.0% of statements

Now using version go1.18.10
go version go1.18.10 darwin/amd64
ok  	github.com/sttk-go/reasonederror	0.356s	coverage: 100.0% of statements

Now using version go1.19.10
go version go1.19.10 darwin/amd64
ok  	github.com/sttk-go/reasonederror	0.344s	coverage: 100.0% of statements

Now using version go1.20.5
go version go1.20.5 darwin/amd64
ok  	github.com/sttk-go/reasonederror	0.351s	coverage: 100.0% of statements

Back to go1.20.5
Now using version go1.20.5
```

<a name="license"></a>
## License

Copyright (C) 2021-2023 Takayuki Sato

This program is free software under MIT License.
See the file LICENSE in this distribution for more details.

[repo-url]: https://github.com/sttk/reasonederror
[ci-img]: https://github.com/sttk/reasonederror/actions/workflows/go.yml/badge.svg?branch=main
[ci-url]: https://github.com/sttk/reasonederror/actions
[pkg-dev-img]: https://pkg.go.dev/badge/github.com/sttk/reasonederror.svg
[pkg-dev-url]: https://pkg.go.dev/github.com/sttk-go/reasonederror
[mit-img]: https://img.shields.io/badge/license-MIT-green.svg
[mit-url]: https://opensource.org/licenses/MIT
