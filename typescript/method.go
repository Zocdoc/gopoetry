package typescript

import (
	"strings"
)

// MethodDeclaration declares an method
type MethodDeclaration struct {
	name      string
	returns   string
	modifiers []string
	hasParams bool
	params    []Writable
	body      Writable
	CommentBlockDeclaration
}

// Returns sets the return type of the method
func (method *MethodDeclaration) Returns(returnType string) *MethodDeclaration {
	method.returns = returnType
	return method
}

func (method *MethodDeclaration) addModifier(modifier string) *MethodDeclaration {
	method.modifiers = append(method.modifiers, modifier)
	return method
}

// Private marks the method as private
func (method *MethodDeclaration) Private() *MethodDeclaration {
	return method.addModifier("private")
}

// Public marks the method as public
func (method *MethodDeclaration) Public() *MethodDeclaration {
	return method.addModifier("public")
}

// Async marks the method as async
func (method *MethodDeclaration) Async() *MethodDeclaration {
	return method.addModifier("async")
}

// Static marks the method as static
func (method *MethodDeclaration) Static() *MethodDeclaration {
	return method.addModifier("static")
}

// AddParams adds parameters to the method
func (method *MethodDeclaration) AddParams(params ...Writable) *MethodDeclaration {
	method.params = append(method.params, params...)
	return method
}

// Body starts building the method body. returns BlockDeclaration
func (method *MethodDeclaration) Body(lines ...string) *BlockDeclaration {
	body := Block(lines...)
	method.body = body
	return body
}

// Param adds a parameter to the method
func (method *MethodDeclaration) Param(typeName string, name string) *ParamDeclaration {
	param := Param(typeName, name)
	method.AddParams(param)
	return param
}

// Method starts building a new method. returns MethodDeclaration
func Method(name string) *MethodDeclaration {
	return &MethodDeclaration{
		name:      name,
		returns:   "void",
		modifiers: []string{},
		hasParams: true,
		params:    []Writable{},
		body:      nil,
	}
}

// Constructor starts building a new constructor. returns MethodDeclaration
func Constructor() *MethodDeclaration {
	return &MethodDeclaration{
		name:      "constructor",
		returns:   "",
		modifiers: []string{},
		hasParams: true,
		params:    []Writable{},
		body:      Block(),
	}
}

// WriteCode writes the method to the writer
func (method *MethodDeclaration) WriteCode(writer CodeWriter) {
	method.WriteComments(writer)
	if len(method.modifiers) > 0 {
		writer.Write(strings.Join(method.modifiers, " ") + " ")
	}

	writer.Write(method.name)

	if method.hasParams {
		writer.Write("(")
		for i, param := range method.params {
			param.WriteCode(writer)
			if i < len(method.params)-1 {
				writer.Write(", ")
			}
		}
		writer.Write(")")
	}

	if method.returns != "" {
		writer.Write(": ")
		writer.Write(method.returns)
	}

	if method.body != nil {
		method.body.WriteCode(writer)
	} else {
		writer.Write(";")
		writer.Eol()
	}
}
