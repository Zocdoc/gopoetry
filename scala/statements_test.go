package scala

import (
	"gotest.tools/assert"
	"strings"
	"testing"
)

var block = Block
var scope = Scope
var code = Code
var line = Line

func TestStatementsSimple(t *testing.T) {
	expected := `
line1()
line2()
`
	statements :=
		Statements(
			line("line1()"),
			line("line2()"),
		)

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
	statements :=
		Statements(
			line("line1()"),
			code("line2 "),
			Scope(
				line("nextedLine1()"),
				line("nextedLine2()"),
			),
		)
	assertCode(t, statements, expected)
}

func TestStatementsWithBlockNoScope(t *testing.T) {
	expected := `
request =>
  nextedLine1()
  nextedLine2()
`
	statements := Statements(
		Line("request =>"),
		Block(
			Line("nextedLine1()"),
			Line("nextedLine2()"),
		),
	)
	assertCode(t, statements, expected)
}

func TestStatementsWithScopeNoBlock(t *testing.T) {
	expected := `collection.map { process }`
	statements := Statements(
		Code("collection.map "),
		ScopeInline(Code("process")),
	)
	assertCode(t, statements, expected)
}

var expectedCode = `
params match {
  case Failure(ex) => Future { BadRequest }
  case Success(params) =>
    val (param1, param2) = params
    val result = api.makeCall(param1, param2)
    result.map(_.toResult.toPlay).recover { case _: Exception => InternalServerError }
}
`

func TestBla(t *testing.T) {
	lambda := Statements(
		code("params match "),
		scope(
			line("case Failure(ex) => Future { BadRequest }"),
			line("case Success(params) =>"),
			block(
				line("val (param1, param2) = params"),
				line("val result = api.makeCall(param1, param2)"),
				line("result.map(_.toResult.toPlay).recover { case _: Exception => InternalServerError }"),
			),
		),
	)
	writer := CreateWriter()
	lambda.WriteCode(&writer)
	code := writer.Code()
	assert.Equal(t, code, strings.TrimPrefix(expectedCode, "\n"))
}
