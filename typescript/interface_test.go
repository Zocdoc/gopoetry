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
interface MyInterface
{
    MyMethod(): void;
}
`
	iface := Interface("MyInterface")
	iface.Method("MyMethod")
	assertCode(t, iface, expected)
}

func TestExportedInterfaceWithProperty(t *testing.T) {
	expected := `
export interface MyInterface extends IfaceOne, IfaceTwo
{
    MyProperty?: Result;
}
`
	iface := Interface("MyInterface").Export().Extends("IfaceOne", "IfaceTwo")
	_ = iface.Property("Result", "MyProperty").Optional()
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
