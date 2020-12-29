package scala

import "testing"

func TestCodePrimitives(t *testing.T) {
	assertCode(t, Code("some code"), "some code")
	//assertCode(t, Line("some line"), "some line\n")
	assertCode(t, NoCode, "")
	assertCode(t, Str("some string"), `"some string"`)
	assertCode(t, Int(123), "123")
	assertCode(t, True, "true")
	assertCode(t, False, "false")
	assertCode(t, Null, "null")
}

