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

func TestStatementsDef(t *testing.T) {
	expected := `
def the_method()
end
`
	statements := Statements()
	statements.Def("the_method")
	assertCode(t, statements, expected)
}