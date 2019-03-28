package csharp

import "testing"

func TestUnitEmpty(t *testing.T) {
	assertCode(t, Unit(), "")
}

func TestUnitImportWithNamespace(t *testing.T) {
	unit := Unit().Import("MyNamespace").ImportStatic("MyOther.Namespace")
	unit.Namespace("Some.Namespace")
	unit.Namespace("Other.Namespace")

	expected := `
import MyNamespace;
import static MyOther.Namespace;

namespace Some.Namespace
{
}

namespace Other.Namespace
{
}
`

	assertCode(t, unit, expected)
}
