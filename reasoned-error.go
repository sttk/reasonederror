// Copyright (C) 2021 Takayuki Sato. All Rights Reserved.
// This program is free software under MIT License.
// See the file LICENSE in this distribution for more details.

package reasonederror

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
)

// ReasonedError is a structure type for an error with a reason.
type ReasonedError struct {
	reason interface{}
	file   string
	line   int
	cause  error
}

// NoError is a structure type for a reason of Ok which is a global of Err and indicates no error..
type NoError struct{}

// Ok is a globak Err value which indicates no error.
var Ok = ReasonedError{reason: NoError{}}

// By function creates a new ReasonedError with a reason by a structure and a
// cause of this error.
// Either a value or a pointer of a structure type is fine for a reason.
func By(reason interface{}, cause ...error) ReasonedError {
	var re ReasonedError
	re.reason = reason

	if len(cause) > 0 {
		re.cause = cause[0]
	}

	_, file, line, ok := runtime.Caller(1)
	if ok {
		re.file = filepath.Base(file)
		re.line = line
	}

	notify(re)

	return re
}

// IsOk method determines whether an Err indicates no error.
func (err ReasonedError) IsOk() bool {
	switch err.reason.(type) {
	case NoError, *NoError:
		return true
	default:
		return false
	}
}

// Reason method returns the reason structure.
func (re ReasonedError) Reason() interface{} {
	return re.reason
}

// ReasonName method returns the name of the reason structure type.
func (re ReasonedError) ReasonName() string {
	t := reflect.TypeOf(re.reason)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	return t.Name()
}

// ReasonPackage method returns the package path of the reason structure type.
func (re ReasonedError) ReasonPackage() string {
	t := reflect.TypeOf(re.reason)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	return t.PkgPath()
}

// Situation method returns a map containing parameters which represent the
// situation when this error is caused.
func (re ReasonedError) Situation() map[string]interface{} {
	v := reflect.ValueOf(re.reason)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	m := map[string]interface{}{}

	t := v.Type()
	n := v.NumField()

	for i := 0; i < n; i++ {
		k := t.Field(i).Name

		f := v.Field(i)
		if f.CanInterface() { // false if field is not public.
			m[k] = f.Interface()
		}
	}

	return m
}

// Situation method returns a parameter value of a specified name, which
// represents the situation when this error is caused.
func (re ReasonedError) SituationValue(name string) interface{} {
	v := reflect.ValueOf(re.reason)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	t := v.Type()
	n := v.NumField()
	for i := 0; i < n; i++ {
		k := t.Field(i).Name
		if k == name {
			f := v.Field(i)
			if f.CanInterface() {
				return f.Interface()
			}
		}
	}

	return nil
}

// FileName method returns the the source file name where this error was
// caused.
func (re ReasonedError) FileName() string {
	return re.file
}

// LineNumber method returns the line number in the source file where this
// error was caused.
func (re ReasonedError) LineNumber() int {
	return re.line
}

// Cause method returns the causal error of this error.
func (re ReasonedError) Cause() error {
	return re.cause
}

// Error method returns a string which expresses this error.
func (re ReasonedError) Error() string {
	v := reflect.ValueOf(re.reason)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	t := v.Type()

	s := "reason=" + t.Name()

	n := v.NumField()
	for i := 0; i < n; i++ {
		k := t.Field(i).Name

		f := v.Field(i)
		if f.CanInterface() {
			s += fmt.Sprintf(", %s=%v", k, f.Interface())
		}
	}

	if re.cause != nil {
		s += ", cause=" + re.cause.Error()
	}

	return s
}

// Unwrap returns an error that this error wraps.
func (re ReasonedError) Unwrap() error {
	return re.cause
}
