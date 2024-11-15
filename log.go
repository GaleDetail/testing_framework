package testframework

import (
	"fmt"
)

func (test *Test) Log(message string, args ...interface{}) {
	test.Helper()
	test.T.Logf(message, args...)
}

func (test *Test) Debug(message string, args ...interface{}) {
	test.Helper()
	fmt.Printf("DEBUG: "+message+"\n", args...)
}
