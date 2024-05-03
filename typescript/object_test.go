package typescript

import "testing"

func TestObjectSpreads(t *testing.T) {
	expected := `
let myConfig = {
    ...someOtherObject,
    a: 45,
    b: 'test',
};
`

	obj := ObjectValue{}
	obj.AddSpread("someOtherObject")
	obj.AddProp("a", Code("45"))
	obj.AddProp("b", Str("test"))

	myConfig := Let("myConfig", nil, &obj)
	assertCode(t, myConfig, expected)
}
