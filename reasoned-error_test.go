package reasonederror_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/sttk-go/reasonederror"
	"testing"
)

type /* Error */ (
	InvalidValue struct {
		Value string
	}
	FailToGetValue struct {
		Name string
	}
)

func TestBy_reasonIsValue(t *testing.T) {
	re := reasonederror.By(InvalidValue{Value: "abc"})

	assert.Equal(t, re.Error(), "{reason=InvalidValue, Value=abc}")
	assert.Equal(t, re.FileName(), "reasoned-error_test.go")
	assert.Equal(t, re.LineNumber(), 20)

	switch re.Reason().(type) {
	case InvalidValue:
	default:
		assert.Fail(t, re.Error())
	}

	assert.False(t, re.IsOk())
	assert.Equal(t, re.ReasonName(), "InvalidValue")
	assert.Equal(t, re.ReasonPackage(), "github.com/sttk-go/reasonederror_test")
	assert.Equal(t, re.SituationValue("Value"), "abc")
	assert.Nil(t, re.SituationValue("value"))

	m := re.Situation()
	assert.Equal(t, len(m), 1)
	assert.Equal(t, m["Value"], "abc")
	assert.Nil(t, m["value"])

	assert.Nil(t, re.Cause())
	assert.Nil(t, re.Unwrap())
}

func TestBy_reasonIsPointer(t *testing.T) {
	re := reasonederror.By(&InvalidValue{Value: "abc"})

	assert.Equal(t, re.Error(), "{reason=InvalidValue, Value=abc}")
	assert.Equal(t, re.FileName(), "reasoned-error_test.go")
	assert.Equal(t, re.LineNumber(), 48)

	switch re.Reason().(type) {
	case *InvalidValue:
	default:
		assert.Fail(t, re.Error())
	}

	assert.False(t, re.IsOk())
	assert.Equal(t, re.ReasonName(), "InvalidValue")
	assert.Equal(t, re.ReasonPackage(), "github.com/sttk-go/reasonederror_test")
	assert.Equal(t, re.SituationValue("Value"), "abc")
	assert.Nil(t, re.SituationValue("value"))

	m := re.Situation()
	assert.Equal(t, len(m), 1)
	assert.Equal(t, m["Value"], "abc")
	assert.Nil(t, m["value"])

	assert.Nil(t, re.Cause())
	assert.Nil(t, re.Unwrap())
}

func TestBy_withCause(t *testing.T) {
	cause := errors.New("def")
	re := reasonederror.By(InvalidValue{Value: "abc"}, cause)

	assert.Equal(t, re.Error(), "{reason=InvalidValue, Value=abc, cause=def}")
	assert.Equal(t, re.FileName(), "reasoned-error_test.go")
	assert.Equal(t, re.LineNumber(), 77)

	switch re.Reason().(type) {
	case InvalidValue:
	default:
		assert.Fail(t, re.Error())
	}

	assert.False(t, re.IsOk())
	assert.Equal(t, re.ReasonName(), "InvalidValue")
	assert.Equal(t, re.ReasonPackage(), "github.com/sttk-go/reasonederror_test")
	assert.Equal(t, re.SituationValue("Value"), "abc")
	assert.Nil(t, re.SituationValue("value"))

	m := re.Situation()
	assert.Equal(t, len(m), 1)
	assert.Equal(t, m["Value"], "abc")
	assert.Nil(t, m["value"])

	assert.Equal(t, re.Cause(), cause)
	assert.Equal(t, re.Unwrap(), cause)
	assert.Equal(t, errors.Unwrap(re), cause)
}

func TestBy_causeIsAlsoReasonedError(t *testing.T) {
	cause := reasonederror.By(FailToGetValue{Name: "foo"})
	re := reasonederror.By(InvalidValue{Value: "abc"}, cause)

	assert.Equal(t, re.Error(), "{reason=InvalidValue, Value=abc, cause={reason=FailToGetValue, Name=foo}}")
	assert.Equal(t, re.FileName(), "reasoned-error_test.go")
	assert.Equal(t, re.LineNumber(), 107)

	switch re.Reason().(type) {
	case InvalidValue:
	default:
		assert.Fail(t, re.Error())
	}

	assert.False(t, re.IsOk())
	assert.Equal(t, re.ReasonName(), "InvalidValue")
	assert.Equal(t, re.ReasonPackage(), "github.com/sttk-go/reasonederror_test")
	assert.Equal(t, re.SituationValue("Value"), "abc")
	assert.Equal(t, re.SituationValue("Name"), "foo")
	assert.Nil(t, re.SituationValue("value"))

	m := re.Situation()
	assert.Equal(t, len(m), 2)
	assert.Equal(t, m["Value"], "abc")
	assert.Equal(t, m["Name"], "foo")
	assert.Nil(t, m["value"])

	assert.Equal(t, re.Cause(), cause)
	assert.Equal(t, re.Unwrap(), cause)
	assert.Equal(t, errors.Unwrap(re), cause)
}

func TestOk(t *testing.T) {
	re := reasonederror.Ok

	assert.Equal(t, re.Error(), "{reason=NoError}")
	assert.Equal(t, re.FileName(), "")
	assert.Equal(t, re.LineNumber(), 0)

	switch re.Reason().(type) {
	case reasonederror.NoError:
	default:
		assert.Fail(t, re.Error())
	}

	assert.True(t, re.IsOk())
	assert.Equal(t, re.ReasonName(), "NoError")
	assert.Equal(t, re.ReasonPackage(), "github.com/sttk-go/reasonederror")
	assert.Nil(t, re.SituationValue("Value"))
	assert.Nil(t, re.SituationValue("value"))

	m := re.Situation()
	assert.Equal(t, len(m), 0)

	assert.Nil(t, re.Cause())
	assert.Nil(t, re.Unwrap())
}
