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

func TestCommentBlockDeclarationMultilineSingleString(t *testing.T) {
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
	fn.AddComments("my comment\nmy second line\nmy third line")
	assertCode(t, fn, expected)
}

func TestCommentBlockEscapesBlockCommentTerminator(t *testing.T) {
	expected := `
/**
 * This comment contains a block terminator: * / but it's escaped
 * Another line with * / in the middle
 */
function myFunction() {
}
`
	fn := DeclareFunction("myFunction")
	fn.AddComments("This comment contains a block terminator: */ but it's escaped", "Another line with */ in the middle")
	assertCode(t, fn, expected)
}
