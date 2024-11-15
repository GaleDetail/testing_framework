package testframework

import (
	"reflect"
)

func (test *Test) Equal(expected, actual interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		test.Helper()
		test.Errorf("Expected: %v (%T), got: %v (%T)", expected, expected, actual, actual)
	}
}

func (test *Test) NotEqual(expected, actual interface{}) {
	if reflect.DeepEqual(expected, actual) {
		test.Helper()
		test.Errorf("Did not expect: %v (%T), but got: %v (%T)", expected, expected, actual, actual)
	}
}

func (test *Test) True(value bool) {
	if !value {
		test.Helper()
		test.Errorf("Expected true, but got false")
	}
}

func (test *Test) False(value bool) {
	if value {
		test.Helper()
		test.Errorf("Expected false, but got true")
	}
}
