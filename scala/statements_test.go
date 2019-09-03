package scala

import "testing"

func TestStatementsSimple(t *testing.T) {
	expected := `
line1()
line2()
`
	statements :=
		Statements().
			AddLn("line1()").
			AddLn("line2()")

	assertCode(t, statements, expected)
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
	statements.AddLn("line1()")
	statements.Add("line2 ")
	statements.Block(true).
		AddLn("nextedLine1()").
		AddLn("nextedLine2()")
	assertCode(t, statements, expected)
}

func TestStatementsWithBlockNoScope(t *testing.T) {
	expected := `
request =>
  nextedLine1()
  nextedLine2()
`
	statements := Statements()
	statements.AddLn("request =>")
	statements.Block(false).
		AddLn("nextedLine1()").
		AddLn("nextedLine2()")
	assertCode(t, statements, expected)
}
