package typescript

import (
	"fmt"
	"strings"
)

// EnumDeclaration declares an enum
type EnumDeclaration struct {
	name      string
	modifiers []string
	members   []Writable
}

func (enum *EnumDeclaration) addModifier(modifier string) *EnumDeclaration {
	enum.modifiers = append(enum.modifiers, modifier)
	return enum
}

// Const makes the enum a constant
func (enum *EnumDeclaration) Const() *EnumDeclaration {
	return enum.addModifier("const")
}

// Exported makes the enum a exported
func (enum *EnumDeclaration) Export() *EnumDeclaration {
	return enum.addModifier("export")
}

// AddMembers adds members to the enum
func (enum *EnumDeclaration) AddMembers(members ...Writable) *EnumDeclaration {
	enum.members = append(enum.members, members...)
	return enum
}

// Member adds a member to the enum. returns EnumMemberDeclaration
func (enum *EnumDeclaration) Member(name string) *EnumMemberDeclaration {
	member := EnumMember(name)
	enum.AddMembers(member)
	return member
}

// Enum starts building a new enum. returns EnumDeclaration
func Enum(name string) *EnumDeclaration {
	return &EnumDeclaration{
		name:      name,
		modifiers: []string{},
		members:   []Writable{},
	}
}

// WriteCode writes the enum to the writer
func (enum *EnumDeclaration) WriteCode(writer CodeWriter) {
	declaration := fmt.Sprintf("enum %s", enum.name)
	if len(enum.modifiers) > 0 {
		declaration = strings.Join(enum.modifiers, " ") + " " + declaration
	}

	writer.Write(declaration)
	writer.Begin()
	for _, member := range enum.members {
		member.WriteCode(writer)
		writer.Write(",")
		writer.Eol()
	}
	writer.End()
}
