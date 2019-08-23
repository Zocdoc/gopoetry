package java

import (
	"gopoetry/util"
	"testing"
)

func TestEnumBasic(t *testing.T) {
	expected := `
enum MyEnum {
}
`
	util.AssertCode(t, Enum("MyEnum"), expected)
}

func TestEnumPublic(t *testing.T) {
	expected := `
public enum MyEnum {
}
`
	util.AssertCode(t, Enum("MyEnum").Public(), expected)
}

func TestEnumPrivate(t *testing.T) {
	expected := `
private enum MyEnum {
}
`
	util.AssertCode(t, Enum("MyEnum").Private(), expected)
}

func TestEnumWithMembers(t *testing.T) {
	expected := `
enum MyEnum {
  foo("bar");

  void MyMethod() {
  }
}
`
	enum := Enum("MyEnum")
	enum.AddEnumMembers("foo", "bar")
	method := Method("MyMethod").Returns("void")
	method.Define().Block()
	enum.AddMembers(method)
	util.AssertCode(t, enum, expected)
}

func TestEnumWithConstructor(t *testing.T) {
	expected := `
enum MyEnum {

  MyEnum(String input) {
    String argh = input;
  }
}
`

	enum := Enum("MyEnum")
	constructor := enum.Constructor()
	constructor.Param("input", "String")
	constructor.Define().Block().Line("String argh = input;")
	util.AssertCode(t, enum, expected)
}
