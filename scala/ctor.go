package scala

import "strings"

type CtorDeclaration struct {
	modifiers      []string
	attributes     []Writable
	params         []Writable
	implicitParams []Writable
	paramPerLine   bool
	forCaseClass   bool
}

func (self *CtorDeclaration) addModifier(modifier string) *CtorDeclaration {
	self.modifiers = append(self.modifiers, modifier)
	return self
}

func (self *CtorDeclaration) Private() *CtorDeclaration {
	return self.addModifier("private")
}

func (self *CtorDeclaration) Public() *CtorDeclaration {
	return self.addModifier("public")
}

func (self *CtorDeclaration) Override() *CtorDeclaration {
	return self.addModifier("override")
}

func (self *CtorDeclaration) AddAttributes(attributes ...Writable) *CtorDeclaration {
	self.attributes = append(self.attributes, filterNils(attributes)...)
	return self
}

func (self *CtorDeclaration) Attribute(code string) *CtorDeclaration {
	return self.AddAttributes(Attribute(code))
}

func (self *CtorDeclaration) AddParams(params ...Writable) *CtorDeclaration {
	self.params = append(self.params, filterNils(params)...)
	return self
}

func (self *CtorDeclaration) AddImplicitParams(params ...Writable) *CtorDeclaration {
	self.implicitParams = append(self.implicitParams, filterNils(params)...)
	return self
}

func (self *CtorDeclaration) Param(name string, type_ string) *CtorDeclaration {
	param := NewFieldDeclaration(name, type_)
	self.AddParams(param)
	return self
}

func (self *CtorDeclaration) Val(name string, type_ string) *CtorDeclaration {
	param := Val(name, type_)
	self.AddParams(param)
	return self
}

func (self *CtorDeclaration) Var(name string, type_ string) *CtorDeclaration {
	param := Var(name, type_)
	self.AddParams(param)
	return self
}

func (self *CtorDeclaration) ImplicitParam(name string, type_ string) *CtorDeclaration {
	param := NewFieldDeclaration(name, type_)
	self.AddImplicitParams(param)
	return self
}

func (self *CtorDeclaration) ParamPerLine() *CtorDeclaration {
	self.paramPerLine = true
	return self
}

func (self *CtorDeclaration) ForCaseClass() *CtorDeclaration {
	self.forCaseClass = true
	return self
}

func Constructor() *CtorDeclaration {
	return &CtorDeclaration{
		modifiers:      []string{},
		attributes:     []Writable{},
		params:         []Writable{},
		implicitParams: []Writable{},
	}
}

func (self *CtorDeclaration) WriteCode(writer CodeWriter) {
	if len(self.attributes) > 0 {
		for _, attribute := range self.attributes {
			writer.Write(" ")
			attribute.WriteCode(writer)
		}
	}

	if len(self.modifiers) > 0 {
		writer.Write(" ")
		writer.Write(strings.Join(self.modifiers, " "))
		writer.Write(" ")
	}

	if self.forCaseClass || len(self.params) > 0 || len(self.attributes) > 0 || len(self.modifiers) > 0 {
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
}