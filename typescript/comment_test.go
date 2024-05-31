package typescript

import "testing"

func TestCommentBlockDeclaration(t *testing.T) {
	expected := `
/** my comment */
function myFunction() {
}
`
	fn := DeclareFunction("myFunction")
	fn.AddComments("my comment")
	assertCode(t, fn, expected)
}

func TestCommentBlockDeclarationMultiline(t *testing.T) {
	expected := `
/**
 * my comment
 * my second line
 * my third line
 */
function myFunction() {
}
`
	fn := DeclareFunction("myFunction")
	fn.AddComments("my comment", "my second line", "my third line")
	assertCode(t, fn, expected)
}
