package java

import (
	"fmt"
	"gopoetry/util"
	"strings"
)

// ClassDeclaration declares a class
type EnumDeclaration struct {
	name        string
	enumMembers map[string]string
	members     []util.Writable
	modifiers   []string
}

func (cls *EnumDeclaration) addModifier(modifier string) *EnumDeclaration {
	cls.modifiers = append(cls.modifiers, modifier)
	return cls
}

func Enum(name string) *EnumDeclaration {
	return &EnumDeclaration{
		name:        name,
		enumMembers: make(map[string]string),
		members:     []util.Writable{},
	}
}

func (cls *EnumDeclaration) Private() *EnumDeclaration {
	return cls.addModifier("private")
}

func (cls *EnumDeclaration) Public() *EnumDeclaration {
	return cls.addModifier("public")
}

// AddMembers adds methods to the enum
func (cls *EnumDeclaration) AddEnumMembers(enumMemberInCode string, enumMemberStringValue string) *EnumDeclaration {
	cls.enumMembers[enumMemberInCode] = enumMemberStringValue
	return cls
}

func (cls *EnumDeclaration) AddMembers(members ...util.Writable) *EnumDeclaration {
	cls.members = append(cls.members, members...)
	return cls
}

func (cls *EnumDeclaration) Constructor() *MethodDeclaration {
	ctor := Method(cls.name)
	cls.AddMembers(ctor)
	return ctor
}

// WriteCode writes the class to the writer
func (cls *EnumDeclaration) WriteCode(writer util.CodeWriter) {
	if len(cls.modifiers) > 0 {
		writer.Write(strings.Join(cls.modifiers, " ") + " ")
	}

	declaration := fmt.Sprintf("enum %s", cls.name)
	writer.Write(declaration)

	writer.Begin()
	index := 0
	for memberNameInCode, memberNameAsString := range cls.enumMembers {
		if index > 0 {
			writer.Eol()
		}

		enumElement := fmt.Sprintf("%s(\"%s\")", memberNameInCode, strings.ToLower(memberNameAsString))
		writer.Write(enumElement)
		if index < len(cls.enumMembers)-1 {
			writer.Write(",")
		} else {
			writer.Write(";")
		}
		writer.Eol()
		index++
	}

	for _, member := range cls.members {
		writer.Eol()
		member.WriteCode(writer)
	}

	writer.End()
}
