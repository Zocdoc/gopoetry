package scala

import "testing"

func TestValSimple(t *testing.T) {
	assertCode(t, Val("prop", "String"), `val prop: String`)
}

func TestVarSimple(t *testing.T) {
	assertCode(t, Var("prop", "String"), `var prop: String`)
}

func TestValWithAttribute(t *testing.T) {
	assertCode(t, Val("prop", "String").Attribute("JsonProperty"), `@JsonProperty val prop: String`)
}

func TestValWithOverride(t *testing.T) {
	assertCode(t, Val("prop", "String").Override(), `override val prop: String`)
}
