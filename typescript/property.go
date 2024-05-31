package typescript

import (
	"fmt"
	"strings"
)

// PropertyDeclaration declares a property
type PropertyDeclaration struct {
	name        string
	typeName    string
	modifiers   []string
	optional    bool
	initializer string
	CommentBlockDeclaration
}

func (prop *PropertyDeclaration) addModifier(modifier string) *PropertyDeclaration {
	prop.modifiers = append(prop.modifiers, modifier)
	return prop
}

// Private marks the property as private
func (prop *PropertyDeclaration) Private() *PropertyDeclaration {
	return prop.addModifier("private")
}

// Readonly marks the property as readonly
func (prop *PropertyDeclaration) Readonly() *PropertyDeclaration {
	return prop.addModifier("readonly")
}

// Public marks the property as public
func (prop *PropertyDeclaration) Public() *PropertyDeclaration {
	return prop.addModifier("public")
}

// Static marks the property as static. Can only be used on classes, not interfaces
func (prop *PropertyDeclaration) Static() *PropertyDeclaration {
	return prop.addModifier("static")
}

// Optional marks the property as optional
func (prop *PropertyDeclaration) Optional() *PropertyDeclaration {
	prop.optional = true
	return prop
}

// Initializer sets the properties initializer
func (prop *PropertyDeclaration) Initializer(initializer string) *PropertyDeclaration {
	prop.initializer = initializer
	return prop
}

// Property starts building a new property. returns PropertyDeclaration
func Property(typeName string, name string) *PropertyDeclaration {
	return &PropertyDeclaration{
		name:      name,
		typeName:  typeName,
		modifiers: []string{},
	}
}

// WriteCode writes the property to the writer
func (prop *PropertyDeclaration) WriteCode(writer CodeWriter) {
	prop.WriteComments(writer)
	optionalStr := ""
	if prop.optional {
		optionalStr = "?"
	}

	declaration := fmt.Sprintf("%s%s: %s", prop.name, optionalStr, prop.typeName)
	if len(prop.modifiers) > 0 {
		declaration = strings.Join(prop.modifiers, " ") + " " + declaration
	}

	if prop.initializer != "" {
		declaration = declaration + " = " + prop.initializer
	}

	declaration = declaration + ";"

	writer.Write(declaration)
	writer.Eol()
}
