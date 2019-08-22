package java

import (
	"github.com/zocdoc/gopoetry/util"
	"testing"
)

func TestValSimple(t *testing.T) {
	util.AssertCode(t, Param("prop", "String"), `String prop`)
}
