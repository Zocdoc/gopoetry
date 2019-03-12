package cs

type Writable interface {
	WriteCode(writer CodeWriter)
}

type codeDefinition struct {
	code string
}

func (self *codeDefinition) WriteCode(writer CodeWriter) {
	writer.Write(self.code)
}

func Code(code string) *codeDefinition {
	return &codeDefinition{code: code}
}

func C(code string) *codeDefinition {
	return Code(code)
}