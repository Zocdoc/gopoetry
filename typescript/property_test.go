package typescript

import "testing"

func TestPropertyPublic(t *testing.T) {
	expected := `
public MyProperty: MyType;
`
	property := Property("MyType", "MyProperty").Public()
	assertCode(t, property, expected)
}

func TestPropertyOptional(t *testing.T) {
	expected := `
private MyProperty?: MyType;
`
	property := Property("MyType", "MyProperty")
	property.Private().Optional()
	assertCode(t, property, expected)
}

func TestPropertyOptionalWithInitializer(t *testing.T) {
	expected := `
private MyProperty?: MyType = 'example';
`
	property := Property("MyType", "MyProperty")
	property.Private().Optional().Initializer("'example'")
	assertCode(t, property, expected)
}
