package typescript

import (
	"testing"
)

func TestVarDecl(t *testing.T) {
	expected := `
var myNum = 42;
`

	myNum := Var("myNum", nil, Code("42"))
	assertCode(t, myNum, expected)
}

func TestTypeOnly(t *testing.T) {
	expected := `
var myNum: number;
`

	myNum := Var("myNum", Code("number"), nil)
	assertCode(t, myNum, expected)
}

func TestConstWithValAndType(t *testing.T) {
	expected := `
const myNum: number = 45;
`

	myNum := Const("myNum", Code("number"), Code("45"))
	assertCode(t, myNum, expected)
}

func TestLetObject(t *testing.T) {
	expected := `
let myConfig = {
    a: 45,
    b: 'test',
};
`

	obj := ObjectValue{}
	obj.AddProp("a", Code("45"))
	obj.AddProp("b", Str("test"))

	myConfig := Let("myConfig", nil, &obj)
	assertCode(t, myConfig, expected)
}

func TestExportedNested(t *testing.T) {
	expected := `
export const myConfig = {
    a: 45,
    b: 'test',
    nest: {
        bar: 1,
        baz: true,
    },
};
`

	nested := &ObjectValue{}
	nested.AddProp("bar", Code("1"))
	nested.AddProp("baz", True)

	obj := ObjectValue{}
	obj.AddProp("a", Code("45"))
	obj.AddProp("b", Str("test"))
	obj.AddProp("nest", nested)

	myConfig := Const("myConfig", nil, &obj).Export()
	assertCode(t, myConfig, expected)
}
