package csharp

import (
	"fmt"
	"strings"
)

type ClassDeclaration struct {
	name       string
	inherits   []string
	modifiers  []string
	attributes []Writable
	members    []Writable
	summary    SummaryDeclaration
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

func (self *ClassDeclaration) Sealed() *ClassDeclaration {
	return self.addModifier("sealed")
}

func (self *ClassDeclaration) Inherits(types ...string) *ClassDeclaration {
	self.inherits = append(self.inherits, types...)
	return self
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

func (self *ClassDeclaration) Constructor() *MethodDeclaration {
	ctor := Constructor(self.name)
	self.AddMembers(ctor)
	return ctor
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

func (self *ClassDeclaration) Summary(summary string) *ClassDeclaration {
	self.summary.AddDescription(summary)
	return self
}

func Class(name string) *ClassDeclaration {
	return &ClassDeclaration{
		name:       name,
		modifiers:  []string{},
		attributes: []Writable{},
		members:    []Writable{},
		summary:    SummaryDeclaration{},
	}
}

func (self *ClassDeclaration) WriteCode(writer CodeWriter) {
	self.summary.WriteCode(writer)

	declaration := fmt.Sprintf("class %s", self.name)
	if len(self.modifiers) > 0 {
		declaration = strings.Join(self.modifiers, " ") + " " + declaration
	}

	if len(self.inherits) > 0 {
		declaration += " : " + strings.Join(self.inherits, ", ")
	}

	if len(self.attributes) > 0 {
		writer.Write("[")
		for i, attribute := range self.attributes {
			if i > 0 {
				writer.Write(", ")
			}
			attribute.WriteCode(writer)
		}
		writer.Write("]")
		writer.Eol()
	}

	writer.Write(declaration)
	writer.Begin()
	for index, member := range self.members {
		if index > 0 {
			writer.Eol()
		}
		member.WriteCode(writer)
	}
	writer.End()
}
