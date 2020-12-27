package scala

import (
	"testing"
)

func TestTraitBasic(t *testing.T) {
	expected := `trait MyTrait`
	assertCode(t, Trait("MyTrait"), expected)
}

func TestTraitExtends(t *testing.T) {
	expected := `trait MyTrait extends BaseTrait`
	assertCode(t, Trait("MyTrait").Extends("BaseTrait"), expected)
}

func TestTraitExtendsMultiple(t *testing.T) {
	expected := `trait MyTrait extends BaseTrait with BaseTrait2`
	assertCode(t, Trait("MyTrait").Extends("BaseTrait", "BaseTrait2"), expected)
}

func TestTraitMethod(t *testing.T) {
	expected := `
trait MyTrait {
  def MyMethod(): Unit
}
`
	trait := Trait("MyTrait")
	trait.Define().Def("MyMethod").Returns("Unit")
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
	trait.Define().Def("MyMethod").Returns("Unit")
	assertCode(t, trait, expected)
}

func TestTraitSealed(t *testing.T) {
	expected := `sealed trait MyTrait`
	trait := Trait("MyTrait").Sealed()
	assertCode(t, trait, expected)
}
