package scala

import (
	"fmt"
	"strings"
)

type FieldKind int

const (
	NONE FieldKind = 0
	VAL  FieldKind = 1
	VAR  FieldKind = 2
)

type FieldDeclaration struct {
	name       string
	type_      string
	kind       FieldKind
	modifiers  []string
	attributes []Writable
	init       Writable
}

func (self *FieldDeclaration) addModifier(modifier string) *FieldDeclaration {
	self.modifiers = append(self.modifiers, modifier)
	return self
}

func (self *FieldDeclaration) Val() *FieldDeclaration {
	self.kind = VAL
	return self
}

func (self *FieldDeclaration) Var() *FieldDeclaration {
	self.kind = VAR
	return self
}

func (self *FieldDeclaration) Override() *FieldDeclaration {
	return self.addModifier("override")
}

func (self *FieldDeclaration) Init(init Writable) *FieldDeclaration {
	self.init = init
	return self
}

func (self *FieldDeclaration) AddAttributes(attributes ...Writable) *FieldDeclaration {
	self.attributes = append(self.attributes, attributes...)
	return self
}

func (self *FieldDeclaration) Attribute(code string) *FieldDeclaration {
	return self.AddAttributes(Attribute(code))
}

func (self *FieldDeclaration) WriteCode(writer CodeWriter) {
	if len(self.attributes) > 0 {
		for i, attribute := range self.attributes {
			if i > 0 {
				writer.Write(" ")
			}
			attribute.WriteCode(writer)
		}
		writer.Write(" ")
	}

	if len(self.modifiers) > 0 {
		writer.Write(strings.Join(self.modifiers, " ") + " ")
	}

	switch self.kind {
	case VAL:
		writer.Write("val ")
	case VAR:
		writer.Write("var ")
	default:
	}

	writer.Write(fmt.Sprintf("%s: %s", self.name, self.type_))

	if self.init != nil {
		writer.Write(" = ")
		self.init.WriteCode(writer)
	}
}

func NewFieldDeclaration(name string, type_ string) *FieldDeclaration {
	return &FieldDeclaration{
		name:       name,
		type_:      type_,
		kind:       NONE,
		modifiers:  []string{},
		attributes: []Writable{},
		init:       nil,
	}
}

func Val(name string, type_ string) *FieldDeclaration {
	return NewFieldDeclaration(name, type_).Val()
}

func Var(name string, type_ string) *FieldDeclaration {
	return NewFieldDeclaration(name, type_).Var()
}
