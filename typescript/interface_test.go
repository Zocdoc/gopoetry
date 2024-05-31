package typescript

import (
	"testing"
)

func TestInterfaceBasic(t *testing.T) {
	expected := `
interface MyInterface
{
}
`
	assertCode(t, Interface("MyInterface"), expected)
}

func TestInterfaceMethod(t *testing.T) {
	expected := `
/** makes a thing */
interface MyInterface
{
    /** does a thing */
    MyMethod(): void;
}
`
	iface := Interface("MyInterface")
	iface.AddComments("makes a thing")
	iface.Method("MyMethod").AddComments("does a thing")

	assertCode(t, iface, expected)
}

func TestExportedInterfaceWithProperty(t *testing.T) {
	expected := `
export interface MyInterface extends IfaceOne, IfaceTwo
{
    /** MyProperty */
    MyProperty?: Result;
}
`
	iface := Interface("MyInterface").Export().Extends("IfaceOne", "IfaceTwo")
	_ = iface.Property("Result", "MyProperty").Optional().AddComments("MyProperty")
	assertCode(t, iface, expected)
}

func TestTypeLiteralInterface(t *testing.T) {
	expected := `
export interface MyInterface
{
    paths: {
        foo: {
            get: string;
        };
    };
}
`
	iface := Interface("MyInterface").Export()
	iface.AddMembers(&PropertySig{
		Name: "paths",
		TypeAnnotation: NewObjectType().AddProp(&PropertySig{
			Name: "foo",
			TypeAnnotation: NewObjectType().AddProp(&PropertySig{
				Name:           "get",
				TypeAnnotation: Code("string"),
			}),
		}),
	})

	assertCode(t, iface, expected)
}
