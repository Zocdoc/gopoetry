package typescript

import (
	"testing"
)

func TestTypeDecBasic(t *testing.T) {
	expected := `
type UUID = string;
`
	assertCode(t, DeclareType("UUID", UnionType("string")), expected)
}

func TestUnionTypeDec(t *testing.T) {
	expected := `
type NumOrString = string | number;
`
	typeDec := DeclareType("NumOrString", UnionType("string", "number"))
	assertCode(t, typeDec, expected)
}

func TestUnionLiteralTypeDec(t *testing.T) {
	expected := `
type DiningType = 'in_door' | 'out_door';
`
	typeDec := DeclareType("DiningType", UnionType("'in_door'", "'out_door'"))
	assertCode(t, typeDec, expected)
}
