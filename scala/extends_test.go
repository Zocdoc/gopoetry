package scala

import "testing"

func TestExtendsBaseClassNoParamsNoWiths(t *testing.T) {
	assertCode(t, Extends("BaseClass"), "extends BaseClass")
}

func TestExtendsBaseClassAndParamsNoWiths(t *testing.T) {
	assertCode(t, Extends("BaseClass", "foo", "bar"), "extends BaseClass(foo, bar)")
}

func TestExtendsBaseClassAndParamsAndWiths(t *testing.T) {
	assertCode(t, Extends("BaseClass", "foo", "bar").With("BaseTrait"),
		"extends BaseClass(foo, bar) with BaseTrait")
}

func TestExtendsBaseClassAndParamsAndMultipleWiths(t *testing.T) {
	assertCode(t, Extends("BaseClass", "foo", "bar").With("BaseTrait1").With("BaseTrait2"),
		"extends BaseClass(foo, bar) with BaseTrait1 with BaseTrait2")
}

func TestExtendsBaseClassAndNoParamsAndMultipleWiths(t *testing.T) {
	assertCode(t, Extends("BaseClass").With("BaseTrait1").With("BaseTrait2"),
		"extends BaseClass with BaseTrait1 with BaseTrait2")
}
