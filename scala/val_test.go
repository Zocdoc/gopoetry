package scala

import "testing"

func TestValSimple(t *testing.T) {
	assertCode(t, Val("prop", "String"), `prop: String`)
}

func TestValWithAttribute(t *testing.T) {
	assertCode(t, Val("prop", "String").WithAttribute("JsonProperty"), `@JsonProperty prop: String`)
}
