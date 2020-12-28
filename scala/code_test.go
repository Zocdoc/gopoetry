package scala

import "testing"

func TestCodePrimitives(t *testing.T) {
	assertCode(t, Code("some code"), "some code")
	assertCode(t, CodeIf(false,"some code not printed"), "")
	assertCode(t, CodeIf(true,"some code printed"), "some code printed")
	//assertCode(t, Line("some line"), "some line\n")
	//assertCode(t, LineIf(false,"some line not printed"), "")
	//assertCode(t, LineIf(true, "some line printed"), "some line printed\n")
	assertCode(t, NoCode(), "")
	assertCode(t, Str("some string"), `"some string"`)
	assertCode(t, Int(123), "123")
	assertCode(t, True, "true")
	assertCode(t, False, "false")
	assertCode(t, Null, "null")
}

