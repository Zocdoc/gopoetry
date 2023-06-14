package swift

import "testing"

func TestEmptyEnum(t *testing.T) {
	myStruct := NewEnum("MyEnum")

	assertCode(t, myStruct, "enum MyEnum {\n}")
}

func TestEnumOneCase(t *testing.T) {
	myStruct := NewEnum("MyEnum")
	myStruct.AddCases(NewCase("a"))

	expected := `
enum MyEnum {
    case a
}
`
	assertCode(t, myStruct, expected)
}

func TestEnumMultipleCases(t *testing.T) {
	myStruct := NewEnum("MyEnum")
	myStruct.AddCases(
		NewCase("one"),
		NewCase("two"),
		NewCase("three"),
	)

	expected := `
enum MyEnum {
    case one
    case two
    case three
}
`
	assertCode(t, myStruct, expected)
}
