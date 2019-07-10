package typescript

import "testing"

func TestNamedImport(t *testing.T) {
	assertCode(t, NamedImport("modname", "foo"), "import { foo } from 'modname';")
}

func TestDefaultImport(t *testing.T) {
	assertCode(t, DefaultImport("modname", "foo"), "import * as foo from 'modname';")
}
