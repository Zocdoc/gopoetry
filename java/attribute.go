package java

import "github.com/zocdoc/gopoetry/util"

type AttributeDeclaration struct {
	code util.Writable
}

func Attribute(code string) *AttributeDeclaration {
	return &AttributeDeclaration{
		code: Code(code),
	}
}

func (self *AttributeDeclaration) WriteCode(writer util.CodeWriter) {
	writer.Write("@")
	self.code.WriteCode(writer)
}
