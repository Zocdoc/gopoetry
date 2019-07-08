package typescript

import (
	"fmt"
	"strings"
)

// ParamDeclaration declares an parameter
type ParamDeclaration struct {
	name         string
	typeName     string
	defaultValue Writable
	modifiers    []string
}

// Param creates a new ParamDeclaration
func Param(typeName string, name string) *ParamDeclaration {
	return &ParamDeclaration{
		name:         name,
		typeName:     typeName,
		defaultValue: nil,
		modifiers:    []string{},
	}
}

// Default sets the default value of the param
func (param *ParamDeclaration) Default(defaultValue Writable) *ParamDeclaration {
	param.defaultValue = defaultValue
	return param
}

// Private marks the parameter as private. Only for constructors
func (param *ParamDeclaration) Private() *ParamDeclaration {
	return param.addModifier("private")
}

// Public marks the parameter as public. Only for constructors
func (param *ParamDeclaration) Public() *ParamDeclaration {
	return param.addModifier("public")
}

// Readonly marks the parameter as readonly. Only for constructors
func (param *ParamDeclaration) Readonly() *ParamDeclaration {
	return param.addModifier("readonly")
}

func (param *ParamDeclaration) addModifier(modifier string) *ParamDeclaration {
	param.modifiers = append(param.modifiers, modifier)
	return param
}

// WriteCode writes the param to the writer
func (param *ParamDeclaration) WriteCode(writer CodeWriter) {
	if len(param.modifiers) > 0 {
		writer.Write(strings.Join(param.modifiers, " ") + " ")
	}

	writer.Write(fmt.Sprintf("%s: %s", param.name, param.typeName))
	if param.defaultValue != nil {
		writer.Write(" = ")
		param.defaultValue.WriteCode(writer)
	}
}
