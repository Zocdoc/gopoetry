package scala

import (
	"strings"
)

type MethodDeclaration struct {
	name           string
	returns        string
	modifiers      []string
	attributes     []Writable
	params         []Writable
	implicitParams []Writable
	body           Writable
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

func (self *MethodDeclaration) Async() *MethodDeclaration {
	return self.addModifier("async")
}

func (self *MethodDeclaration) AddAttributes(attributes ...Writable) *MethodDeclaration {
	self.attributes = append(self.attributes, attributes...)
	return self
}

func (self *MethodDeclaration) WithAttribute(code string) *MethodDeclaration {
	return self.AddAttributes(Attribute(code))
}

func (self *MethodDeclaration) AddParams(params ...Writable) *MethodDeclaration {
	self.params = append(self.params, params...)
	return self
}

func (self *MethodDeclaration) AddImplicitParams(params ...Writable) *MethodDeclaration {
	self.implicitParams = append(self.implicitParams, params...)
	return self
}

func (self *MethodDeclaration) Body(lines ...string) *BlockDeclaration {
	body := Block(lines...)
	self.body = body
	return body
}

func (self *MethodDeclaration) Param(type_ string, name string) *FieldDeclaration {
	param := Field(type_, name)
	self.AddParams(param)
	return param
}

func (self *MethodDeclaration) ImplicitParam(type_ string, name string) *FieldDeclaration {
	param := Field(type_, name)
	self.AddImplicitParams(param)
	return param
}

func Method(name string) *MethodDeclaration {
	return &MethodDeclaration{
		name:           name,
		returns:        "Unit",
		modifiers:      []string{},
		attributes:     []Writable{},
		params:         []Writable{},
		implicitParams: []Writable{},
		body:           nil,
	}
}

func (self *MethodDeclaration) WriteCode(writer CodeWriter) {
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
	writer.Write("def "+self.name)

	writer.Write("(")
	for i, param := range self.params {
		param.WriteCode(writer)
		if i < len(self.params)-1 {
			writer.Write(", ")
		}
	}
	writer.Write(")")

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

		writer.Write(": ")
	writer.Write(self.returns)

	if self.body != nil {
		self.body.WriteCode(writer)
	} else {
		writer.Eol()
	}
}
