package testframework

import (
	"testing"
)

type Test struct {
	*testing.T
}

func New(t *testing.T) *Test {
	return &Test{T: t}
}
