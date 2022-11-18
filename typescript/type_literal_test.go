package typescript

import "testing"

func TestObjectType(t *testing.T) {
	expected := `
type MyType = {
    foo: string;
};
	`

	ot := &ObjectType{}
	ot.AddMember("foo", Code("string"))

	assertCode(t, DeclareType("MyType", ot), expected)
}

func TestNestedObjectType(t *testing.T) {
	expected := `
type MyType = {
    foo: {
        foo: string;
    };
    bar: 'test';
};
`

	objectTypeDecl := DeclareType(
		"MyType",
		NewObjectType(
			"foo",
			NewObjectType("foo", Code("string")),
		).
			AddMember("bar", Str("test")),
	)

	assertCode(t, objectTypeDecl, expected)
}
