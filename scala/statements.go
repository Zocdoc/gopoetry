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

func (self *StatementsDeclaration) Scope(block bool) *StatementsDeclaration {
	body := Scope(block)
	self.AddCode(body)
	return body
}

func (self *StatementsDeclaration) Def(name string) *MethodDeclaration {
	method := Method(name)
	self.AddCode(method)
	return method
}

func (self *StatementsDeclaration) Val(name string, type_ string) *FieldDeclaration {
	field := Val(name, type_)
	self.AddCode(field)
	return field
}

func (self *StatementsDeclaration) Var(name string, type_ string) *FieldDeclaration {
	field := Var(name, type_)
	self.AddCode(field)
	return field
}

func Statements(block bool, scope bool) *StatementsDeclaration {
	return &StatementsDeclaration{statements: []Writable{}, block: block, scope: scope}
}

func Block(scope bool) *StatementsDeclaration {
	return &StatementsDeclaration{statements: []Writable{}, block: true, scope: scope}
}

func Scope(block bool) *StatementsDeclaration {
	return &StatementsDeclaration{statements: []Writable{}, block: block, scope: true}
}

func (self *StatementsDeclaration) WriteCode(writer CodeWriter) {
	if self.scope {
		writer.Write("{")
		if self.block {
			writer.Eol()
		} else {
			writer.Write(" ")
		}
	}
	if self.block {
		writer.Indent()
	}
	for _, statement := range self.statements {
		statement.WriteCode(writer)
	}
	if self.block {
		writer.UnIndent()
	}
	if self.scope {
		if !self.block {
			writer.Write(" ")
		}
		writer.Write("}")
		writer.Eol()
	}
}
