package scala

import "testing"

func TestMethodSimple(t *testing.T) {
	assertCode(t, Def("someMethod"), `def someMethod()`)
}

func TestMethodReturn(t *testing.T) {
	assertCode(t, Def("someMethod").Returns("String"), `def someMethod(): String`)
}

func TestMethodZeroParams(t *testing.T) {
	assertCode(t, Def("someMethod").NoParams(), `def someMethod`)
}

func TestMethodParams(t *testing.T) {
	expected := `def someMethod(param1: String, param2: Int)`
	method :=
		Def("someMethod").AddParams(
			Param("param1", "String"),
			Param("param2", "Int"),
		)
	assertCode(t, method, expected)
}

func TestMethodParamsNil(t *testing.T) {
	expected := `def someMethod(param1: String, param2: Int)`
	method :=
		Def("someMethod").AddParams(
			Param("param1", "String"),
			nil,
			Param("param2", "Int"),
		)
	assertCode(t, method, expected)
}

func TestMethodParamPerLine(t *testing.T) {
	expected := `
def someMethod(
  param1: String,
  param2: Int
)`
	method :=
		Def("someMethod").ParamPerLine().
			Param("param1", "String").
			Param("param2", "Int")
	assertCode(t, method, expected)
}

func TestMethodWithBody(t *testing.T) {
	expected := `
def someMethod() = {
}`
	method := Def("someMethod")
	method.Body()
	assertCode(t, method, expected)
}

func TestMethodPrivate(t *testing.T) {
	assertCode(t, Def("someMethod").Private(), `private def someMethod()`)
}
