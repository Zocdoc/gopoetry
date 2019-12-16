package ruby

import "testing"

func TestMethodSimple(t *testing.T) {
	expected := `
def my_method()
end
`
	assertCode(t, Method("my_method"), expected)
}

func TestMethodParams(t *testing.T) {
	method := Method("my_method")
	method.Arg("arg1")
	method.KeywordArg("named_arg2")
	expected := `
def my_method(arg1, named_arg2:)
end
`
	assertCode(t, method, expected)
}

func TestMethodParamPerLine(t *testing.T) {
	method := Method("my_method").ParamPerLine()
	method.Arg("arg1")
	method.KeywordArg("named_arg2")
	expected := `
def my_method(
  arg1,
  named_arg2:
)
end
`
	assertCode(t, method, expected)
}

func TestMethodDefinition(t *testing.T) {
	method := Method("my_method")
	method_ := method.Body()
	method_.AddLn("line1()")
	method_.AddLn("line2()")

	expected := `
def my_method()
  line1()
  line2()
end
`
	assertCode(t, method, expected)
}