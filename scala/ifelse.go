package scala

type IfElseCode struct {
	condition       bool
	trueStatements  *StatementsDeclaration
	falseStatements *StatementsDeclaration
}

func If(condition bool) *IfElseCode {
	return &IfElseCode{condition: condition, trueStatements: nil, falseStatements: nil}
}

func (self *IfElseCode) Then(statements ...Writable) *IfElseCode {
	self.trueStatements = Statements(statements...)
	return self
}

func (self *IfElseCode) Else(statements ...Writable) *IfElseCode {
	self.falseStatements = Statements(statements...)
	return self
}

func (self *IfElseCode) WriteCode(writer CodeWriter) {
	if self.condition {
		if self.trueStatements != nil {
			self.trueStatements.WriteCode(writer)
		}
	} else {
		if self.falseStatements != nil {
			self.falseStatements.WriteCode(writer)
		}
	}
}