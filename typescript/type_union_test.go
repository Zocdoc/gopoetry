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

func TestUnionBuilderFunction(t *testing.T) {
	expected := `
type DiningType = 'in_door' | 'out_door' | 'underwater' | 'up_in_the_air';
`
	elements := UnionType("'underwater'", "'up_in_the_air'")
	doors := UnionType("'in_door'", "'out_door'")

	typeDec := DeclareType("DiningType", doors.Union(elements))
	assertCode(t, typeDec, expected)
}

func TestExportUnion(t *testing.T) {
	expected := `
export type DiningType = 'always';
`
	typeDec := DeclareType("DiningType", Str("always")).Export()
	assertCode(t, typeDec, expected)
}
