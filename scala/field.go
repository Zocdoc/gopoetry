package scala

import (
	"fmt"
)

type FieldDeclaration struct {
	name       string
	type_      string
	attributes []Writable
	init       Writable
}

func (self *FieldDeclaration) Init(init Writable) *FieldDeclaration {
	self.init = init
	return self
}

func (self *FieldDeclaration) AddAttributes(attributes ...Writable) *FieldDeclaration {
	self.attributes = append(self.attributes, attributes...)
	return self
}

func (self *FieldDeclaration) WithAttribute(code string) *FieldDeclaration {
	return self.AddAttributes(Attribute(code))
}

func (self *FieldDeclaration) WriteCode(writer CodeWriter) {
	if len(self.attributes) > 0 {
		for i, attribute := range self.attributes {
			if i > 0 {
				writer.Write(" ")
			}
			writer.Write("@")
			attribute.WriteCode(writer)
		}
		writer.Write(" ")
	}

	writer.Write(fmt.Sprintf("%s: %s", self.name, self.type_))

	if self.init != nil {
		writer.Write(" = ")
		self.init.WriteCode(writer)
	}
}

func Field(type_ string, name string) *FieldDeclaration {
	return &FieldDeclaration{
		name:       name,
		type_:      type_,
		attributes: []Writable{},
		init:       nil,
	}
}
