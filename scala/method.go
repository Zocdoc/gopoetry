package scala

import "strings"

type MethodDeclaration struct {
	name           string
	returns        *string
	modifiers      []string
	attributes     []Writable
	params         []Writable
	noParams       bool
	implicitParams []Writable
	definition     Writable
	paramPerLine   bool
}

func (self *MethodDeclaration) Returns(returnType string) *MethodDeclaration {
	self.returns = &returnType
	return self
}

func (self *MethodDeclaration) addModifier(modifier string) *MethodDeclaration {
	self.modifiers = append(self.modifiers, modifier)
	return self
}

func (self *MethodDeclaration) Private() *MethodDeclaration {
	return self.addModifier("private")
}

func (self *MethodDeclaration) Public() *MethodDeclaration {
	return self.addModifier("public")
}

func (self *MethodDeclaration) Override() *MethodDeclaration {
	return self.addModifier("override")
}

func (self *MethodDeclaration) Async() *MethodDeclaration {
	return self.addModifier("async")
}

func (self *MethodDeclaration) AddAttributes(attributes ...Writable) *MethodDeclaration {
	self.attributes = append(self.attributes, attributes...)
	return self
}

func (self *MethodDeclaration) Attribute(code string) *MethodDeclaration {
	return self.AddAttributes(Attribute(code))
}

func (self *MethodDeclaration) AddParams(params ...Writable) *MethodDeclaration {
	self.params = append(self.params, params...)
	return self
}

func (self *MethodDeclaration) AddImplicitParams(params ...Writable) *MethodDeclaration {
	self.implicitParams = append(self.implicitParams, params...)
	return self
}

func (self *MethodDeclaration) NoParams() *MethodDeclaration {
	self.noParams = true
	return self
}

func (self *MethodDeclaration) ParamPerLine() *MethodDeclaration {
	self.paramPerLine = true
	return self
}

func (self *MethodDeclaration) Define(statements *StatementsDeclaration) *MethodDeclaration {
	if statements != nil {
		self.definition = statements
	} else {
		self.definition = Statements()
	}
	return self
}

func (self *MethodDeclaration) Param(name string, type_ string) *MethodDeclaration {
	param := NewFieldDeclaration(name, type_)
	self.AddParams(param)
	return self
}

func (self *MethodDeclaration) ImplicitParam(name string, type_ string) *MethodDeclaration {
	param := NewFieldDeclaration(name, type_)
	self.AddImplicitParams(param)
	return self
}

func (self *MethodDeclaration) IsConstructor() bool {
	return self.name == ""
}

func Method(name string) *MethodDeclaration {
	return &MethodDeclaration{
		name:           name,
		returns:        nil,
		modifiers:      []string{},
		attributes:     []Writable{},
		params:         []Writable{},
		implicitParams: []Writable{},
		definition:     nil,
		paramPerLine:   false,
	}
}

func Constructor() *MethodDeclaration {
	return Method("")
}

func (self *MethodDeclaration) WriteCode(writer CodeWriter) {
	if len(self.attributes) > 0 {
		if self.IsConstructor() {
			writer.Write(" ")
		}
		for i, attribute := range self.attributes {
			if i > 0 {
				writer.Write(" ")
			}
			attribute.WriteCode(writer)
		}
		if !self.IsConstructor() {
			writer.Eol()
		}
	}

	if len(self.modifiers) > 0 {
		if self.IsConstructor() {
			writer.Write(" ")
		}
		writer.Write(strings.Join(self.modifiers, " "))
		writer.Write(" ")
	}

	if !self.IsConstructor() {
		writer.Write("def " + self.name)
	}

	if !self.noParams {
		writer.Write("(")
		if self.paramPerLine {
			writer.Indent()
			writer.Eol()
		}
		for i, param := range self.params {
			param.WriteCode(writer)
			if i < len(self.params)-1 {
				writer.Write(",")
			}
			if self.paramPerLine {
				writer.Eol()
			} else {
				if i < len(self.params)-1 {
					writer.Write(" ")
				}
			}
		}
		if self.paramPerLine {
			writer.UnIndent()
		}
		writer.Write(")")
	}

	if len(self.implicitParams) > 0 {
		writer.Write("(implicit ")
		for i, param := range self.implicitParams {
			param.WriteCode(writer)
			if i < len(self.implicitParams)-1 {
				writer.Write(", ")
			}
		}
		writer.Write(")")
	}

	if self.returns != nil {
		writer.Write(": ")
		writer.Write(*self.returns)
	}

	if self.definition != nil {
		writer.Write(" = ")
		self.definition.WriteCode(writer)
	} else {
		if !self.IsConstructor() {
			writer.Eol()
		}
	}
}
