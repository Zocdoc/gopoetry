package scala

import (
"fmt"
"strings"
)

type CaseClassDeclaration struct {
	name       string
	modifiers  []string
	attributes []Writable
	fields     []Writable
}

func (self *CaseClassDeclaration) addModifier(modifier string) *CaseClassDeclaration {
	self.modifiers = append(self.modifiers, modifier)
	return self
}

func (self *CaseClassDeclaration) Private() *CaseClassDeclaration {
	return self.addModifier("private")
}

func (self *CaseClassDeclaration) Public() *CaseClassDeclaration {
	return self.addModifier("public")
}

func (self *CaseClassDeclaration) Static() *CaseClassDeclaration {
	return self.addModifier("static")
}

func (self *CaseClassDeclaration) Sealed() *CaseClassDeclaration {
	return self.addModifier("sealed")
}

func (self *CaseClassDeclaration) AddAttributes(attributes ...Writable) *CaseClassDeclaration {
	self.attributes = append(self.attributes, attributes...)
	return self
}

func (self *CaseClassDeclaration) Attribute(code string) *CaseClassDeclaration {
	return self.AddAttributes(Attribute(code))
}

func (self *CaseClassDeclaration) Param(name string, type_ string) *ValDeclaration {
	field := Val(name, type_)
	self.fields = append(self.fields, field)
	return field
}

func CaseClass(name string) *CaseClassDeclaration {
	return &CaseClassDeclaration{
		name:       name,
		modifiers:  []string{},
		attributes: []Writable{},
		fields:     []Writable{},
	}
}

func (self *CaseClassDeclaration) WriteCode(writer CodeWriter) {
	if len(self.attributes) > 0 {
		for i, attribute := range self.attributes {
			if i > 0 {
				writer.Write(" ")
			}
			attribute.WriteCode(writer)
		}
		writer.Eol()
	}

	declaration := fmt.Sprintf("case class %s", self.name)
	if len(self.modifiers) > 0 {
		declaration = strings.Join(self.modifiers, " ") + " " + declaration
	}

	writer.Write(declaration)

	writer.Write("(")
	writer.Indent()
	if len(self.fields) > 0 {
		writer.Eol()
		for i, param := range self.fields {
			param.WriteCode(writer)
			if i < len(self.fields)-1 {
				writer.Write(", ")
			}
			writer.Eol()
		}
	}
	writer.Unindent()
	writer.Write(")")
	writer.Eol()
}
