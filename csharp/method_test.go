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
	method := Method("MyMethod").Returns("Result")
	method.Body()
	assertCode(t, method, expected)
}

func TestMethodAttributed(t *testing.T) {
	expected := `
[MyAttribute]
void MyMethod();
`
	assertCode(t, Method("MyMethod").WithAttribute("MyAttribute"), expected)
}

func TestConstructor(t *testing.T) {
	expected := `
MyType(string myString)
{
}
`
	ctor := Constructor("MyType")
	ctor.Param("string", "myString")
	ctor.Body()
	assertCode(t, ctor, expected)
}

func TestConstructorWithBase(t *testing.T) {
	expected := `
MyType(string myString, bool myBool) : base(myString)
{
}
`
	ctor := Constructor("MyType")
	ctor.Param("string", "myString")
	ctor.Param("bool", "myBool")
	ctor.WithBase("myString")
	ctor.Body()
	assertCode(t, ctor, expected)
}
