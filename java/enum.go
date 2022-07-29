package java

import (
	"fmt"
	"strings"
)


type EnumValue struct {
    name string
    value string
}

type EnumDeclaration struct {
	name        string
	enumMembers []EnumValue
	members     []Writable
	modifiers   []string
}

func (cls *EnumDeclaration) addModifier(modifier string) *EnumDeclaration {
	cls.modifiers = append(cls.modifiers, modifier)
	return cls
}

func Enum(name string) *EnumDeclaration {
	return &EnumDeclaration{
		name:        name,
		enumMembers: make([]EnumValue, 0),
		members:     []Writable{},
	}
}

func (cls *EnumDeclaration) Private() *EnumDeclaration {
	return cls.addModifier("private")
}

func (cls *EnumDeclaration) Public() *EnumDeclaration {
	return cls.addModifier("public")
}

func (cls *EnumDeclaration) AddEnumMembers(enumMemberInCode string, enumMemberStringValue string) *EnumDeclaration {
	cls.enumMembers = append(cls.enumMembers, EnumValue{
        name:  enumMemberInCode,
        value: enumMemberStringValue,
    })
	return cls
}

func (cls *EnumDeclaration) AddMembers(members ...Writable) *EnumDeclaration {
	cls.members = append(cls.members, members...)
	return cls
}

func (cls *EnumDeclaration) Constructor() *MethodDeclaration {
	ctor := Method(cls.name)
	cls.AddMembers(ctor)
	return ctor
}

func (cls *EnumDeclaration) WriteCode(writer CodeWriter) {
	if len(cls.modifiers) > 0 {
		writer.Write(strings.Join(cls.modifiers, " ") + " ")
	}

	declaration := fmt.Sprintf("enum %s", cls.name)
	writer.Write(declaration)

	writer.Begin()
	index := 0
	for _, enum := range cls.enumMembers {
		if index > 0 {
			writer.Eol()
		}

		enumElement := fmt.Sprintf("%s(\"%s\")", enum.name, enum.value)
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
