package csharp

import (
	"fmt"
)

type InterfaceDeclaration struct {
	name       string
	attributes []Writable
	members    []Writable
}

func (self *InterfaceDeclaration) AddMembers(members ...Writable) *InterfaceDeclaration {
	self.members = append(self.members, members...)
	return self
}

func (self *InterfaceDeclaration) AddAttributes(attributes ...Writable) *InterfaceDeclaration {
	self.attributes = append(self.attributes, attributes...)
	return self
}

func (self *InterfaceDeclaration) WithAttribute(code string) *InterfaceDeclaration {
	return self.AddAttributes(Attribute(code))
}

func (self *InterfaceDeclaration) Method(name string) *MethodDeclaration {
	method := Method(name)
	self.AddMembers(method)
	return method
}

func (self *InterfaceDeclaration) Property(type_ string, name string) *PropertyDeclaration {
	property := Property(type_, name)
	self.AddMembers(property)
	return property
}

func Interface(name string) *InterfaceDeclaration {
	return &InterfaceDeclaration{
		name:       name,
		attributes: []Writable{},
		members:    []Writable{},
	}
}

func (self *InterfaceDeclaration) WriteCode(writer CodeWriter) {
	declaration := fmt.Sprintf("interface %s", self.name)

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
