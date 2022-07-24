package reasonederror_test

import (
	"errors"
	"fmt"
	"github.com/sttk-go/reasonederror"
)

func ExampleBy() {
	type /* Error */ (
		FailToDoSomething           struct{}
		FailToDoSomethingWithParams struct {
			Param1 string
			Param2 int
		}
	)

	// (1) Creates a ReasonedError with no situation parameter.
	re := reasonederror.By(FailToDoSomething{})
	fmt.Printf("(1) %v\n", re)

	// (2) Creates a ReasonedError with situation parameters.
	re = reasonederror.By(FailToDoSomethingWithParams{
		Param1: "ABC",
		Param2: 123,
	})
	fmt.Printf("(2) %v\n", re)

	err := errors.New("Causal error")

	// (3) Creates a ReasonedError with a causal error.
	re = reasonederror.By(FailToDoSomething{}, err)
	fmt.Printf("(3) %v\n", re)

	// (4) Creates a ReasonedError with situation parameters and a causal error.
	re = reasonederror.By(FailToDoSomethingWithParams{
		Param1: "ABC",
		Param2: 123,
	}, err)
	fmt.Printf("(4) %v\n", re)

	// Output:
	// (1) {reason=FailToDoSomething}
	// (2) {reason=FailToDoSomethingWithParams, Param1=ABC, Param2=123}
	// (3) {reason=FailToDoSomething, cause=Causal error}
	// (4) {reason=FailToDoSomethingWithParams, Param1=ABC, Param2=123, cause=Causal error}
}

func ExampleReasonedError_Cause() {
	type FailToDoSomething struct{}

	err := errors.New("Causal error")

	re := reasonederror.By(FailToDoSomething{}, err)
	fmt.Printf("%v\n", re.Cause())

	// Output:
	// Causal error
}

func ExampleReasonedError_Error() {
	type FailToDoSomething struct {
		Param1 string
		Param2 int
	}

	err := errors.New("Causal error")

	re := reasonederror.By(FailToDoSomething{
		Param1: "ABC",
		Param2: 123,
	}, err)
	fmt.Printf("%v\n", re.Error())

	// Output:
	// {reason=FailToDoSomething, Param1=ABC, Param2=123, cause=Causal error}
}

func ExampleReasonedError_FileName() {
	type FailToDoSomething struct{}

	re := reasonederror.By(FailToDoSomething{})
	fmt.Printf("%v\n", re.FileName())

	// Output:
	// example_reasoned-error_test.go
}

func ExampleReasonedError_LineNumber() {
	type FailToDoSomething struct{}

	re := reasonederror.By(FailToDoSomething{})
	fmt.Printf("%v\n", re.LineNumber())

	// Output:
	// 92
}

func ExampleReasonedError_Reason() {
	type FailToDoSomething struct {
		Param1 string
	}

	re := reasonederror.By(FailToDoSomething{Param1: "value1"})
	switch re.Reason().(type) {
	case FailToDoSomething:
		fmt.Println("The reason of the error is: FailToDoSomething")
		reason := re.Reason().(FailToDoSomething)
		fmt.Printf("The value of reason.Param1 is: %v\n", reason.Param1)
	}

	re = reasonederror.By(&FailToDoSomething{Param1: "value1"})
	switch re.Reason().(type) {
	case *FailToDoSomething:
		fmt.Println("The reason of the error is: *FailToDoSomething")
		reason := re.Reason().(*FailToDoSomething)
		fmt.Printf("The value of reason.Param1 is: %v\n", reason.Param1)
	}

	// Output:
	// The reason of the error is: FailToDoSomething
	// The value of reason.Param1 is: value1
	// The reason of the error is: *FailToDoSomething
	// The value of reason.Param1 is: value1
}

func ExampleReasonedError_ReasonName() {
	type FailToDoSomething struct{}

	re := reasonederror.By(FailToDoSomething{})
	fmt.Printf("%v\n", re.ReasonName())

	// Output:
	// FailToDoSomething
}

func ExampleReasonedError_ReasonPackage() {
	type FailToDoSomething struct{}

	re := reasonederror.By(FailToDoSomething{})
	fmt.Printf("%v\n", re.ReasonPackage())

	// Output:
	// github.com/sttk-go/reasonederror_test
}

func ExampleReasonedError_Situation() {
	type FailToDoSomething struct {
		Param1 string
		Param2 int
	}

	re := reasonederror.By(FailToDoSomething{
		Param1: "ABC",
		Param2: 123,
	})
	fmt.Printf("%v\n", re.Situation())

	// Output:
	// map[Param1:ABC Param2:123]
}

func ExampleReasonedError_SituationValue() {
	type FailToDoSomething struct {
		Param1 string
		Param2 int
	}

	re := reasonederror.By(FailToDoSomething{
		Param1: "ABC",
		Param2: 123,
	})
	fmt.Printf("Param1=%v\n", re.SituationValue("Param1"))
	fmt.Printf("Param2=%v\n", re.SituationValue("Param2"))
	fmt.Printf("Param3=%v\n", re.SituationValue("Param3"))

	// Output:
	// Param1=ABC
	// Param2=123
	// Param3=<nil>
}

func ExampleReasonedError_Unwrap() {
	type FailToDoSomething struct{}

	err1 := errors.New("Causal error 1")
	err2 := errors.New("Causal error 2")

	re := reasonederror.By(FailToDoSomething{}, err1)

	fmt.Printf("re.Unwrap() = %v\n", re.Unwrap())
	fmt.Printf("errors.Is(re, err1) = %v\n", errors.Is(re, err1))
	fmt.Printf("errors.Is(re, err2) = %v\n", errors.Is(re, err2))

	// Output:
	// re.Unwrap() = Causal error 1
	// errors.Is(re, err1) = true
	// errors.Is(re, err2) = false
}
