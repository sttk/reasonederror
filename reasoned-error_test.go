package reasonederror_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/sttk-go/reasonederror"
	"testing"
)

type /* Error */ (
	InvalidValue struct {
		Value string
	}
)

func TestBy_reasonIsValue(t *testing.T) {
	re := reasonederror.By(InvalidValue{Value: "abc"})

	assert.Equal(t, re.Error(), "reason=InvalidValue, Value=abc")
	assert.Equal(t, re.FileName(), "reasoned-error_test.go")
	assert.Equal(t, re.LineNumber(), 16)

	switch re.Reason().(type) {
	case InvalidValue:
	default:
		t.Errorf("re.Reason() = %v\n", re.Reason())
	}

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

	assert.Equal(t, re.Error(), "reason=InvalidValue, Value=abc")
	assert.Equal(t, re.FileName(), "reasoned-error_test.go")
	assert.Equal(t, re.LineNumber(), 43)

	switch re.Reason().(type) {
	case *InvalidValue:
	default:
		t.Errorf("re.Reason() = %v\n", re.Reason())
	}

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
