package scala

import (
"fmt"
"strings"
)

type TraitDeclaration struct {
	name       string
	inherits   []string
	attributes []Writable
	members    []Writable
}

func (self *TraitDeclaration) Inherits(types ...string) *TraitDeclaration {
	self.inherits = append(self.inherits, types...)
	return self
}

func (self *TraitDeclaration) AddDefinitions(members ...Writable) *TraitDeclaration {
	self.members = append(self.members, members...)
	return self
}

func (self *TraitDeclaration) AddAttributes(attributes ...Writable) *TraitDeclaration {
	self.attributes = append(self.attributes, attributes...)
	return self
}

func (self *TraitDeclaration) WithAttribute(code string) *TraitDeclaration {
	return self.AddAttributes(Attribute(code))
}

func (self *TraitDeclaration) Def(name string) *MethodDeclaration {
	method := Method(name)
	self.AddDefinitions(method)
	return method
}

func Trait(name string) *TraitDeclaration {
	return &TraitDeclaration{
		name:       name,
		attributes: []Writable{},
		members:    []Writable{},
	}
}

func (self *TraitDeclaration) WriteCode(writer CodeWriter) {
	declaration := fmt.Sprintf("trait %s", self.name)

	if len(self.inherits) > 0 {
		declaration += ": "+strings.Join(self.inherits, ", ")
	}

	for _, attribute := range self.attributes {
		attribute.WriteCode(writer)
		writer.Eol()
	}

	writer.Write(declaration)
	writer.Write(" ")
	writer.Begin()
	for _, member := range self.members {
		member.WriteCode(writer)
	}
	writer.End()
}

