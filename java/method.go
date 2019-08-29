package java

import (
	"gopoetry/util"
	"strings"
)

type MethodDeclaration struct {
	name       string
	returns    *string
	modifiers  []string
	attributes []util.Writable
	params     []util.Writable
	definition util.Writable
}

func (self *MethodDeclaration) Returns(returnType string) *MethodDeclaration {
	self.returns = &returnType
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

func (self *MethodDeclaration) Static() *MethodDeclaration {
	return self.addModifier("static")
}

func (self *MethodDeclaration) AddAttributes(attributes ...util.Writable) *MethodDeclaration {
	self.attributes = append(self.attributes, attributes...)
	return self
}

func (self *MethodDeclaration) Attribute(code string) *MethodDeclaration {
	return self.AddAttributes(Attribute(code))
}

func (self *MethodDeclaration) AddParams(params ...util.Writable) *MethodDeclaration {
	self.params = append(self.params, params...)
	return self
}

func (self *MethodDeclaration) Define() *StatementsDeclaration {
	statements := Statements()
	self.definition = statements
	return statements
}

func (self *MethodDeclaration) Param(name string, type_ string) *ParamDeclaration {
	param := Param(name, type_)
	self.AddParams(param)
	return param
}

func Method(name string) *MethodDeclaration {
	return &MethodDeclaration{
		name:       name,
		returns:    nil,
		modifiers:  []string{},
		attributes: []util.Writable{},
		params:     []util.Writable{},
		definition: nil,
	}
}

func (self *MethodDeclaration) WriteCode(writer util.CodeWriter) {
	if len(self.attributes) > 0 {
		for i, attribute := range self.attributes {
			if i > 0 {
				writer.Write(" ")
			}
			attribute.WriteCode(writer)
		}
		writer.Eol()
	}

	if len(self.modifiers) > 0 {
		writer.Write(strings.Join(self.modifiers, " "))
		writer.Write(" ")
	}

	if self.returns != nil {
		writer.Write(*self.returns)
		writer.Write(" ")
	}

	writer.Write(self.name)

	writer.Write("(")
	for i, param := range self.params {
		param.WriteCode(writer)
		if i < len(self.params)-1 {
			writer.Write(", ")
		}
	}
	writer.Write(")")

	if self.definition != nil {
		self.definition.WriteCode(writer)
	} else {
		writer.Eol()
	}
}
