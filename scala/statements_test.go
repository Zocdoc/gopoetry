package scala

import "testing"

func TestStatementsSimple(t *testing.T) {
	expected := `
line1()
line2()
`
	assertCode(t, Statements("line1()", "line2()"), expected)
}

func TestStatementsWithBlock(t *testing.T) {
	expected := `
line1()
line2 {
  nextedLine1()
  nextedLine2()
}
`
	statements := Statements()
	statements.AppendLine("line1()")
	statements.Append("line2 ")
	statements.Block(
		"nextedLine1()",
		"nextedLine2()",
	)
	assertCode(t, statements, expected)
}
