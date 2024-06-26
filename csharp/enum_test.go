package csharp

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

func TestEnumInherits(t *testing.T) {
	expected := `
enum MyEnum: int
{
}
`
	assertCode(t, Enum("MyEnum").Inherits("int"), expected)
}

func TestEnumModifier(t *testing.T) {
	expected := `
private enum MyEnum
{
}
`
	assertCode(t, Enum("MyEnum").Private(), expected)
}

func TestEnumAttributed(t *testing.T) {
	expected := `
[MyAttribute]
enum MyEnum
{
}
`
	assertCode(t, Enum("MyEnum").WithAttribute("MyAttribute"), expected)
}

func TestEnumWithMembers(t *testing.T) {
	expected := `
/// <summary>
/// Enum Summary
/// </summary>
enum MyEnum
{
    /// <summary>
    /// member 1
    /// </summary>
    Member1,
    /// <summary>
    /// member 2
    /// </summary>
    Member2,
}
`
	enum := Enum("MyEnum")
	enum.Summary("Enum Summary")

	enum.Member("Member1").Summary("member 1")
	enum.Member("Member2").Summary("member 2")
	assertCode(t, enum, expected)
}
