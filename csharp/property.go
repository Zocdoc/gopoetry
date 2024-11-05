package csharp

import (
	"fmt"
	"strings"
)

type PropertyDeclaration struct {
	name       string
	type_      string
	modifiers  []string
	attributes []Writable
	getter     *MethodDeclaration
	setter     *MethodDeclaration
	init       Writable
	summary    SummaryDeclaration
	expression *SingleLineExpressionDeclaration
}

func (self *PropertyDeclaration) addModifier(modifier string) *PropertyDeclaration {
	self.modifiers = append(self.modifiers, modifier)
	return self
}

func (self *PropertyDeclaration) Private() *PropertyDeclaration {
	return self.addModifier("private")
}

func (self *PropertyDeclaration) Public() *PropertyDeclaration {
	return self.addModifier("public")
}

func (self *PropertyDeclaration) Get() *MethodDeclaration {
	self.getter = Get()
	return self.getter
}

func (self *PropertyDeclaration) Set() *MethodDeclaration {
	self.setter = Set()
	return self.setter
}

func (self *PropertyDeclaration) Init() *MethodDeclaration {
	self.setter = Init()
	return self.setter
}

func (self *PropertyDeclaration) WithGet() *PropertyDeclaration {
	self.Get()
	return self
}

func (self *PropertyDeclaration) WithSet() *PropertyDeclaration {
	self.Set()
	return self
}

func (self *PropertyDeclaration) WithInit() *PropertyDeclaration {
	self.Init()
	return self
}

func (self *PropertyDeclaration) AddAttributes(attributes ...Writable) *PropertyDeclaration {
	self.attributes = append(self.attributes, attributes...)
	return self
}

func (self *PropertyDeclaration) WithAttribute(code string) *PropertyDeclaration {
	return self.AddAttributes(Attribute(code))
}

func (self *PropertyDeclaration) Initialize(init Writable) *PropertyDeclaration {
	self.init = init
	return self
}

func (self *PropertyDeclaration) ExpressionBodiedMember(expression string) *PropertyDeclaration {
	self.expression = SingleLineExpression(expression)
	return self
}

func (self *PropertyDeclaration) Summary(summary string) *PropertyDeclaration {
	self.summary.AddDescription(summary)
	return self
}

func Property(type_ string, name string) *PropertyDeclaration {
	return &PropertyDeclaration{
		name:      name,
		type_:     type_,
		modifiers: []string{},
		getter:    nil,
		setter:    nil,
		summary:   SummaryDeclaration{},
	}
}

func (self *PropertyDeclaration) WriteCode(writer CodeWriter) {
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
	declaration := fmt.Sprintf("%s %s", self.type_, self.name)
	if len(self.modifiers) > 0 {
		declaration = strings.Join(self.modifiers, " ") + " " + declaration
	}
	writer.Write(declaration)

	if self.expression != nil {
		writer.Lambda()
		self.expression.WriteCode(writer)
		writer.Write(";")
		writer.Eol()
	} else {

		inlineBlock := (self.getter == nil || self.getter.IsSimpleGetSetOrInit()) &&
			(self.setter == nil || self.setter.IsSimpleGetSetOrInit())

		if inlineBlock {
			writer.Write(" { ")

			if self.getter != nil {
				self.getter.WriteCode(writer)
			}
			if self.setter != nil {
				if self.getter != nil {
					writer.Write(" ")
				}
				self.setter.WriteCode(writer)
			}

			writer.Write(" }")

			if self.init != nil {
				writer.Write(" = ")
				self.init.WriteCode(writer)
				writer.Write(";")
			}

			writer.Eol()
		} else {
			writer.Begin()

			if self.getter != nil {
				self.getter.WriteCode(writer)
			}
			if self.setter != nil {
				self.setter.WriteCode(writer)
			}

			writer.End()

			if self.init != nil {
				writer.Write(" = ")
				self.init.WriteCode(writer)
				writer.Write(";")
				writer.Eol()
			}
		}
	}
}
