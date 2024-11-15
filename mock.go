package testframework

type MockFunc struct {
	Calls int
	Args  [][]interface{}
	Func  func(args ...interface{}) interface{}
}

func (m *MockFunc) Call(args ...interface{}) interface{} {
	m.Calls++
	m.Args = append(m.Args, args)
	if m.Func != nil {
		return m.Func(args...)
	}
	return nil
}
