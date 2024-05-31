package csharp

import (
	"fmt"
	"strings"
)

type EnumDeclaration struct {
	name       string
	inherits   *string
	modifiers  []string
	attributes []Writable
	members    []Writable
	summary    SummaryDeclaration
}

func (self *EnumDeclaration) addModifier(modifier string) *EnumDeclaration {
	self.modifiers = append(self.modifiers, modifier)
	return self
}

func (self *EnumDeclaration) Private() *EnumDeclaration {
	return self.addModifier("private")
}

func (self *EnumDeclaration) Public() *EnumDeclaration {
	return self.addModifier("public")
}

func (self *EnumDeclaration) Static() *EnumDeclaration {
	return self.addModifier("static")
}

func (self *EnumDeclaration) Inherits(type_ string) *EnumDeclaration {
	self.inherits = &type_
	return self
}

func (self *EnumDeclaration) AddMembers(members ...Writable) *EnumDeclaration {
	self.members = append(self.members, members...)
	return self
}

func (self *EnumDeclaration) AddAttributes(attributes ...Writable) *EnumDeclaration {
	self.attributes = append(self.attributes, attributes...)
	return self
}

func (self *EnumDeclaration) WithAttribute(code string) *EnumDeclaration {
	return self.AddAttributes(Attribute(code))
}

func (self *EnumDeclaration) Member(name string) *EnumMemberDeclaration {
	member := EnumMember(name)
	self.AddMembers(member)
	return member
}

func (self *EnumDeclaration) Summary(summary string) *EnumDeclaration {
	self.summary.AddDescription(summary)
	return self
}

func Enum(name string) *EnumDeclaration {
	return &EnumDeclaration{
		name:       name,
		modifiers:  []string{},
		attributes: []Writable{},
		members:    []Writable{},
		summary:    SummaryDeclaration{},
	}
}

func (self *EnumDeclaration) WriteCode(writer CodeWriter) {
	self.summary.WriteCode(writer)
	declaration := fmt.Sprintf("enum %s", self.name)
	if len(self.modifiers) > 0 {
		declaration = strings.Join(self.modifiers, " ") + " " + declaration
	}

	if self.inherits != nil {
		declaration += fmt.Sprintf(": %s", *self.inherits)
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

	writer.Write(declaration)
	writer.Begin()
	for _, member := range self.members {
		member.WriteCode(writer)
		writer.Write(",")
		writer.Eol()
	}
	writer.End()
}
