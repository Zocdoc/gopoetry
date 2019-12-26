package scala

import (
	"testing"
)

func TestTraitBasic(t *testing.T) {
	expected := `trait MyTrait`
	assertCode(t, Trait("MyTrait"), expected)
}

func TestTraitMethod(t *testing.T) {
	expected := `
trait MyTrait {
  def MyMethod(): Unit
}
`
	trait := Trait("MyTrait")
	trait.Define(true).Def("MyMethod").Returns("Unit")
	assertCode(t, trait, expected)
}

func TestTraitWithAttribute(t *testing.T) {
	expected := `
@MyAttribute
trait MyTrait {
  def MyMethod(): Unit
}
`
	trait := Trait("MyTrait").Attribute("MyAttribute")
	trait.Define(true).Def("MyMethod").Returns("Unit")
	assertCode(t, trait, expected)
}

func TestTraitSealed(t *testing.T) {
	expected := `sealed trait MyTrait`
	trait := Trait("MyTrait").Sealed()
	assertCode(t, trait, expected)
}
