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
class MyClass : MyParent
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

func TestMultipleClassModifier(t *testing.T) {
	expected := `
private abstract class MyClass
{
}
`
	assertCode(t, Class("MyClass").Private().Abstract(), expected)
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
    Result MyProperty { get; set; }
}
`
	class := Class("MyClass")
	property := class.Property("Result", "MyProperty")
	property.Get()
	property.Set()
	assertCode(t, class, expected)
}

func TestClassPropertyWithInitializer(t *testing.T) {
	expected := `
class MyClass
{
    Result MyProperty { get; set; } = "bar";
}
`
	class := Class("MyClass")
	property := class.Property("Result", "MyProperty")
	property.Get()
	property.Set()
	property.Initialize(Str("bar"))
	assertCode(t, class, expected)
}

func TestClassPropertyWithInit(t *testing.T) {
	expected := `
class MyClass
{
    Result MyProperty { get; init; }
}
`
	class := Class("MyClass")
	property := class.Property("Result", "MyProperty")
	property.Get()
	property.Init()
	assertCode(t, class, expected)
}

func TestClassWithSummary(t *testing.T) {
	expected := `
/// <summary>
/// my class summary
/// </summary>
class MyClass
{
}
`
	assertCode(t, Class("MyClass").Summary("my class summary"), expected)
}
