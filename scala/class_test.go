package scala

import (
	"testing"
)

func TestClassBasic(t *testing.T) {
	expected := `class MyClass`
	assertCode(t, Class("MyClass"), expected)
}

func TestClassExtends(t *testing.T) {
	expected := `class MyClass extends BaseClass with Trait`
	assertCode(t, Class("MyClass").Extends("BaseClass").With("Trait"), expected)
}

func TestClassMethod(t *testing.T) {
	expected := `
class MyClass {
  def MyMethod(): Unit = {
  }
}`
	class :=
		Class("MyClass").Members(
			Def("MyMethod").Returns("Unit").Body(),
		)
	assertCode(t, class, expected)
}

func TestClassWithAttribute(t *testing.T) {
	expected := `
@MyAttribute
class MyClass {
  def MyMethod(): Unit = {
  }
}`
	class :=
		Class("MyClass").Attribute("MyAttribute").Members(
			Def("MyMethod").Returns("Unit").Body(),
		)
	assertCode(t, class, expected)
}

func TestClassWithCtorParam(t *testing.T) {
	expected := `class MyClass(param1: String)`
	class := Class("MyClass")
	class.Constructor(Constructor().AddParams(Param("param1", "String")))
	assertCode(t, class, expected)
}

func TestClassWithPrivateCtor(t *testing.T) {
	expected := `class MyClass private (param1: String)`
	class := Class("MyClass")
	class.Constructor(Constructor().Private().AddParams(Param("param1", "String")))
	assertCode(t, class, expected)
}

func TestClassWithCtorImplicitParam(t *testing.T) {
	expected := `class MyClass(implicit param1: String)`
	class := Class("MyClass")
	class.Constructor(Constructor().AddImplicitParams(Param("param1", "String")))
	assertCode(t, class, expected)
}

func TestClassWithCtorAttribute(t *testing.T) {
	expected := `
class MyClass @MyAttribute()() {
  def MyMethod(): Unit = {
  }
}`
	class :=
		Class("MyClass").
			Constructor(Constructor().Attribute("MyAttribute()")).
			Members(
				Def("MyMethod").Returns("Unit").Body(),
			)
	assertCode(t, class, expected)
}

func TestObjectBasic(t *testing.T) {
	expected := `object MyObject`
	assertCode(t, Object("MyObject"), expected)
}

func TestEnumCaseObject(t *testing.T) {
	expected := `case object Yes extends Answer { override def toString = "yes" }`
	object :=
		CaseObject("Yes").Extends("Answer").MembersInline(
			Def("toString").Override().NoParams().BodyInline(Code(`"yes"`)),
		)
	assertCode(t, object, expected)
}

func TestCaseClass(t *testing.T) {
	expected := `
case class SomeClass(
  field1: Int,
  field2: String
)`
	class :=
		CaseClass(`SomeClass`).
			CtorParamPerLine().
			CtorParams(
				Param(`field1`, `Int`),
				Param(`field2`, `String`),
			)

	assertCode(t, class, expected)
}