package typescript

// IfStatement represents an if/else statement block
type IfStatement struct {
	condition Writable
	ifBlock   OTBSBlock
	elseBlock *OTBSBlock
}

func If(cond Writable) *IfStatement {
	return &IfStatement{condition: cond}
}

func (ifB *IfStatement) IfBlockAppend(lines ...Writable) *IfStatement {
	ifB.ifBlock.lines = append(ifB.ifBlock.lines, lines...)
	return ifB
}

func (ifB *IfStatement) ElseBlockAppend(lines ...Writable) *IfStatement {
	if ifB.elseBlock == nil {
		ifB.elseBlock = &OTBSBlock{}
	}
	ifB.elseBlock.lines = append(ifB.elseBlock.lines, lines...)
	return ifB
}

func (ifS *IfStatement) WriteCode(writer CodeWriter) {
	writer.Write("if (")
	ifS.condition.WriteCode(writer)
	writer.Write(") ")
	ifS.ifBlock.WriteCode(writer)

	if ifS.elseBlock != nil {
		writer.Write(" else ")
		ifS.elseBlock.WriteCode(writer)
	}
}
