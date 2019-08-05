package scala

import (
	"testing"
)

func TestClassBasic(t *testing.T) {
	expected := `class MyClass`
	assertCode(t, Class("MyClass"), expected)
}

func TestClassExtends(t *testing.T) {
	expected := `class MyClass extends BaseClass`
	assertCode(t, Class("MyClass").Extends("BaseClass"), expected)
}


func TestClassMethod(t *testing.T) {
	expected := `
class MyClass {
  def MyMethod(): Unit = {
  }
}
`
	class := Class("MyClass")
	class.Def("MyMethod").Returns("Unit").Define().Block()
	assertCode(t, class, expected)
}

func TestClassWithAttribute(t *testing.T) {
	expected := `
@MyAttribute
class MyClass {
  def MyMethod(): Unit = {
  }
}
`
	class := Class("MyClass").Attribute("MyAttribute")
	class.Def("MyMethod").Returns("Unit").Define().Block()
	assertCode(t, class, expected)
}

func TestClassWithCtorParam(t *testing.T) {
	expected := `class MyClass(param1: String)`
	class := Class("MyClass")
	class.Contructor().Param("param1", "String")
	assertCode(t, class, expected)
}

func TestClassWithCtorImplicitParam(t *testing.T) {
	expected := `class MyClass(implicit param1: String)`
	class := Class("MyClass")
	class.Contructor().NoParams().ImplicitParam("param1", "String")
	assertCode(t, class, expected)
}

func TestClassWithCtorAttribute(t *testing.T) {
	expected := `
class MyClass @MyAttribute()() {
  def MyMethod(): Unit = {
  }
}
`
	class := Class("MyClass")
	class.Contructor().Attribute("MyAttribute()")
	class.Def("MyMethod").Returns("Unit").Define().Block()
	assertCode(t, class, expected)
}
