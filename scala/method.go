package scala

import "strings"

type MethodDeclaration struct {
	name               string
	returns            *string
	modifiers          []string
	attributes         []Writable
	params             []Writable
	noParams           bool
	implicitParams     []Writable
	body               *StatementsDeclaration
	paramPerLine       bool
	implicitsOnNewLine bool
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
	self.attributes = append(self.attributes, filterNils(attributes)...)
	return self
}

func (self *MethodDeclaration) Attribute(code string) *MethodDeclaration {
	return self.AddAttributes(Attribute(code))
}

func (self *MethodDeclaration) AddParams(params ...Writable) *MethodDeclaration {
	self.params = append(self.params, filterNils(params)...)
	return self
}

func (self *MethodDeclaration) AddImplicitParams(params ...Writable) *MethodDeclaration {
	self.implicitParams = append(self.implicitParams, filterNils(params)...)
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

func (self *MethodDeclaration) ImplicitsOnNewLine() *MethodDeclaration {
	self.implicitsOnNewLine = true
	return self
}

func (self *MethodDeclaration) ImplicitsOnSingleLine() *MethodDeclaration {
	self.implicitsOnNewLine = false
	return self
}

func (self *MethodDeclaration) Body(statements ...Writable) *MethodDeclaration {
	self.body = Scope(statements...)
	return self
}

func (self *MethodDeclaration) BodyInline(statements ...Writable) *MethodDeclaration {
	self.body = Statements(statements...)
	return self
}

func (self *MethodDeclaration) AddStatements(statements ...Writable) *MethodDeclaration {
	if self.body == nil {
		self.body = Scope()
	}
	self.body.Add(statements...)
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

func Def(name string) *MethodDeclaration {
	return &MethodDeclaration{
		name:           name,
		returns:        nil,
		modifiers:      []string{},
		attributes:     []Writable{},
		params:         []Writable{},
		implicitParams: []Writable{},
		body:           nil,
		paramPerLine:   false,
	}
}

func (self *MethodDeclaration) WriteCode(writer CodeWriter) {
	if len(self.attributes) > 0 {
		for i, attribute := range self.attributes {
			if i > 0 {
				writer.Write(" ")
			}
			attribute.WriteCode(writer)
		}
		writer.Eol()
	}

	if len(self.modifiers) > 0 {
		writer.Write(strings.Join(self.modifiers, " "))
		writer.Write(" ")
	}

	writer.Write("def " + self.name)

	if !self.noParams {
		writer.Write("(")
		if self.paramPerLine {
			writer.Indent()
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
			writer.UnIndent()
		}
		writer.Write(")")
	}

	if len(self.implicitParams) > 0 {
		writer.Write("(")

		if self.implicitsOnNewLine {
			writer.Eol()
			writer.Indent()
			writer.Indent()
		}

		writer.Write("implicit ")
		for i, param := range self.implicitParams {
			param.WriteCode(writer)
			if i < len(self.implicitParams)-1 {
				writer.Write(", ")
			}
		}

		if self.implicitsOnNewLine {
			writer.Eol()
			writer.UnIndent()
			writer.UnIndent()
		}

		writer.Write(")")
	}

	if self.returns != nil {
		writer.Write(": ")
		writer.Write(*self.returns)
	}

	if self.body != nil {
		writer.Write(" = ")
		self.body.WriteCode(writer)
	} else {
		writer.Eol()
	}
}
