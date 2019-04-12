package csharp

type BlockDeclaration struct {
	lines       []Writable
}

func (self *BlockDeclaration) AppendCode(code Writable) *BlockDeclaration {
	self.lines = append(self.lines, code)
	return self
}

func (self *BlockDeclaration) Append(code string) *BlockDeclaration {
	self.lines = append(self.lines, Code(code))
	return self
}

func Block(lines ...string) *BlockDeclaration {
	codeLines := []Writable{}
	for _, line := range lines {
		codeLines = append(codeLines, Code(line))
	}
	return &BlockDeclaration{ lines: codeLines }
}

func (self *BlockDeclaration) WriteCode(writer CodeWriter) {
	writer.Begin()
	for _, line := range self.lines {
		line.WriteCode(writer)
		writer.Eol()
	}
	writer.End()
}


