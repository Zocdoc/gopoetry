package scala

import (
	"strings"
)

type ClassDeclaration struct {
	name       string
	extends    []string
	modifiers  []string
	attributes []Writable
	members    []Writable
	ctor       Writable
	isObject   bool
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

func (self *ClassDeclaration) Sealed() *ClassDeclaration {
	return self.addModifier("sealed")
}

func (self *ClassDeclaration) Case() *ClassDeclaration {
	return self.addModifier("case")
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

func (self *ClassDeclaration) Contructor() *MethodDeclaration {
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
		name:       name,
		modifiers:  []string{},
		attributes: []Writable{},
		members:    []Writable{},
		ctor:       nil,
		isObject:   false,
	}
}

func Object(name string) *ClassDeclaration {
	return &ClassDeclaration{
		name:       name,
		modifiers:  []string{},
		attributes: []Writable{},
		members:    []Writable{},
		ctor:       nil,
		isObject:   true,
	}
}

func (self *ClassDeclaration) WriteCode(writer CodeWriter) {
	if len(self.attributes) > 0 {
		for _, attribute := range self.attributes {
			attribute.WriteCode(writer)
		}
		writer.Eol()
	}

	if len(self.modifiers) > 0 {
		writer.Write(strings.Join(self.modifiers, " ") + " ")
	}

	if self.isObject {
		writer.Write("object ")
	} else {
		writer.Write("class ")
	}
	writer.Write(self.name)

	if self.ctor != nil {
		self.ctor.WriteCode(writer)
	}

	if len(self.extends) > 0 {
		writer.Write(" extends " + strings.Join(self.extends, ", "))
	}

	if len(self.members) > 0 {
		writer.Write(" ")
		WriteMembers(writer, self.members)
	} else {
		writer.Eol()
	}
}
