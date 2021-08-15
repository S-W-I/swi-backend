package utility

import (
	"log"
	"runtime/debug"
	"testing"
)



func ValidateError(t *testing.T, err error) {
	if err != nil {
		log.Fatal(err)
		t.FailNow()
	}
}

func ValidateErrorExistence(t *testing.T, err error) {
	if err == nil {
		// log.Fatal(err)
		t.FailNow()
	}
}

func Assert(t *testing.T, expr bool) {
	if !expr {
		debug.PrintStack()
		t.FailNow()
	}
}
