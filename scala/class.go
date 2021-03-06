package scala

import (
	"strings"
)

type ClassDeclaration struct {
	name       string
	extends    []string
	modifiers  []string
	attributes []Writable
	definition *StatementsDeclaration
	ctor       Writable
	isObject   bool
}

func (self *ClassDeclaration) addModifier(modifier string) *ClassDeclaration {
	self.modifiers = append(self.modifiers, modifier)
	return self
}

func (self *ClassDeclaration) Private() *ClassDeclaration {
	return self.addModifier("private")
}

func (self *ClassDeclaration) Public() *ClassDeclaration {
	return self.addModifier("public")
}

func (self *ClassDeclaration) Sealed() *ClassDeclaration {
	return self.addModifier("sealed")
}

func (self *ClassDeclaration) Case() *ClassDeclaration {
	return self.addModifier("case")
}

func (self *ClassDeclaration) Abstract() *ClassDeclaration {
	return self.addModifier("abstract")
}

func (self *ClassDeclaration) Extends(types ...string) *ClassDeclaration {
	self.extends = append(self.extends, types...)
	return self
}

func (self *ClassDeclaration) Define(block bool) *StatementsDeclaration {
	self.definition = Statements(block, true)
	return self.definition
}

func (self *ClassDeclaration) AddAttributes(attributes ...Writable) *ClassDeclaration {
	self.attributes = append(self.attributes, attributes...)
	return self
}

func (self *ClassDeclaration) Attribute(code string) *ClassDeclaration {
	return self.AddAttributes(Attribute(code))
}

func (self *ClassDeclaration) Contructor() *MethodDeclaration {
	ctor := Method("")
	self.ctor = ctor
	return ctor
}

func Class(name string) *ClassDeclaration {
	return &ClassDeclaration{
		name:       name,
		modifiers:  []string{},
		attributes: []Writable{},
		definition: nil,
		ctor:       nil,
		isObject:   false,
	}
}

func Object(name string) *ClassDeclaration {
	return &ClassDeclaration{
		name:       name,
		modifiers:  []string{},
		attributes: []Writable{},
		definition: nil,
		ctor:       nil,
		isObject:   true,
	}
}

func (self *ClassDeclaration) WriteCode(writer CodeWriter) {
	if len(self.attributes) > 0 {
		for _, attribute := range self.attributes {
			attribute.WriteCode(writer)
		}
		writer.Eol()
	}

	if len(self.modifiers) > 0 {
		writer.Write(strings.Join(self.modifiers, " ") + " ")
	}

	if self.isObject {
		writer.Write("object ")
	} else {
		writer.Write("class ")
	}
	writer.Write(self.name)

	if self.ctor != nil {
		self.ctor.WriteCode(writer)
	}

	if len(self.extends) > 0 {
		writer.Write(" extends " + strings.Join(self.extends, " with "))
	}

	if self.definition != nil {
		writer.Write(" ")
		self.definition.WriteCode(writer)
	} else {
		writer.Eol()
	}
}
