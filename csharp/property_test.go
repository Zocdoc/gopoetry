package csharp

import "testing"

func TestPropertyGetSet(t *testing.T) {
	expected := `
MyType MyProperty { get; set; }
`
	property := Property("MyType", "MyProperty")
	property.Get()
	property.Set()
	assertCode(t, property, expected)
}

func TestPropertyGetWithInitializer(t *testing.T) {
	expected := `
string MyProperty { get; } = "foo";
`
	property := Property("string", "MyProperty")
	property.Get()
	property.Initialize(Str("foo"))
	assertCode(t, property, expected)
}

func TestPropertyPublic(t *testing.T) {
	expected := `
public MyType MyProperty { get; set; }
`
	property := Property("MyType", "MyProperty").Public()
	property.Get()
	property.Set()
	assertCode(t, property, expected)
}

func TestPropertyPrivateSetter(t *testing.T) {
	expected := `
MyType MyProperty { get; private set; }
`
	property := Property("MyType", "MyProperty")
	property.Get()
	property.Set().Private()
	assertCode(t, property, expected)
}

func TestPropertyAttributed(t *testing.T) {
	expected := `
[MyAttribute]
public MyType MyProperty { get; set; }
`
	property := Property("MyType", "MyProperty").Public().WithAttribute("MyAttribute")
	property.Get()
	property.Set()
	assertCode(t, property, expected)
}

func TestPropertyWithSummary(t *testing.T) {
	expected := `
/// <summary>
/// my property summary
/// </summary>
public int intParam { get; set; }
`
	property := Property("int", "intParam").Public().Summary("my property summary")
	property.Get()
	property.Set()
	assertCode(t, property, expected)
}

func TestPropertyWithExpressionBodiedMember(t *testing.T) {
	expected := `
public new LogLevel? MyLogLevel => SomeAwesomeLogLevel.Ignore;
`
	property := Property("new LogLevel?", "MyLogLevel").Public()
	property.ExpressionBodiedMember("SomeAwesomeLogLevel.Ignore")
	assertCode(t, property, expected)
}
