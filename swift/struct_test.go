package swift

import (
	"fmt"
	"testing"
)

func TestEmptyStruct(t *testing.T) {
	myStruct := NewStruct("MyStruct")

	assertCode(t, myStruct, "struct MyStruct {}")
}

func TestStructSimpleMember(t *testing.T) {
	myStruct := NewStruct("MyStruct")
	myStruct.AddMembers(Var("name", "String"))

	expected := `
struct MyStruct {
    var name: String
}
`
	assertCode(t, myStruct, expected)
}

func TestStructSimpleMemberInit(t *testing.T) {
	myStruct := NewStruct("MyStruct")
	myStruct.AddMembers(
		Var("name", "String").InitWith(Str("some_name")),
	)

	expected := `
struct MyStruct {
    var name: String = "some_name"
}
`
	assertCode(t, myStruct, expected)
}

func TestPuplicStructSimpleMemberInit(t *testing.T) {
	myStruct := NewStruct("MyStruct")
	myStruct.AddMembers(
		Var("name", "String").
			InitWith(Str("some_name")).
			Public(),
	)

	expected := `
struct MyStruct {
    public var name: String = "some_name"
}
`
	assertCode(t, myStruct, expected)
}

func TestStructInStruct(t *testing.T) {
	myStruct := NewStruct("MyStruct")
	myStruct.AddMembers(
		NewStruct("InnerStruct").
			AddMembers(Var("age", "Int")),
	)

	expected := `
struct MyStruct {
    struct InnerStruct {
        var age: Int
    }
}
`
	assertCode(t, myStruct, expected)
}

func TestPublicStructInStruct(t *testing.T) {
	myStruct := NewStruct("MyStruct")
	myStruct.AddMembers(
		NewStruct("InnerStruct").
			AddMembers(Var("age", "Int")).
			Public(),
	)

	expected := `
struct MyStruct {
    public struct InnerStruct {
        var age: Int
    }
}
`
	assertCode(t, myStruct, expected)
}

func TestStructAccessModifiers(t *testing.T) {
	myStruct := NewStruct("MyStruct")

	modifiers := []func() *StructDecl{
		myStruct.Public,
		myStruct.FilePrivate,
		myStruct.Private,
		myStruct.Internal,
	}

	expected := []string{
		"public",
		"fileprivate",
		"private",
		"internal",
	}

	for i, modify := range modifiers {
		modify()
		expected := fmt.Sprintf("%s struct MyStruct {}", expected[i])
		assertCode(t, myStruct, expected)
	}
}
