package typescript

import "testing"

func TestObjectType(t *testing.T) {
	expected := `
/** my type */
type MyType = {
    /** my prop */
    foo: string;
};
	`

	ot := &ObjectType{}
	prop := &PropertySig{
		Name:           "foo",
		TypeAnnotation: Code("string"),
	}
	prop.AddComments("my prop")
	ot.AddProp(prop)

	dt := DeclareType("MyType", ot)
	dt.AddComments("my type")

	assertCode(t, dt, expected)
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
