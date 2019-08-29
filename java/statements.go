package java

import (
	"gopoetry/util"
)

type StatementsDeclaration struct {
	statements []util.Writable
	isBlock    bool
}

func (self *StatementsDeclaration) AppendCode(code util.Writable) *StatementsDeclaration {
	self.statements = append(self.statements, code)
	return self
}

func (self *StatementsDeclaration) Append(code string) *StatementsDeclaration {
	self.statements = append(self.statements, Code(code))
	return self
}

func (self *StatementsDeclaration) Line(code string) *StatementsDeclaration {
	self.
		AppendCode(Code(code)).
		AppendCode(Eol())
	return self
}

func (self *StatementsDeclaration) Lines(lines ...string) *StatementsDeclaration {
	for _, line := range lines {
		self.Line(line)
	}
	return self
}

func (self *StatementsDeclaration) Block() *StatementsDeclaration {
	body := Block()
	self.AppendCode(body)
	return body
}

func Statements() *StatementsDeclaration {
	return &StatementsDeclaration{statements: []util.Writable{}, isBlock: false}
}

func Block() *StatementsDeclaration {
	return &StatementsDeclaration{statements: []util.Writable{}, isBlock: true}
}

func (self *StatementsDeclaration) WriteCode(writer util.CodeWriter) {
	if self.isBlock {
		writer.Begin()
	}
	for _, line := range self.statements {
		line.WriteCode(writer)
	}
	if self.isBlock {
		writer.End()
	}
}
