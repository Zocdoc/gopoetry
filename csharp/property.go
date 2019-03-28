package csharp

import (
	"fmt"
	"strings"
)

type PropertyDeclaration struct {
	name       string
	type_      string
	modifiers  []string
	attributes []Writable
	hasGet     bool
	hasSet     bool
}

func (self *PropertyDeclaration) addModifier(modifier string) *PropertyDeclaration {
	self.modifiers = append(self.modifiers, modifier)
	return self
}

func (self *PropertyDeclaration) Private() *PropertyDeclaration {
	return self.addModifier("private")
}

func (self *PropertyDeclaration) Public() *PropertyDeclaration {
	return self.addModifier("public")
}

func (self *PropertyDeclaration) Get() *PropertyDeclaration {
	self.hasGet = true
	return self
}

func (self *PropertyDeclaration) Set() *PropertyDeclaration {
	self.hasSet = true
	return self
}

func (self *PropertyDeclaration) AddAttributes(attributes ...Writable) *PropertyDeclaration {
	self.attributes = append(self.attributes, attributes...)
	return self
}

func (self *PropertyDeclaration) WithAttribute(code string) *PropertyDeclaration {
	return self.AddAttributes(Attribute(code))
}

func Property(type_ string, name string) *PropertyDeclaration {
	return &PropertyDeclaration{
		name:      name,
		type_:     type_,
		modifiers: []string{},
		hasGet:    false,
		hasSet:    false,
	}
}

func (self *PropertyDeclaration) WriteCode(writer CodeWriter) {
	for _, attribute := range self.attributes {
		attribute.WriteCode(writer)
		writer.Eol()
	}
	declaration := fmt.Sprintf("%s %s", self.type_, self.name)
	if len(self.modifiers) > 0 {
		declaration = strings.Join(self.modifiers, " ") + " " + declaration
	}
	writer.Write(declaration)
	writer.Begin()
	if self.hasGet {
		writer.Write("get;")
		writer.Eol()
	}
	if self.hasSet {
		writer.Write("set;")
		writer.Eol()
	}
	writer.End()
}
