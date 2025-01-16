package csharp

import (
	"testing"
)

func TestInterfaceBasic(t *testing.T) {
	expected := `
interface MyInterface
{
}
`
	assertCode(t, Interface("MyInterface"), expected)
}

func TestInterfaceMethod(t *testing.T) {
	expected := `
interface MyInterface
{
    void MyMethod();
}
`
	iface := Interface("MyInterface")
	iface.Method("MyMethod")
	assertCode(t, iface, expected)
}

func TestInterfaceProperty(t *testing.T) {
	expected := `
interface MyInterface
{
    Result MyProperty { get; set; }
}
`
	iface := Interface("MyInterface")
	property := iface.Property("Result", "MyProperty")
	property.Get()
	property.Set()
	assertCode(t, iface, expected)
}

func TestInterfaceSummary(t *testing.T) {
	expected := `
/// <summary>
/// This is a summary
/// </summary>
interface MyInterface
{
}
`
	iface := Interface("MyInterface")
	iface.Summary("This is a summary")
	assertCode(t, iface, expected)
}
