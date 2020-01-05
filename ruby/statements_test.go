package ruby

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

func TestStatementsScope(t *testing.T) {
	expected := `
line1()
  line2()
`
	statements := Statements().AddLn("line1()")
	statements.Scope().AddLn("line2()")

	assertCode(t, statements, expected)
}