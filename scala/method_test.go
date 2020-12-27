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

func TestMethodParams(t *testing.T) {
	expected := `
def someMethod(param1: String, param2: Int)
`
	method := Method("someMethod")
	method.Param("param1", "String")
	method.Param("param2", "Int")
	assertCode(t, method, expected)
}

func TestMethodParamPerLine(t *testing.T) {
	expected := `
def someMethod(
  param1: String,
  param2: Int
)
`
	method := Method("someMethod").ParamPerLine()
	method.Param("param1", "String")
	method.Param("param2", "Int")
	assertCode(t, method, expected)
}

func TestMethodWithBody(t *testing.T) {
	expected := `
def someMethod() = {
}
`
	method := Method("someMethod")
	method.Define().Scope()
	assertCode(t, method, expected)
}

func TestMethodPrivate(t *testing.T) {
	assertCode(t, Method("someMethod").Private(), `private def someMethod()`)
}
