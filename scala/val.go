package scala

import (
	"fmt"
	"strings"
)

type ValDeclaration struct {
	name       string
	type_      string
	modifiers  []string
	attributes []Writable
	init       Writable
}

func (self *ValDeclaration) addModifier(modifier string) *ValDeclaration {
	self.modifiers = append(self.modifiers, modifier)
	return self
}

func (self *ValDeclaration) Val() *ValDeclaration {
	return self.addModifier("val")
}

func (self *ValDeclaration) Var() *ValDeclaration {
	return self.addModifier("var")
}

func (self *ValDeclaration) Override() *ValDeclaration {
	return self.addModifier("override")
}

func (self *ValDeclaration) Init(init Writable) *ValDeclaration {
	self.init = init
	return self
}

func (self *ValDeclaration) AddAttributes(attributes ...Writable) *ValDeclaration {
	self.attributes = append(self.attributes, attributes...)
	return self
}

func (self *ValDeclaration) Attribute(code string) *ValDeclaration {
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

	if len(self.modifiers) > 0 {
		writer.Write(strings.Join(self.modifiers, " ") + " ")
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
