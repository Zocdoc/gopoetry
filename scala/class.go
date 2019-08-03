package scala

import (
	"strings"
)

type ClassDeclaration struct {
	name           string
	extends        []string
	modifiers      []string
	attributes     []Writable
	members        []Writable
	ctor           Writable
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

func (self *ClassDeclaration) Extends(types ...string) *ClassDeclaration {
	self.extends = append(self.extends, types...)
	return self
}

func (self *ClassDeclaration) AddDefinitions(members ...Writable) *ClassDeclaration {
	self.members = append(self.members, members...)
	return self
}

func (self *ClassDeclaration) AddAttributes(attributes ...Writable) *ClassDeclaration {
	self.attributes = append(self.attributes, attributes...)
	return self
}

func (self *ClassDeclaration) Attribute(code string) *ClassDeclaration {
	return self.AddAttributes(Attribute(code))
}

func (self *ClassDeclaration) Ctor() *MethodDeclaration {
	ctor := Method("")
	self.ctor = ctor
	return ctor
}

func (self *ClassDeclaration) Def(name string) *MethodDeclaration {
	method := Method(name)
	self.AddDefinitions(method)
	return method
}

func (self *ClassDeclaration) Val(name string, type_ string) *ValDeclaration {
	field := Val(name, type_)
	self.AddDefinitions(field)
	return field
}

func Class(name string) *ClassDeclaration {
	return &ClassDeclaration{
		name:           name,
		modifiers:      []string{},
		attributes:     []Writable{},
		members:        []Writable{},
		ctor:           nil,
	}
}

func (self *ClassDeclaration) WriteCode(writer CodeWriter) {
	if len(self.attributes) > 0 {
		for i, attribute := range self.attributes {
			if i > 0 {
				writer.Write(" ")
			}
			attribute.WriteCode(writer)
		}
		writer.Eol()
	}

	if len(self.modifiers) > 0 {
		writer.Write(strings.Join(self.modifiers, " ") + " ")
	}

	writer.Write("class "+self.name)

	if(self.ctor != nil) {
		self.ctor.WriteCode(writer)
	}

	if len(self.extends) > 0 {
		writer.Write(" extends "+strings.Join(self.extends, ", "))
	}

	writer.Write(" ")
	writer.Begin()
	for index, member := range self.members {
		if index > 0 { writer.Eol() }
		member.WriteCode(writer)
	}
	writer.End()
}

