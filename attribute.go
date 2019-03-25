package poetrycs

type AttributeDeclaration struct {
	code Writable
}

func Attribute(code string) *AttributeDeclaration {
	return &AttributeDeclaration{
		code: Code(code),
	}
}

func (self *AttributeDeclaration) WriteCode(writer CodeWriter) {
	writer.Write("[")
	self.code.WriteCode(writer)
	writer.Write("]")
}
