package csharp

import (
	"fmt"
	"strings"
)

type ClassDeclaration struct {
	name       string
	modifiers  []string
	attributes []Writable
	members    []Writable
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

func (self *ClassDeclaration) Static() *ClassDeclaration {
	return self.addModifier("static")
}

func (self *ClassDeclaration) AddMembers(members ...Writable) *ClassDeclaration {
	self.members = append(self.members, members...)
	return self
}

func (self *ClassDeclaration) AddAttributes(attributes ...Writable) *ClassDeclaration {
	self.attributes = append(self.attributes, attributes...)
	return self
}

func (self *ClassDeclaration) WithAttribute(code string) *ClassDeclaration {
	return self.AddAttributes(Attribute(code))
}

func (self *ClassDeclaration) Method(name string) *MethodDeclaration {
	method := Method(name)
	self.AddMembers(method)
	return method
}

func (self *ClassDeclaration) Field(type_ string, name string) *FieldDeclaration {
	field := Field(type_, name)
	self.AddMembers(field)
	return field
}

func (self *ClassDeclaration) Property(type_ string, name string) *PropertyDeclaration {
	property := Property(type_, name)
	self.AddMembers(property)
	return property
}

func Class(name string) *ClassDeclaration {
	return &ClassDeclaration{
		name:       name,
		modifiers:  []string{},
		attributes: []Writable{},
		members:    []Writable{},
	}
}

func (self *ClassDeclaration) WriteCode(writer CodeWriter) {
	declaration := fmt.Sprintf("class %s", self.name)
	if len(self.modifiers) > 0 {
		declaration = strings.Join(self.modifiers, " ") + " " + declaration
	}

	for _, attribute := range self.attributes {
		attribute.WriteCode(writer)
		writer.Eol()
	}

	writer.Write(declaration)
	writer.Begin()
	for _, member := range self.members {
		member.WriteCode(writer)
	}
	writer.End()
}
