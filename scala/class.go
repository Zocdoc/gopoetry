package scala

import (
	"strings"
)

type ClassDeclaration struct {
	name       string
	extends    *ExtendsDeclaration
	modifiers  []string
	attributes []Writable
	members    *StatementsDeclaration
	ctor       *CtorDeclaration
	isObject   bool
	isCase     bool
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
	self.isCase = true
	if !self.isObject {
		self.addConstructor()
		self.ctor.ForCaseClass()
	}
	return self
}

func (self *ClassDeclaration) Abstract() *ClassDeclaration {
	return self.addModifier("abstract")
}

func (self *ClassDeclaration) Extends(baseClassName string, params ...string) *ClassDeclaration {
	self.extends = Extends(baseClassName, params...)
	return self
}

func (self *ClassDeclaration) With(traitName string) *ClassDeclaration {
	self.extends.With(traitName)
	return self
}

func (self *ClassDeclaration) addConstructor() {
	if self.ctor == nil {
		self.ctor = Constructor()
	}
}

func (self *ClassDeclaration) Members(statements ...Writable) *ClassDeclaration {
	self.members = Scope(statements...)
	return self
}

func (self *ClassDeclaration) MembersInline(statements ...Writable) *ClassDeclaration {
	self.members = ScopeInline(statements...)
	return self
}

func (self *ClassDeclaration) Add(members ...Writable) *ClassDeclaration {
	if self.members == nil {
		self.members = Scope()
	}
	self.members.Add(members...)
	return self
}

func (self *ClassDeclaration) AddAttributes(attributes ...Writable) *ClassDeclaration {
	self.attributes = append(self.attributes, attributes...)
	return self
}

func (self *ClassDeclaration) Attribute(code string) *ClassDeclaration {
	return self.AddAttributes(Attribute(code))
}

func (self *ClassDeclaration) Constructor(ctor *CtorDeclaration) *ClassDeclaration {
	self.ctor = ctor
	if self.isCase && !self.isObject {
		self.ctor.ForCaseClass()
	}
	return self
}

func Class(name string) *ClassDeclaration {
	return &ClassDeclaration{
		name:       name,
		modifiers:  []string{},
		attributes: []Writable{},
		members:    nil,
		ctor:       nil,
		isObject:   false,
		extends:    nil,
	}
}

func Object(name string) *ClassDeclaration {
	return &ClassDeclaration{
		name:       name,
		modifiers:  []string{},
		attributes: []Writable{},
		members:    nil,
		ctor:       nil,
		isObject:   true,
		extends:    nil,
	}
}

func CaseClass(name string) *ClassDeclaration {
	return Class(name).Case()
}

func CaseObject(name string) *ClassDeclaration {
	return Object(name).Case()
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

	if self.isCase {
		writer.Write("case ")
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

	if self.extends != nil {
		writer.Write(" ")
		self.extends.WriteCode(writer)
	}

	if self.members != nil {
		writer.Write(" ")
		self.members.WriteCode(writer)
	} else {
		writer.Eol()
	}
}
