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
		Name:           "foo",
		TypeAnnotation: Code("string"),
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
				Name: "foo",
				TypeAnnotation: NewObjectType().AddProp(&PropertySig{
					Name:           "foo",
					TypeAnnotation: Code("string"),
				}),
			}).
			AddProp(&PropertySig{
				Name:           "bar",
				Optional:       true,
				TypeAnnotation: Str("test"),
			}),
	)

	assertCode(t, objectTypeDecl, expected)
}
