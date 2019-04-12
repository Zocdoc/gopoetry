package csharp

import "fmt"

type ParamDeclaration struct {
	name  string
	type_ string
	defaultValue Writable
}

func Param(type_ string, name string) *ParamDeclaration {
	return &ParamDeclaration{
		name:  name,
		type_: type_,
		defaultValue: nil,
	}
}

func (self *ParamDeclaration) Default(defaultValue Writable) *ParamDeclaration {
	self.defaultValue = defaultValue
	return self
}

func (self *ParamDeclaration) WriteCode(writer CodeWriter) {
	writer.Write(fmt.Sprintf("%s %s", self.type_, self.name))
	if self.defaultValue != nil {
		writer.Write(" = ")
		self.defaultValue.WriteCode(writer)
	}
}
