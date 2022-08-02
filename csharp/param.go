package csharp

import "fmt"

type ParamDeclaration struct {
	name         string
	type_        string
	defaultValue Writable
	attributes   []Writable
}

func Param(type_ string, name string) *ParamDeclaration {
	return &ParamDeclaration{
		name:         name,
		type_:        type_,
		defaultValue: nil,
		attributes:   []Writable{},
	}
}

func (self *ParamDeclaration) Default(defaultValue Writable) *ParamDeclaration {
	self.defaultValue = defaultValue
	return self
}

func (self *ParamDeclaration) AddAttributes(attributes ...Writable) *ParamDeclaration {
	self.attributes = append(self.attributes, attributes...)
	return self
}

func (self *ParamDeclaration) WithAttribute(code string) *ParamDeclaration {
	return self.AddAttributes(Attribute(code))
}

func (self *ParamDeclaration) WriteCode(writer CodeWriter) {
	if len(self.attributes) > 0 {
		writer.Write("[")
		for i, attribute := range self.attributes {
			if i > 0 {
				writer.Write(", ")
			}
			attribute.WriteCode(writer)
		}
		writer.Write("]")
		writer.Write(" ")
	}
	writer.Write(fmt.Sprintf("%s %s", self.type_, self.name))
	if self.defaultValue != nil {
		writer.Write(" = ")
		self.defaultValue.WriteCode(writer)
	}
}
