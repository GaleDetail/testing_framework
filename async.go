package testframework

import (
	"time"
)

func (test *Test) WaitFor(condition func() bool, timeout time.Duration) {
	start := time.Now()
	for {
		if condition() {
			return
		}
		if time.Since(start) > timeout {
			test.Helper()
			test.Fatalf("Condition not met within %v", timeout)
		}
		time.Sleep(10 * time.Millisecond)
	}
}
