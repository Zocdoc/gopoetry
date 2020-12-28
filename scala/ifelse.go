package scala

type IfElseCode struct {
	condition       bool
	trueStatements  StatementsGetter
	falseStatements StatementsGetter
}

func If(condition bool) *IfElseCode {
	return &IfElseCode{condition: condition, trueStatements: nil, falseStatements: nil}
}

func Only(statements ...Writable) *IfElseCode {
	return &IfElseCode{condition: false, trueStatements: func() *StatementsDeclaration { return Statements(statements...) }, falseStatements: nil}
}

func OnlyLazy(trueStatements StatementsGetter) *IfElseCode {
	return &IfElseCode{condition: false, trueStatements: trueStatements, falseStatements: nil}
}

func (self *IfElseCode) If(condition bool) *IfElseCode {
	self.condition = condition
	return self
}

func (self *IfElseCode) Then(statements ...Writable) *IfElseCode {
	self.trueStatements = func() *StatementsDeclaration { return Statements(statements...) }
	return self
}

func (self *IfElseCode) Else(statements ...Writable) *IfElseCode {
	self.falseStatements = func() *StatementsDeclaration { return Statements(statements...) }
	return self
}

func (self *IfElseCode) ThenLazy(trueStatements StatementsGetter) *IfElseCode {
	self.trueStatements = trueStatements
	return self
}

func (self *IfElseCode) ElseLazy(falseStatements StatementsGetter) *IfElseCode {
	self.falseStatements = falseStatements
	return self
}

func (self *IfElseCode) WriteCode(writer CodeWriter) {
	if self.condition {
		if self.trueStatements != nil {
			trueStatements := self.trueStatements()
			trueStatements.WriteCode(writer)
		}
	} else {
		if self.falseStatements != nil {
			falseStatements := self.falseStatements()
			falseStatements.WriteCode(writer)
		}
	}
}