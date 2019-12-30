package ruby

import (
	"fmt"
)

// TODO: New lines between members

type ClassDeclaration struct {
	name       string
	superClass *string
	members    []Writable
}

func (self *ClassDeclaration) Inherits(superClass string) *ClassDeclaration {
	self.superClass = &superClass
	return self
}

func (self *ClassDeclaration) AddMembers(members ...Writable) *ClassDeclaration {
	self.members = append(self.members, members...)
	return self
}

func (self *ClassDeclaration) AddCode(code string) *ClassDeclaration {
	self.members = append(self.members, Code(code))
	return self
}

func (self *ClassDeclaration) Method(name string) *MethodDeclaration {
	method := Method(name)
	self.AddMembers(method)
	return method
}

func (self *ClassDeclaration) Initialize() *MethodDeclaration {
	initialize := Method("initialize")
	self.AddMembers(initialize)
	return initialize
}

func Class(name string) *ClassDeclaration {
	return &ClassDeclaration{
		name:       name,
		superClass: nil,
		members:    []Writable{},
	}
}

func (self *ClassDeclaration) WriteCode(writer CodeWriter) {
	declaration := fmt.Sprintf("class %s", self.name)

	if self.superClass != nil {
		declaration += " < " + *self.superClass
	}

	writer.Begin(declaration)
	for index, member := range self.members {
		if index > 0 { writer.Eol() }
		member.WriteCode(writer)
	}
	writer.End()
}

