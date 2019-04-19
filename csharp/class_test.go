package csharp

import (
	"testing"
)

func TestClassBasic(t *testing.T) {
	expected := `
class MyClass
{
}
`
	assertCode(t, Class("MyClass"), expected)
}

func TestClassInherits(t *testing.T) {
	expected := `
class MyClass: MyParent
{
}
`
	assertCode(t, Class("MyClass").Inherits("MyParent"), expected)
}

func TestClassModifier(t *testing.T) {
	expected := `
private class MyClass
{
}
`
	assertCode(t, Class("MyClass").Private(), expected)
}

func TestClassAttributed(t *testing.T) {
	expected := `
[MyAttribute]
class MyClass
{
}
`
	assertCode(t, Class("MyClass").WithAttribute("MyAttribute"), expected)
}

func TestClassMultipleAttributes(t *testing.T) {
	expected := `
[MyAttribute1, MyAttribute2]
class MyClass
{
}
`
	assertCode(t, Class("MyClass").WithAttribute("MyAttribute1").WithAttribute("MyAttribute2"), expected)
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
	class := Class("MyClass")
	class.Method("MyMethod").Body()
	assertCode(t, class, expected)
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
	class := Class("MyClass")
	class.Property("Result", "MyProperty").Get().Set()
	assertCode(t, class, expected)
}
