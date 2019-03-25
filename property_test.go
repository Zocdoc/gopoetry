package poetrycs

import "testing"

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

func TestPropertyPublic(t *testing.T) {
	expected := `
public MyType MyProperty
{
    get;
    set;
}
`
	assertCode(t, Property("MyType", "MyProperty").Get().Set().Public(), expected)
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
	assertCode(t, Property("MyType", "MyProperty").Get().Set().Public().WithAttribute("MyAttribute"), expected)
}
