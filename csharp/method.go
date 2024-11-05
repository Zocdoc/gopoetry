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
	summary    SummaryDeclaration
	expression *SingleLineExpressionDeclaration
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

func (self *MethodDeclaration) Param(type_ string, name string) *ParamDeclaration {
	param := Param(type_, name)
	self.AddParams(param)
	self.summary.AddParam(name, "")
	return param
}

func (self *MethodDeclaration) ParamWithDescription(type_ string, name string, description string) *ParamDeclaration {
	param := Param(type_, name)
	self.AddParams(param)
	self.summary.AddParam(name, description)
	return param
}

func (self *MethodDeclaration) ParamWithDescriptionAlreadyEscaped(type_ string, name string, description string) *ParamDeclaration {
	param := Param(type_, name)
	self.AddParams(param)
	self.summary.AddParamAlreadyEscaped(name, description)
	return param
}

func (self *MethodDeclaration) WithBase(args ...string) *MethodDeclaration {
	self.base = Base(args)
	return self
}

func (self *MethodDeclaration) ReturnsSummary(returns string) *MethodDeclaration {
	self.summary.AddReturns(returns)
	return self
}

func (self *MethodDeclaration) Summary(summary string) *MethodDeclaration {
	self.summary.AddDescription(summary)
	return self
}

func (self *MethodDeclaration) ExpressionBodiedMember(expression string) *MethodDeclaration {
	self.expression = SingleLineExpression(expression)
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
		summary:    SummaryDeclaration{},
		expression: nil,
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
		summary:    SummaryDeclaration{},
		expression: nil,
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
		summary:    SummaryDeclaration{},
		expression: nil,
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
		summary:    SummaryDeclaration{},
		expression: nil,
	}
}

func Init() *MethodDeclaration {
	return &MethodDeclaration{
		name:       "init",
		returns:    "",
		modifiers:  []string{},
		attributes: []Writable{},
		hasParams:  false,
		params:     []Writable{},
		body:       nil,
		hasBase:    false,
		base:       nil,
		summary:    SummaryDeclaration{},
		expression: nil,
	}
}

func (self *MethodDeclaration) WriteCode(writer CodeWriter) {
	self.summary.WriteCode(writer)

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

	if self.expression != nil {
		writer.Lambda()
		self.expression.WriteCode(writer)
		writer.Write(";")
		writer.Eol()
	} else if self.body != nil {
		self.body.WriteCode(writer)
	} else {
		writer.Write(";")
		writer.Eol()
	}
}
