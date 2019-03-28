package csharp

import "testing"

func TestMethodVoid(t *testing.T) {
	assertCode(t, Method("MyMethod"), "void MyMethod();")
}

func TestMethodReturns(t *testing.T) {
	assertCode(t, Method("MyMethod").Returns("Result"), "Result MyMethod();")
}

func TestMethodPublic(t *testing.T) {
	assertCode(t, Method("MyMethod").Public(), "public void MyMethod();")
}

func TestMethodParams(t *testing.T) {
	method := Method("MyMethod").Public()
	method.Param("int", "intParam")
	method.Param("string", "stringParam")
	assertCode(t, method, "public void MyMethod(int intParam, string stringParam);")
}

func TestMethodBody(t *testing.T) {
	expected := `
Result MyMethod()
{
}
`
	assertCode(t, Method("MyMethod").Returns("Result").Body(), expected)
}

func TestMethodAttributed(t *testing.T) {
	expected := `
[MyAttribute]
void MyMethod();
`
	assertCode(t, Method("MyMethod").WithAttribute("MyAttribute"), expected)
}
