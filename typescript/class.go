package typescript

import (
	"fmt"
	"strings"
)

// ClassDeclaration declares a class
type ClassDeclaration struct {
	name       string
	implements []string
	extends    []string
	modifiers  []string
	members    []Writable
}

func (cls *ClassDeclaration) addModifier(modifier string) *ClassDeclaration {
	cls.modifiers = append(cls.modifiers, modifier)
	return cls
}

// Private marks a class as private
func (cls *ClassDeclaration) Private() *ClassDeclaration {
	return cls.addModifier("private")
}

// Exported marks a class as exported
func (cls *ClassDeclaration) Exported() *ClassDeclaration {
	return cls.addModifier("export")
}

// Public marks a class as public
func (cls *ClassDeclaration) Public() *ClassDeclaration {
	return cls.addModifier("public")
}

// Implements adds types that the classes implements
func (cls *ClassDeclaration) Implements(types ...string) *ClassDeclaration {
	cls.implements = append(cls.implements, types...)
	return cls
}

// Extends adds types that the classes extends
func (cls *ClassDeclaration) Extends(types ...string) *ClassDeclaration {
	cls.extends = append(cls.extends, types...)
	return cls
}

// AddMembers adds methods or properties to a class
func (cls *ClassDeclaration) AddMembers(members ...Writable) *ClassDeclaration {
	cls.members = append(cls.members, members...)
	return cls
}

// Method starts building a new method on the class. returns MethodDeclaration
func (cls *ClassDeclaration) Method(name string) *MethodDeclaration {
	method := Method(name)
	cls.AddMembers(method)
	return method
}

// Constructor starts building a constructor on the class. returns MethodDeclaration
func (cls *ClassDeclaration) Constructor() *MethodDeclaration {
	ctor := Constructor()
	cls.AddMembers(ctor)
	return ctor
}

// Property adds a property to the class. returns PropertyDeclaration
func (cls *ClassDeclaration) Property(typeName string, name string) *PropertyDeclaration {
	property := Property(typeName, name)
	cls.AddMembers(property)
	return property
}

// Class starts building a new class. returns ClassDeclaration
func Class(name string) *ClassDeclaration {
	return &ClassDeclaration{
		name:      name,
		modifiers: []string{},
		members:   []Writable{},
	}
}

// WriteCode writes the class to the writer
func (cls *ClassDeclaration) WriteCode(writer CodeWriter) {
	declaration := fmt.Sprintf("class %s", cls.name)
	if len(cls.modifiers) > 0 {
		declaration = strings.Join(cls.modifiers, " ") + " " + declaration
	}

	if len(cls.extends) > 0 {
		declaration += " extends " + strings.Join(cls.extends, ", ")
	}

	if len(cls.implements) > 0 {
		declaration += " implements " + strings.Join(cls.implements, ", ")
	}

	writer.Write(declaration)
	writer.Begin()
	for index, member := range cls.members {
		if index > 0 {
			writer.Eol()
		}
		member.WriteCode(writer)
	}
	writer.End()
}
