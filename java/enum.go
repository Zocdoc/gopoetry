package java

import (
	"fmt"
	"strings"
)

// ClassDeclaration declares a class
type EnumDeclaration struct {
	name        string
	enumMembers []string
	members     []Writable
	ctor        Writable
}

// AddMembers adds methods to the enum
func (cls *EnumDeclaration) AddEnumMembers(enumMembers ...string) *EnumDeclaration {
	cls.enumMembers = append(cls.enumMembers, enumMembers...)
	return cls
}

func (cls *EnumDeclaration) AddMembers(members ...Writable) *EnumDeclaration {
	cls.members = append(cls.members, members...)
	return cls
}

func (cls *EnumDeclaration) Constructor() *MethodDeclaration {
	ctor := Method("")
	cls.AddMembers(ctor)
	return ctor
}

// WriteCode writes the class to the writer
func (cls *EnumDeclaration) WriteCode(writer CodeWriter) {
	declaration := fmt.Sprintf("public enum %s", cls.name)

	writer.Write(declaration)
	writer.Begin()
	for index, member := range cls.enumMembers {
		if index > 0 {
			writer.Eol()
		}
		enumElement := fmt.Sprintf("%s(\"%s\")", member, strings.ToLower(member))
		writer.Write(enumElement)
		if index < len(cls.members)-1 {
			writer.Write(",")
		} else {
			writer.Write(";")
		}
		writer.Eol()
	}

	cls.ctor.WriteCode(writer)

	for _, member := range cls.members {
		member.WriteCode(writer)
	}

	writer.End()
}
