package scala

type StatementsDeclaration struct {
	statements []Writable
	block      bool
	scope      bool
}

func (self *StatementsDeclaration) AddCode(code Writable) *StatementsDeclaration {
	self.statements = append(self.statements, code)
	return self
}

func (self *StatementsDeclaration) Add(code string) *StatementsDeclaration {
	self.statements = append(self.statements, Code(code))
	return self
}

func (self *StatementsDeclaration) AddLn(code string) *StatementsDeclaration {
	self.
		AddCode(Code(code)).
		AddCode(Eol())
	return self
}

func (self *StatementsDeclaration) Block(scope bool) *StatementsDeclaration {
	body := Block(scope)
	self.AddCode(body)
	return body
}

func Statements() *StatementsDeclaration {
	return &StatementsDeclaration{statements: []Writable{}, block: false, scope: false}
}

func Block(scope bool) *StatementsDeclaration {
	return &StatementsDeclaration{statements: []Writable{}, block: true, scope: scope}
}

func (self *StatementsDeclaration) WriteCode(writer CodeWriter) {
	if self.block {
		if self.scope {
			writer.Write("{")
			writer.Eol()
		}
		writer.Indent()
	}
	for _, statement := range self.statements {
		statement.WriteCode(writer)
	}
	if self.block {
		writer.UnIndent()
		if self.scope {
			writer.Write("}")
			writer.Eol()
		}
	}
}
