package scala

type IfElseCode struct {
	condition       bool
	trueStatements  Writable
	falseStatements Writable
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

func Only(statements ...Writable) *IfElseCode {
	return &IfElseCode{condition: false, trueStatements: Statements(statements...), falseStatements: nil}
}

func (self *IfElseCode) If(condition bool) *IfElseCode {
	self.condition = condition
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