package scala

import "testing"

func TestMethodSimple(t *testing.T) {
	assertCode(t, Method("someMethod"), `def someMethod()`)
}

func TestMethodReturn(t *testing.T) {
	assertCode(t, Method("someMethod").Returns("String"), `def someMethod(): String`)
}

func TestMethodZeroParams(t *testing.T) {
	assertCode(t, Method("someMethod").NoParams(), `def someMethod`)
}

func TestMethodWithBody(t *testing.T) {
	expected := `
def someMethod() = {
}
`
	method := Method("someMethod")
	method.Define().Block()
	assertCode(t, method, expected)
}
