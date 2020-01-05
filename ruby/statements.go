package ruby

type StatementsDeclaration struct {
	statements []Writable
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
	self.AddCode(Code(code)).AddCode(Eol())
	return self
}

func (self *StatementsDeclaration) Scope() *StatementsDeclaration {
	scope := &StatementsDeclaration{
		statements: []Writable{},
		scope: true,
	}
	self.AddCode(scope).AddCode(Eol())
	return scope
}

func Statements() *StatementsDeclaration {
	return &StatementsDeclaration{
		statements: []Writable{},
		scope: false,
	}
}

func (self *StatementsDeclaration) WriteCode(writer CodeWriter) {
	if self.scope {
		writer.Indent()
	}
	for _, statement := range self.statements {
		statement.WriteCode(writer)
	}
	if self.scope {
		writer.UnIndent()
	}
}
