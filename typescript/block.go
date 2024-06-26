package typescript

type BlockDeclaration struct {
	lines []Writable
}

// Len returns the number of lines defined in the block
func (b *BlockDeclaration) Len() int {
	return len(b.lines)
}

func (self *BlockDeclaration) AppendCode(code ...Writable) *BlockDeclaration {
	self.lines = append(self.lines, code...)
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
	return &BlockDeclaration{lines: codeLines}
}

func (self *BlockDeclaration) WriteCode(writer CodeWriter) {
	writer.Begin()
	for _, line := range self.lines {
		line.WriteCode(writer)
		writer.Eol()
	}
	writer.End()
}

// OTBSBlock
// https://en.wikipedia.org/wiki/Indentation_style#Variant:_1TBS_.28OTBS.29
type OTBSBlock struct {
	BlockDeclaration
}

func (b *OTBSBlock) WriteCode(writer CodeWriter) {
	writer.OpenBlock()
	for _, line := range b.lines {
		line.WriteCode(writer)
		writer.Eol()
	}
	writer.CloseBlock()
}
