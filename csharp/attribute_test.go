package csharp

import "testing"

func TestAttributeSimple(t *testing.T) {
	assertCode(t, Attribute("MyAttribute(a, b, c = c)"), `MyAttribute(a, b, c = c)`)
}
