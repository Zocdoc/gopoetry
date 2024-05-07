package typescript

import "testing"

func TestObjectSpreadsPrePropertyAssignment(t *testing.T) {
	expected := `
let myConfig = {
    ...someOtherObject,
    a: 45,
    b: 'test',
};
`

	obj := ObjectValue{}
	obj.AddSpread("someOtherObject", false)
	obj.AddProp("a", Code("45"))
	obj.AddProp("b", Str("test"))

	myConfig := Let("myConfig", nil, &obj)
	assertCode(t, myConfig, expected)
}

func TestObjectSpreadsPostPropertyAssignment(t *testing.T) {
	expected := `
let myConfig = {
    a: 45,
    b: 'test',
    ...someOtherObject,
};
`

	obj := ObjectValue{}
	obj.AddSpread("someOtherObject", true)
	obj.AddProp("a", Code("45"))
	obj.AddProp("b", Str("test"))

	myConfig := Let("myConfig", nil, &obj)
	assertCode(t, myConfig, expected)
}
