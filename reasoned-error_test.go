package reasonederror_test

import (
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

	// t.Logf("re = %v\n", re)

	ex0 := "reason=InvalidValue, Value=abc"
	if re.Error() != ex0 {
		t.Errorf("re.Error() = %v (differ from %v)\n", re.Error(), ex0)
	}

	ex1 := "reasoned-error_test.go"
	if re.File() != ex1 {
		t.Errorf("re.File() = %v (differ from %v)\n", re.File(), ex1)
	}

	ex2 := 15
	if re.Line() != ex2 {
		t.Errorf("re.Line() = %v (differ from %v)\n", re.Line(), ex2)
	}

	switch re.Reason().(type) {
	case InvalidValue:
	default:
		t.Errorf("re.Reason() = %v\n", re.Reason())
	}

	ex3 := "InvalidValue"
	if re.ReasonName() != ex3 {
		t.Errorf("re.ReasonName() = %v (differ from %v)\n", re.ReasonName(), ex3)
	}

	ex4 := "github.com/sttk-go/reasonederror_test"
	if re.ReasonPackage() != ex4 {
		t.Errorf("re.ReasonPackage() = %v (differ from %v)\n", re.ReasonPackage(), ex4)
	}

	m := re.Situation()
	ex5 := 1
	if len(m) != ex5 {
		t.Errorf("re.Situation():len = %v (differ from %v)\n", len(m), ex5)
	}

	ex6 := "abc"
	if m["Value"] != ex6 {
		t.Errorf("re.Situation():[\"Value\"] = %v (differ from %v)\n", m["Value"], ex6)
	}
}

func TestBy_reasonIsPointer(t *testing.T) {
	re := reasonederror.By(&InvalidValue{Value: "abc"})

	//t.Logf("re = %v\n", re)

	ex0 := "reason=InvalidValue, Value=abc"
	if re.Error() != ex0 {
		t.Errorf("re.Error() = %v (differ from %v)\n", re.Error(), ex0)
	}

	ex1 := "reasoned-error_test.go"
	if re.File() != ex1 {
		t.Errorf("re.File() = %v (differ from %v)\n", re.File(), ex1)
	}

	ex2 := 63
	if re.Line() != ex2 {
		t.Errorf("re.Line() = %v (differ from %v)\n", re.Line(), ex2)
	}

	switch re.Reason().(type) {
	case *InvalidValue:
	default:
		t.Errorf("re.Reason() = %v\n", re.Reason())
	}

	ex3 := "InvalidValue"
	if re.ReasonName() != ex3 {
		t.Errorf("re.ReasonName() = %v (differ from %v)\n", re.ReasonName(), ex3)
	}

	ex4 := "github.com/sttk-go/reasonederror_test"
	if re.ReasonPackage() != ex4 {
		t.Errorf("re.ReasonPackage() = %v (differ from %v)\n", re.ReasonPackage(), ex4)
	}

	m := re.Situation()
	ex5 := 1
	if len(m) != ex5 {
		t.Errorf("re.Situation():len = %v (differ from %v)\n", len(m), ex5)
	}
	ex6 := "abc"
	if m["Value"] != ex6 {
		t.Errorf("re.Situation():[\"Value\"] = %v (differ from %v)\n", m["Value"], ex6)
	}

	ex7 := "abc"
	if re.SituationValue("Value") != ex7 {
		t.Errorf("re.SituationValue(\"Value\") = %v (differ from %v)\n", re.SituationValue("Value"), ex7)
	}

	var ex8 interface{} = nil
	if re.SituationValue("Xxx") != ex8 {
		t.Errorf("re.SituationValue(\"Value\") = %v (differ from %v)\n", re.SituationValue("Value"), ex8)
	}
}
