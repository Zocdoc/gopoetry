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
    Result MyProperty
    {
        get;
        set;
    }
}
`
	iface := Interface("MyInterface")
	iface.Property("Result", "MyProperty").Get().Set()
	assertCode(t, iface, expected)
}
