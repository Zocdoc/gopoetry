package csharp

import (
	"strings"
)

type MethodDeclaration struct {
	name       string
	returns    string
	modifiers  []string
	attributes []Writable
	hasParams  bool
	params     []Writable
	body       Writable
	hasBase    bool
	base       *BaseStatement
	summary    *SummaryDeclaration
}

func (self *MethodDeclaration) Returns(returnType string) *MethodDeclaration {
	self.returns = returnType
	if self.summary != nil {
		self.summary.AddReturnType(returnType)
	}
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

func (self *MethodDeclaration) Static() *MethodDeclaration {
	return self.addModifier("static")
}

func (self *MethodDeclaration) AddParams(params ...Writable) *MethodDeclaration {
	self.params = append(self.params, params...)
	return self
}

func (self *MethodDeclaration) Body(lines ...string) *BlockDeclaration {
	body := Block(lines...)
	self.body = body
	return body
}

func (self *MethodDeclaration) addParamDescription(name string, description string) *MethodDeclaration {
	if self.summary != nil {
		self.summary.AddParam(name, description)
	}
	return self
}

func (self *MethodDeclaration) Param(type_ string, name string) *ParamDeclaration {
	param := Param(type_, name)
	self.AddParams(param)
	self.addParamDescription(name, "")
	return param
}

func (self *MethodDeclaration) ParamWithDescription(type_ string, name string, description string) *ParamDeclaration {
	param := Param(type_, name)
	self.AddParams(param)
	self.addParamDescription(name, description)
	return param
}

func (self *MethodDeclaration) WithBase(args ...string) *MethodDeclaration {
	self.base = Base(args)
	return self
}

func (self *MethodDeclaration) Summary(summary string) *MethodDeclaration {
	self.summary = Summary(summary)
	return self
}

func Method(name string) *MethodDeclaration {
	return &MethodDeclaration{
		name:       name,
		returns:    "void",
		modifiers:  []string{},
		attributes: []Writable{},
		hasParams:  true,
		params:     []Writable{},
		body:       nil,
		hasBase:    false,
		base:       nil,
		summary:    nil,
	}
}

func Constructor(name string) *MethodDeclaration {
	return &MethodDeclaration{
		name:       name,
		returns:    "",
		modifiers:  []string{},
		attributes: []Writable{},
		hasParams:  true,
		params:     []Writable{},
		body:       nil,
		hasBase:    true,
		base:       nil,
		summary:    nil,
	}
}

func Get() *MethodDeclaration {
	return &MethodDeclaration{
		name:       "get",
		returns:    "",
		modifiers:  []string{},
		attributes: []Writable{},
		hasParams:  false,
		params:     []Writable{},
		body:       nil,
		hasBase:    false,
		base:       nil,
		summary:    nil,
	}
}

func Set() *MethodDeclaration {
	return &MethodDeclaration{
		name:       "set",
		returns:    "",
		modifiers:  []string{},
		attributes: []Writable{},
		hasParams:  false,
		params:     []Writable{},
		body:       nil,
		hasBase:    false,
		base:       nil,
		summary:    nil,
	}
}

func (self *MethodDeclaration) WriteCode(writer CodeWriter) {
	if self.summary != nil {
		self.summary.WriteCode(writer)
	}

	if len(self.attributes) > 0 {
		writer.Write("[")
		for i, attribute := range self.attributes {
			if i > 0 {
				writer.Write(", ")
			}
			attribute.WriteCode(writer)
		}
		writer.Write("]")
		writer.Eol()
	}

	if len(self.modifiers) > 0 {
		writer.Write(strings.Join(self.modifiers, " ") + " ")
	}
	if self.returns != "" {
		writer.Write(self.returns)
		writer.Write(" ")
	}
	writer.Write(self.name)

	if self.hasParams {
		writer.Write("(")
		for i, param := range self.params {
			param.WriteCode(writer)
			if i < len(self.params)-1 {
				writer.Write(", ")
			}
		}
		writer.Write(")")

		if self.hasBase && self.base != nil {
			writer.Write(" : ")
			self.base.WriteCode(writer)
		}
	}

	if self.body != nil {
		self.body.WriteCode(writer)
	} else {
		writer.Write(";")
		writer.Eol()
	}
}
