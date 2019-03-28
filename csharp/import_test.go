package csharp

import "testing"

func TestImport(t *testing.T) {
	assertCode(t, Import("Example.Namespace"), "import Example.Namespace;")
}

func TestImportStatic(t *testing.T) {
	assertCode(t, Import("Example.Namespace").Static(), "import static Example.Namespace;")
}
