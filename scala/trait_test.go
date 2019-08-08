package scala

import (
	"testing"
)

func TestTraitBasic(t *testing.T) {
	expected := `
trait MyTrait {
}
`
	assertCode(t, Trait("MyTrait"), expected)
}

func TestTraitMethod(t *testing.T) {
	expected := `
trait MyTrait {
  def MyMethod(): Unit
}
`
	trait := Trait("MyTrait")
	trait.Def("MyMethod")
	assertCode(t, trait, expected)
}

func TestTraitWithAttribute(t *testing.T) {
	expected := `
@MyAttribute
trait MyTrait {
  def MyMethod(): Unit
}
`
	trait := Trait("MyTrait").WithAttribute("MyAttribute")
	trait.Def("MyMethod")
	assertCode(t, trait, expected)
}
