package scala

import (
	"fmt"
)

type ValDeclaration struct {
	name       string
	type_      string
	attributes []Writable
	init       Writable
}

func (self *ValDeclaration) Init(init Writable) *ValDeclaration {
	self.init = init
	return self
}

func (self *ValDeclaration) AddAttributes(attributes ...Writable) *ValDeclaration {
	self.attributes = append(self.attributes, attributes...)
	return self
}

func (self *ValDeclaration) WithAttribute(code string) *ValDeclaration {
	return self.AddAttributes(Attribute(code))
}

func (self *ValDeclaration) WriteCode(writer CodeWriter) {
	if len(self.attributes) > 0 {
		for i, attribute := range self.attributes {
			if i > 0 {
				writer.Write(" ")
			}
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

func Val(name string, type_ string) *ValDeclaration {
	return &ValDeclaration{
		name:       name,
		type_:      type_,
		attributes: []Writable{},
		init:       nil,
	}
}
