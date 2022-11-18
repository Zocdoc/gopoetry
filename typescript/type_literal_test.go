package typescript

import "testing"

func TestObjectType(t *testing.T) {
	expected := `
type MyType = {
    foo: string;
};
	`

	ot := &ObjectType{}
	ot.AddProp(&PropertySig{
		name:           "foo",
		typeAnnotation: Code("string"),
	})

	assertCode(t, DeclareType("MyType", ot), expected)
}

func TestNestedObjectType(t *testing.T) {
	expected := `
type MyType = {
    foo: {
        foo: string;
    };
    bar?: 'test';
};
`

	objectTypeDecl := DeclareType(
		"MyType",
		NewObjectType().
			AddProp(&PropertySig{
				name: "foo",
				typeAnnotation: NewObjectType().AddProp(&PropertySig{
					name:           "foo",
					typeAnnotation: Code("string"),
				}),
			}).
			AddProp(&PropertySig{
				name:           "bar",
				optional:       true,
				typeAnnotation: Str("test"),
			}),
	)

	assertCode(t, objectTypeDecl, expected)
}
