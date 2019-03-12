package cs

import (
	"strings"
	"testing"
	"gotest.tools/assert"
)

func assertCode(t *testing.T, value Writable, expected string) {
	writer := CreateWriter()
	value.WriteCode(&writer)
	code := writer.Code()
	assert.Equal(t, strings.TrimSpace(code), strings.TrimSpace(expected))
}

func TestClassBasic(t *testing.T) {
	expected := `
class MyClass
{
}
`
	assertCode(t, Class("MyClass"), expected)
}

func TestClassModifier(t *testing.T) {
	expected := `
private class MyClass
{
}
`
	assertCode(t, Class("MyClass").Private(), expected)
}

func TestClassMethod(t *testing.T) {
	expected := `
class MyClass
{
    void MyMethod()
    {
    }
}
`
	assertCode(t, Class("MyClass").Members(Method("MyMethod").Body()), expected)
}

func TestClassProperty(t *testing.T) {
	expected := `
class MyClass
{
    Result MyProperty
    {
        get;
        set;
    }
}
`
	assertCode(t, Class("MyClass").Members(Property("Result","MyProperty").Get().Set()), expected)
}

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
	method :=
		Method("MyMethod").Public().Params(
			Param("int", "intParam"),
			Param("string", "stringParam"),
		)
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

func TestPropertyGetSet(t *testing.T) {
	expected := `
MyType MyProperty
{
    get;
    set;
}
`
	assertCode(t, Property("MyType", "MyProperty").Get().Set(), expected)
}

func TestFieldPublicStatic(t *testing.T) {
	assertCode(t, Field("MyType", "myField").Public().Static(), "public static MyType myField;")
}

func TestFieldInitializer(t *testing.T) {
	assertCode(t, Field("int", "myField").Init(C("3")), "int myField = 3;")
}