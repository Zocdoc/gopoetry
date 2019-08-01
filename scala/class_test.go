package scala

import (
	"testing"
)

func TestClassBasic(t *testing.T) {
	expected := `
class MyClass {
}
`
	assertCode(t, Class("MyClass"), expected)
}

func TestClassExtends(t *testing.T) {
	expected := `
class MyClass extends BaseClass {
}
`
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
	class.Def("MyMethod").Returns("Unit").As().Block()
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
	class := Class("MyClass").WithAttribute("MyAttribute")
	class.Def("MyMethod").Returns("Unit").As().Block()
	assertCode(t, class, expected)
}

func TestClassWithCtorAttribute(t *testing.T) {
	expected := `
class MyClass @MyAttribute()() {
  def MyMethod(): Unit = {
  }
}
`
	class := Class("MyClass").WithCtorAttribute("MyAttribute()")
	class.Def("MyMethod").Returns("Unit").As().Block()
	assertCode(t, class, expected)
}
