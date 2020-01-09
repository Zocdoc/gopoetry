package scala

import "strings"

type ExtendsDeclaration struct {
	className string
	params    []string
	withs     []string
	code      Writable
}

func Extends(className string, params ...string) *ExtendsDeclaration {
	return &ExtendsDeclaration{
		className: className,
		params:    params,
	}
}

func (self *ExtendsDeclaration) With(traitName string) *ExtendsDeclaration {
	self.withs = append(self.withs, traitName)
	return self
}

func (self *ExtendsDeclaration) WriteCode(writer CodeWriter) {
	writer.Write("extends " + self.className)
	if self.params != nil && len(self.params) > 0 {
		writer.Write("(")
		writer.Write(strings.Join(self.params, ", "))
		writer.Write(")")
	}

	if len(self.withs) > 0 {
		writer.Write(" with ")
		writer.Write(strings.Join(self.withs, " with "))
	}
}
