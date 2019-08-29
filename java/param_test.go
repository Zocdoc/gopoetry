package java

import (
	"gopoetry/util"
	"testing"
)

func TestValSimple(t *testing.T) {
	util.AssertCode(t, Param("prop", "String"), `String prop`)
}
