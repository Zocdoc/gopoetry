package typescript

import "testing"

func TestEnumMemberBasic(t *testing.T) {
	expected := `MyMember`
	assertCode(t, EnumMember("MyMember"), expected)
}

func TestEnumMemberInit(t *testing.T) {
	expected := `MyMember = 1`
	assertCode(t, EnumMember("MyMember").Value(Int(1)), expected)
}
