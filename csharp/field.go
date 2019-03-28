package csharp

import (
	"fmt"
	"strings"
)

type FieldDeclaration struct {
	name      string
	type_     string
	modifiers []string
	init      Writable
}

func (self *FieldDeclaration) addModifier(modifier string) *FieldDeclaration {
	self.modifiers = append(self.modifiers, modifier)
	return self
}

func (self *FieldDeclaration) Private() *FieldDeclaration {
	return self.addModifier("private")
}

func (self *FieldDeclaration) Public() *FieldDeclaration {
	return self.addModifier("public")
}

func (self *FieldDeclaration) Static() *FieldDeclaration {
	return self.addModifier("static")
}

func (self *FieldDeclaration) Init(init Writable) *FieldDeclaration {
	self.init = init
	return self
}

func (self *FieldDeclaration) WriteCode(writer CodeWriter) {
	declaration := fmt.Sprintf("%s %s", self.type_, self.name)
	if len(self.modifiers) > 0 {
		declaration = strings.Join(self.modifiers, " ") + " " + declaration
	}
	writer.Write(declaration)
	if self.init != nil {
		writer.Write(" = ")
		self.init.WriteCode(writer)
	}
	writer.Write(";")
}

func Field(type_ string, name string) *FieldDeclaration {
	return &FieldDeclaration{
		name:      name,
		type_:     type_,
		modifiers: []string{},
		init:      nil,
	}
}
