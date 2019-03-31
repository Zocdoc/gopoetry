package csharp

import "testing"

func TestUnitEmpty(t *testing.T) {
	assertCode(t, Unit(), "")
}

func TestUnitUsingWithNamespace(t *testing.T) {
	unit := Unit().Using("MyNamespace").UsingStatic("MyOther.Namespace")
	unit.Namespace("Some.Namespace")
	unit.Namespace("Other.Namespace")

	expected := `
using MyNamespace;
using static MyOther.Namespace;

namespace Some.Namespace
{
}

namespace Other.Namespace
{
}
`

	assertCode(t, unit, expected)
}
