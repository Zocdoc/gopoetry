package cs

import (
	"fmt"
	"strings"
)

type ClassDeclaration struct {
	name string
	modifiers []string
	members []Writable
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

func (self *ClassDeclaration) Members(members ...Writable) *ClassDeclaration {
	self.members = append(self.members, members...)
	return self
}

func Class(name string) *ClassDeclaration {
	return &ClassDeclaration{
		name: name,
		modifiers: []string{},
		members: []Writable{},
	}
}

func (self *ClassDeclaration) WriteCode(writer CodeWriter) {
	declaration := fmt.Sprintf("class %s", self.name)
	if len(self.modifiers) > 0 {
		declaration = strings.Join(self.modifiers, " ")+" "+declaration
	}
	writer.Write(declaration)
	writer.Begin()
	for _, member := range self.members {
		member.WriteCode(writer)
	}
	writer.End()
}