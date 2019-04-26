package scala

import (
	"strings"
)

type ClassDeclaration struct {
	name           string
	params         []Writable
	implicitParams []Writable
	inherits       []string
	modifiers      []string
	attributes     []Writable
	members        []Writable
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

func (self *ClassDeclaration) Inherits(types ...string) *ClassDeclaration {
	self.inherits = append(self.inherits, types...)
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

func (self *ClassDeclaration) WithAttribute(code string) *ClassDeclaration {
	return self.AddAttributes(Attribute(code))
}

func (self *ClassDeclaration) Param(type_ string, name string) *ValDeclaration {
	param := Val(type_, name)
	self.params = append(self.params, param)
	return param
}

func (self *ClassDeclaration) ImplicitParam(type_ string, name string) *ValDeclaration {
	param := Val(type_, name)
	self.implicitParams = append(self.implicitParams, param)
	return param
}

func (self *ClassDeclaration) Def(name string) *MethodDeclaration {
	method := Method(name)
	self.AddDefinitions(method)
	return method
}

func (self *ClassDeclaration) Val(type_ string, name string) *ValDeclaration {
	field := Val(type_, name)
	self.AddDefinitions(field)
	return field
}

func Class(name string) *ClassDeclaration {
	return &ClassDeclaration{
		name:       name,
		params:     []Writable{},
		modifiers:  []string{},
		attributes: []Writable{},
		members:    []Writable{},
	}
}

func (self *ClassDeclaration) WriteCode(writer CodeWriter) {
	if len(self.attributes) > 0 {
		if len(self.attributes) > 0 {
			for i, attribute := range self.attributes {
				if i > 0 {
					writer.Write(" ")
				}
				attribute.WriteCode(writer)
			}
			writer.Eol()
		}
	}

	if len(self.modifiers) > 0 {
		writer.Write(strings.Join(self.modifiers, " ") + " ")
	}

	writer.Write("class "+self.name)

	if len(self.params) > 0 {
		writer.Write("(")
		for i, param := range self.params {
			param.WriteCode(writer)
			if i < len(self.params)-1 {
				writer.Write(", ")
			}
		}
		writer.Write(")")
	}

	if len(self.implicitParams) > 0 {
		writer.Write("(implicit ")
		for i, param := range self.implicitParams {
			param.WriteCode(writer)
			if i < len(self.implicitParams)-1 {
				writer.Write(", ")
			}
		}
		writer.Write(")")
	}

	if len(self.inherits) > 0 {
		writer.Write(" extends "+strings.Join(self.inherits, ", "))
	}

	writer.Begin()
	for index, member := range self.members {
		if index > 0 { writer.Eol() }
		member.WriteCode(writer)
	}
	writer.End()
}

