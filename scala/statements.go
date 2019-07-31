package scala

type StatementsDeclaration struct {
	statements []Writable
}

func (self *StatementsDeclaration) AppendCode(code Writable) *StatementsDeclaration {
	self.statements = append(self.statements, code)
	return self
}

func (self *StatementsDeclaration) Append(code string) *StatementsDeclaration {
	self.statements = append(self.statements, Code(code))
	return self
}

func (self *StatementsDeclaration) AppendLine(code string) *StatementsDeclaration {
	self.
		AppendCode(Code(code)).
		AppendCode(Eol())
	return self
}

func (self *StatementsDeclaration) Block(lines ...string) *BlockDeclaration {
	body := Block(lines...)
	self.AppendCode(body)
	return body
}

func Statements(lines ...string) *StatementsDeclaration {
	codeLines := []Writable{}
	for _, line := range lines {

		codeLines = append(codeLines, Code(line), Eol())
	}
	return &StatementsDeclaration{ statements: codeLines }
}

func (self *StatementsDeclaration) WriteCode(writer CodeWriter) {
	for _, line := range self.statements {
		line.WriteCode(writer)
	}
}



