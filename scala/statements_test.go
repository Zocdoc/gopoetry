package scala

import "testing"

func TestStatementsSimple(t *testing.T) {
	expected := `
line1()
line2()
`
	statements :=
		Statements(false, false).
			AddLn("line1()").
			AddLn("line2()")

	assertCode(t, statements, expected)
}

func TestStatementsWithBlockScope(t *testing.T) {
	expected := `
line1()
line2 {
  nextedLine1()
  nextedLine2()
}
`
	statements := Statements(false, false)
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
	statements := Statements(false, false)
	statements.AddLn("request =>")
	statements.Block(false).
		AddLn("nextedLine1()").
		AddLn("nextedLine2()")
	assertCode(t, statements, expected)
}

func TestStatementsWithScopeNoBlock(t *testing.T) {
	expected := `collection.map { process }`
	statements := Statements(false, false)
	statements.Add("collection.map ").Scope(false).Add("process")
	assertCode(t, statements, expected)
}
