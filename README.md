# [sttk-go/reasonederror][repo-url] [![Go Reference][pkg-dev-img]][pkg-dev-url] [![CI Status][ci-img]][ci-url] [![MIT License][mit-img]][mit-url]


The error processing library for Go.


- [What is this?](#what-is-this)
- [Features](#features)
- [Usage](#usage)
- [Supporting Go versions](#supporting-go-versions)
- [License](#license)


<a name="what-is-this"></a>
## What is this?

`github.com/sttk-go/reasonederror` is a library for error processes in Go program.

The central type of this library is `ReasonedError`.
This structure type takes any structure value/pointer as a reason of an error.
The type of this reason is any, therefore this can have any fields which represent a situation when an error is caused.


<a name="features"></a>
## Features

This library provides the following features:

- Error with a reason
- Creation notification

### Error with a reason

`ReasonedError` in this library is an error type with a reason of an error.
A reason is defined by a structure type.
The name of this structure type represents what the reason is.
Since a structure type is always unique in a Go program, a reason by a structure value can identify an error.
This will free you from efforts to implement many error types and their fields and methods for various error situations.

By defining a structure type for a reason in a package which creates its `ReasonedError`, the package can be identified with the structure value because an package of an structure type can be solved with Go reflection.

`ReasonedError` is created with `By` function.
This function can take a value or pointer of a structure type, and can also take parameters by using fields of the structure.
These parameters help to know a situation when an error is caused, and can get their values with `#Situation` or `#SituationValue` method.

`reasonederror.Ok` is a global value of `ReasonedError`, which indicates no error.
Since `By` function returns a `ReasondError` value, not a pointer, it is needed the way to indicate and check no error with a `ReasonedError` value. `reasonederror.Ok` just indicates it, and the method `#IsOk` can check it.

### Creation notification

`ReasonedError` has a function to notify the creation of itself when calling `By` function.
The notifications can be done to synchronous handlers or asynchronous handlers.
A synchronous handler is registered with `AddSyncHandler` function and an asynchronous handler is registered with `AddAsyncHandler` function.
And then the notification is made possible by calling `FixConfiguration` function.


<a name="usage"></a>
## Usage

This section explains the usage of functions, structure types, and methods in this library.

### Creates a `ReasonedError`

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

A way to create a `ReasonedError` is as follows:

```
  return reasonederror.By(FailToDoSomething{})
```

If a structure type for a reason has fields which help to know a situation when an error is caused, set the field values:

```
  return reasonederror.By(FailToDoSomethingWithParams{
    Param1: "abc",
    Param2: 123,
  })
```

If there is a causal error, pass it to `By` function (`ReasonedError` supports `#Unwrap` method):

```
  err, result := ...

  return reasonederror.By(FailToDoSomethingWithParams{
    Param1: "abc",
    Param2: 123,
  }, err)
```

To return `ReasonedError` value which indicates no error, `reasonederror.Ok` is used.

```
  return reasonederror.Ok
```

Then, there are two way to check whether a `ReasonedError` value indicates no error or not.
One way is as follows:

```
  var re readsonederror.ReasoedError
  re = ...
  if re.IsOk() {
    ...
  } else {
    ...
  }
```

And another way is as follows:

```
  var re readsonederror.ReasoedError
  re = ...
  switch re.Reason().(type) {
  case reasonederror.NoError, *reasonederror.NoError:
    ...
  default:
    ...
  }
```

### Registers creation handlers

By registering creation handlers with `AddSyncHandler` or `AddAsyncHandler`, these handlers are notified whenever `ReasonedError`s are created with `By` function.

`AddSyncHandler` registers a handler which is executed synchronously when a `ReasonedError` is created, and `AddAsyncHandler` registers a handler which is executed asynchronously.
These functions are effective only before calling `FixConfiguration` function.

Creation handlers should be registered in start-up process of an application.

```
reasonederror.AddSyncHandler(func(re reasonederror.ReasonedError, dttm time.Time) {
    // (1)
})
reasonederror.AddAsyncHandler(func(re reasonederror.ReasonedError, dttm time.Time) {
    // (2)
})
reasonederror.FixConfiguration() // fixes configuration to disable to add more handlers.
```

Whenever a `ReasonedError` is created with `By` funciton, these registered handlers are called. The (1) handler is executed synchronously, and (2) is executed asynchronously in another goroutine.


<a name="supporting-go-versions"></a>
## Supporting Go versions

This library supports Go 1.13 or later.

### Actually Checked Go versions

- 1.18.4
- 1.17.12
- 1.16.15
- 1.15.15
- 1.14.15
- 1.13.15


<a name="license"></a>
## License

Copyright (C) 2021 Takayuki Sato

This program is free software under MIT License.
See the file LICENSE in this distribution for more details.

[repo-url]: https://github.com/sttk-go/reasonederror
[ci-img]: https://github.com/sttk-go/reasonederror/actions/workflows/go.yml/badge.svg?branch=main
[ci-url]: https://github.com/sttk-go/reasonederror/actions
[pkg-dev-img]: https://pkg.go.dev/badge/github.com/sttk-go/reasonederror.svg
[pkg-dev-url]: https://pkg.go.dev/github.com/sttk-go/reasonederror
[mit-img]: https://img.shields.io/badge/license-MIT-green.svg
[mit-url]: https://opensource.org/licenses/MIT
