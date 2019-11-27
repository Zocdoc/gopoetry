package typescript

import (
	"fmt"
	"strings"
)

// InterfaceDeclaration declares an interface
type InterfaceDeclaration struct {
	name      string
	extends   []string
	modifiers []string
	members   []Writable
}

func (iface *InterfaceDeclaration) addModifier(modifier string) *InterfaceDeclaration {
	iface.modifiers = append(iface.modifiers, modifier)
	return iface
}

// Export marks the interface as exported
func (iface *InterfaceDeclaration) Export() *InterfaceDeclaration {
	return iface.addModifier("export")
}

// Extends adds additional types to extend the interface
func (iface *InterfaceDeclaration) Extends(types ...string) *InterfaceDeclaration {
	iface.extends = append(iface.extends, types...)
	return iface
}

// AddMembers adds additional members to the interface
func (iface *InterfaceDeclaration) AddMembers(members ...Writable) *InterfaceDeclaration {
	iface.members = append(iface.members, members...)
	return iface
}

// Method starts building a new method on the interface. returns MethodDeclaration
func (iface *InterfaceDeclaration) Method(name string) *MethodDeclaration {
	method := Method(name)
	iface.AddMembers(method)
	return method
}

// Property starts building a new property on the interface. returns PropertyDeclaration
func (iface *InterfaceDeclaration) Property(typeName string, name string) *PropertyDeclaration {
	property := Property(typeName, name)
	iface.AddMembers(property)
	return property
}

// Interface starts building a new interface. returns InterfaceDeclaration
func Interface(name string) *InterfaceDeclaration {
	return &InterfaceDeclaration{
		name:    name,
		members: []Writable{},
	}
}

// WriteCode writes the interface to the writer
func (iface *InterfaceDeclaration) WriteCode(writer CodeWriter) {
	declaration := fmt.Sprintf("interface %s", iface.name)

	if len(iface.modifiers) > 0 {
		declaration = strings.Join(iface.modifiers, " ") + " " + declaration
	}

	if len(iface.extends) > 0 {
		declaration += " extends " + strings.Join(iface.extends, ", ")
	}

	writer.Write(declaration)
	writer.Begin()
	for _, member := range iface.members {
		member.WriteCode(writer)
	}
	writer.End()
}
