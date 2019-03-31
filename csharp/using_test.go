package csharp

import "testing"

func TestUsing(t *testing.T) {
	assertCode(t, Using("Example.Namespace"), "using Example.Namespace;")
}

func TestUsingStatic(t *testing.T) {
	assertCode(t, Using("Example.Namespace").Static(), "using static Example.Namespace;")
}
