package java

import (
	"testing"
)

func TestValSimple(t *testing.T) {
	AssertCode(t, Param("prop", "String"), `String prop`)
}
