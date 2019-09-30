package scala

import "testing"

func TestImportSimple(t *testing.T) {
	assertCode(t, Import("com.some.package"), `import com.some.package`)
}
