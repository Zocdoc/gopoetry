package scala

import "testing"

func TestLazyBasic(t *testing.T) {
	assertCode(t, Lazy(func() Writable { return Code("some code") }), "some code")
}