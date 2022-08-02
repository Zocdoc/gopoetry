package csharp

import "testing"

func TestEnumMemberBasic(t *testing.T) {
	expected := `MyMember`
	assertCode(t, EnumMember("MyMember"), expected)
}

func TestEnumMemberInit(t *testing.T) {
	expected := `MyMember = 1`
	assertCode(t, EnumMember("MyMember").Value("1"), expected)
}

func TestEnumMemberAttributed(t *testing.T) {
	expected := `
[MyAttribute]
MyMember
`
	assertCode(t, EnumMember("MyMember").WithAttribute("MyAttribute"), expected)
}
