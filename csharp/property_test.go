package csharp

import "testing"

func TestPropertyGetSet(t *testing.T) {
	expected := `
MyType MyProperty
{
    get;
    set;
}
`
	property := Property("MyType", "MyProperty")
	property.Get()
	property.Set()
	assertCode(t, property, expected)
}

func TestPropertyGetWithInitalizer(t *testing.T) {
	expected := `
string MyProperty
{
    get;
}
 = "foo";
`
	property := Property("string", "MyProperty")
	property.Get()
	property.Init(Str("foo"))
	assertCode(t, property, expected)
}

func TestPropertyPublic(t *testing.T) {
	expected := `
public MyType MyProperty
{
    get;
    set;
}
`
	property := Property("MyType", "MyProperty").Public()
	property.Get()
	property.Set()
	assertCode(t, property, expected)
}

func TestPropertyPrivateSetter(t *testing.T) {
	expected := `
MyType MyProperty
{
    get;
    private set;
}
`
	property := Property("MyType", "MyProperty")
	property.Get()
	property.Set().Private()
	assertCode(t, property, expected)
}

func TestPropertyAttributed(t *testing.T) {
	expected := `
[MyAttribute]
public MyType MyProperty
{
    get;
    set;
}
`
	property := Property("MyType", "MyProperty").Public().WithAttribute("MyAttribute")
	property.Get()
	property.Set()
	assertCode(t, property, expected)
}
