package testframework

func (test *Test) Before(setup func()) {
	setup()
}

func (test *Test) After(teardown func()) {
	test.Cleanup(teardown)
}
