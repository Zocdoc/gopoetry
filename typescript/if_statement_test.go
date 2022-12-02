package typescript

import (
	"testing"
)

func TestBasicIf(t *testing.T) {
	expected := `
if (true) {
}
`
	assertCode(t, If(True), expected)
}

func TestBasicNrach(t *testing.T) {
	expected := `
if (true) {
    return -1;
} else {
    return 1;
}
`

	ifStmt := If(True).
		IfBlockAppend(C("return -1;")).
		ElseBlockAppend(C("return 1;"))

	assertCode(t, ifStmt, expected)
}
