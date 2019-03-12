package cs

import (
	"fmt"
	"strings"
)

type PropertyDeclaration struct {
	name string
	type_ string
	modifiers []string
	hasGet bool
	hasSet bool
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

func (self *PropertyDeclaration) WriteCode(writer CodeWriter) {
	declaration := fmt.Sprintf("%s %s", self.type_, self.name)
	if len(self.modifiers) > 0 {
		declaration = strings.Join(self.modifiers, " ")+" "+declaration
	}
	writer.Write(declaration)
	writer.Begin()
	if self.hasGet {
		writer.Write("get;")
		writer.Eof()
	}
	if self.hasSet {
		writer.Write("set;")
		writer.Eof()
	}
	writer.End()
}

func Property(type_ string, name string) *PropertyDeclaration {
	return &PropertyDeclaration{
		name: name,
		type_: type_,
		modifiers: []string{},
		hasGet: false,
		hasSet: false,
	}
}