package swift

import "testing"

func TestEmptyStruct(t *testing.T) {
	myStruct := NewStruct("MyStruct")

	assertCode(t, myStruct, "struct MyStruct {\n}")
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
