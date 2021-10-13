package reasonederror_test

import (
	"errors"
	"fmt"
	"github.com/sttk-go/reasonederror"
	"testing"
	"time"
)

type /* Error */ (
	FailToDoSomething struct {
		Param1, Param2, Param3, Param4, Param5  string
		Param6, Param7, Param8, Param9, Param10 int
	}
)

func BenchmarkBy(b *testing.B) {
	var re reasonederror.ReasonedError
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		re = reasonederror.By(FailToDoSomething{
			Param1:  "ABC",
			Param2:  "def",
			Param3:  "ghi",
			Param4:  "jkl",
			Param5:  "mno",
			Param6:  123,
			Param7:  456,
			Param8:  789,
			Param9:  987,
			Param10: 654,
		})
	}

	b.StopTimer()
	_ = fmt.Sprintf("%v\n", re)
}

func BenchmarkBy_withCause(b *testing.B) {
	var re reasonederror.ReasonedError
	cause := errors.New("Causal error")
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		re = reasonederror.By(FailToDoSomething{
			Param1:  "ABC",
			Param2:  "def",
			Param3:  "ghi",
			Param4:  "jkl",
			Param5:  "mno",
			Param6:  123,
			Param7:  456,
			Param8:  789,
			Param9:  987,
			Param10: 654,
		}, cause)
	}

	b.StopTimer()
	_ = fmt.Sprintf("%v\n", re)
}

func BenchmarkBy_withNotification(b *testing.B) {
	reasonederror.AddSyncHandler(func(re reasonederror.ReasonedError, dttm time.Time) {})
	reasonederror.AddAsyncHandler(func(re reasonederror.ReasonedError, dttm time.Time) {})
	reasonederror.FixConfiguration()

	var re reasonederror.ReasonedError
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		re = reasonederror.By(FailToDoSomething{
			Param1:  "ABC",
			Param2:  "def",
			Param3:  "ghi",
			Param4:  "jkl",
			Param5:  "mno",
			Param6:  123,
			Param7:  456,
			Param8:  789,
			Param9:  987,
			Param10: 654,
		})
	}

	b.StopTimer()
	_ = fmt.Sprintf("%v\n", re)
}

func BenchmarkReasonedError_Reason(b *testing.B) {
	re := reasonederror.By(FailToDoSomething{
		Param1:  "ABC",
		Param2:  "def",
		Param3:  "ghi",
		Param4:  "jkl",
		Param5:  "mno",
		Param6:  123,
		Param7:  456,
		Param8:  789,
		Param9:  987,
		Param10: 654,
	})

	b.StartTimer()

	ok := false
	for i := 0; i < b.N; i++ {
		switch re.Reason().(type) {
		case FailToDoSomething, *FailToDoSomething:
			ok = true
		}
	}

	b.StopTimer()
	_ = fmt.Sprintf("%v %t\n", re, ok)
}

func BenchmarkReasonedError_ReasonName(b *testing.B) {
	re := reasonederror.By(FailToDoSomething{
		Param1:  "ABC",
		Param2:  "def",
		Param3:  "ghi",
		Param4:  "jkl",
		Param5:  "mno",
		Param6:  123,
		Param7:  456,
		Param8:  789,
		Param9:  987,
		Param10: 654,
	})

	var reason string
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		reason = re.ReasonName()
	}

	b.StopTimer()
	_ = fmt.Sprintf("%v %s\n", re, reason)
}

func BenchmarkReasonedError_Cause(b *testing.B) {
	cause := errors.New("Causal error")
	re := reasonederror.By(FailToDoSomething{
		Param1:  "ABC",
		Param2:  "def",
		Param3:  "ghi",
		Param4:  "jkl",
		Param5:  "mno",
		Param6:  123,
		Param7:  456,
		Param8:  789,
		Param9:  987,
		Param10: 654,
	}, cause)

	var err error
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		err = re.Cause()
	}

	b.StopTimer()
	_ = fmt.Sprintf("%v %t\n", re, err)
}

func BenchmarkReasonedError_Situation(b *testing.B) {
	re := reasonederror.By(FailToDoSomething{
		Param1:  "ABC",
		Param2:  "def",
		Param3:  "ghi",
		Param4:  "jkl",
		Param5:  "mno",
		Param6:  123,
		Param7:  456,
		Param8:  789,
		Param9:  987,
		Param10: 654,
	})

	var m map[string]interface{}

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		m = re.Situation()
	}

	b.StopTimer()
	_ = fmt.Sprintf("%v %v\n", re, m)
}

func BenchmarkReasonedError_SituationValue(b *testing.B) {
	re := reasonederror.By(FailToDoSomething{
		Param1:  "ABC",
		Param2:  "def",
		Param3:  "ghi",
		Param4:  "jkl",
		Param5:  "mno",
		Param6:  123,
		Param7:  456,
		Param8:  789,
		Param9:  987,
		Param10: 654,
	})

	var s string
	var n int

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		s = re.SituationValue("Param1").(string)
		s = re.SituationValue("Param2").(string)
		s = re.SituationValue("Param3").(string)
		s = re.SituationValue("Param4").(string)
		s = re.SituationValue("Param5").(string)
		n = re.SituationValue("Param6").(int)
		n = re.SituationValue("Param7").(int)
		n = re.SituationValue("Param8").(int)
		n = re.SituationValue("Param9").(int)
		n = re.SituationValue("Param10").(int)
	}

	b.StopTimer()
	_ = fmt.Sprintf("%v %s %d\n", re, s, n)
}
