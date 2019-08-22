package java

import (
	"github.com/zocdoc/gopoetry/util"
	"testing"
)

func TestMethodSimple(t *testing.T) {
	util.AssertCode(t, Method("someMethod"), `someMethod()`)
}

func TestMethodReturn(t *testing.T) {
	util.AssertCode(t, Method("someMethod").Returns("String"), `String someMethod()`)
}

func TestMethodWithBody(t *testing.T) {
	expected := `
someMethod() {
}
`
	method := Method("someMethod")
	method.Define().Block()
	util.AssertCode(t, method, expected)
}

func TestMethodPrivate(t *testing.T) {
	util.AssertCode(t, Method("someMethod").Private(), `private someMethod()`)
}

func TestMethodAttribute(t *testing.T) {
	util.AssertCode(t, Method("someMethod").Attribute("JsonCreator"), `
@JsonCreator
someMethod()`)
}
