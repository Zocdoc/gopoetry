package typescript

import (
	"testing"
)

func TestEnumBasic(t *testing.T) {
	expected := `
enum MyEnum
{
}
`
	assertCode(t, Enum("MyEnum"), expected)
}

func TestEnumModifier(t *testing.T) {
	expected := `
export enum MyEnum
{
}
`
	assertCode(t, Enum("MyEnum").Export(), expected)
}

func TestEnumWithMembers(t *testing.T) {
	expected := `
enum MyEnum
{
    Member1,
    Member2,
}
`
	enum := Enum("MyEnum")
	enum.Member("Member1")
	enum.Member("Member2")
	assertCode(t, enum, expected)
}

func TestEnumWithMemberValue(t *testing.T) {
	expected := `
enum MyEnum
{
    Member1 = 'foo',
    Member2,
}
`
	enum := Enum("MyEnum")
	enum.Member("Member1").Value(Str("foo"))
	enum.Member("Member2")
	assertCode(t, enum, expected)
}
