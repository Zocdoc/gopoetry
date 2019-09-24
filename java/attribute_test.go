package java

import (
	"testing"
)

func TestAttributeSimple(t *testing.T) {
	AssertCode(t, Attribute("MyAttribute(a, b, c = c)"), `@MyAttribute(a, b, c = c)`)
}
