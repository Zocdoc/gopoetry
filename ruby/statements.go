package ruby

type StatementsDeclaration struct {
	statements []Writable
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

func (self *StatementsDeclaration) Def(name string) *MethodDeclaration {
	method := Method(name)
	self.AddCode(method)
	return method
}

func Statements() *StatementsDeclaration {
	return &StatementsDeclaration{statements: []Writable{}}
}

func (self *StatementsDeclaration) WriteCode(writer CodeWriter) {
	for _, statement := range self.statements {
		statement.WriteCode(writer)
	}
}
