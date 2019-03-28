package csharp

import (
	"fmt"
	"strings"
)

type MethodDeclaration struct {
	name       string
	returns    string
	modifiers  []string
	attributes []Writable
	params     []Writable
	body       []string
}

func (self *MethodDeclaration) Returns(returnType string) *MethodDeclaration {
	self.returns = returnType
	return self
}

func (self *MethodDeclaration) addModifier(modifier string) *MethodDeclaration {
	self.modifiers = append(self.modifiers, modifier)
	return self
}

func (self *MethodDeclaration) Private() *MethodDeclaration {
	return self.addModifier("private")
}

func (self *MethodDeclaration) Public() *MethodDeclaration {
	return self.addModifier("public")
}

func (self *MethodDeclaration) AddAttributes(attributes ...Writable) *MethodDeclaration {
	self.attributes = append(self.attributes, attributes...)
	return self
}

func (self *MethodDeclaration) WithAttribute(code string) *MethodDeclaration {
	return self.AddAttributes(Attribute(code))
}

func (self *MethodDeclaration) Static() *MethodDeclaration {
	return self.addModifier("static")
}

func (self *MethodDeclaration) AddParams(params ...Writable) *MethodDeclaration {
	self.params = append(self.params, params...)
	return self
}

func (self *MethodDeclaration) Body(lines ...string) *MethodDeclaration {
	if self.body == nil {
		self.body = []string{}
	}
	self.body = append(self.body, lines...)
	return self
}

func (self *MethodDeclaration) Param(type_ string, name string) *ParamDeclaration {
	param := Param(type_, name)
	self.AddParams(param)
	return param
}

func Method(name string) *MethodDeclaration {
	return &MethodDeclaration{
		name:      name,
		returns:   "void",
		modifiers: []string{},
		params:    []Writable{},
		body:      nil,
	}
}

func (self *MethodDeclaration) WriteCode(writer CodeWriter) {
	for _, attribute := range self.attributes {
		attribute.WriteCode(writer)
		writer.Eol()
	}

	if len(self.modifiers) > 0 {
		writer.Write(strings.Join(self.modifiers, " ") + " ")
	}
	writer.Write(fmt.Sprintf("%s %s", self.returns, self.name))
	writer.Write("(")
	for i, param := range self.params {
		param.WriteCode(writer)
		if i < len(self.params)-1 {
			writer.Write(", ")
		}
	}
	writer.Write(")")
	if self.body != nil {
		writer.Begin()
		for _, line := range self.body {
			writer.Write(line)
			writer.Eol()
		}
		writer.End()
	} else {
		writer.Write(";")
		writer.Eol()
	}
}
