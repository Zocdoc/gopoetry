package ruby

type MethodDeclaration struct {
	name         string
	args         []Writable
	noArgs       bool
	body         Writable
	paramPerLine bool
}

func (self *MethodDeclaration) AddParams(params ...Writable) *MethodDeclaration {
	self.args = append(self.args, params...)
	return self
}

func (self *MethodDeclaration) NoParams() *MethodDeclaration {
	self.noArgs = true
	return self
}

func (self *MethodDeclaration) ParamPerLine() *MethodDeclaration {
	self.paramPerLine = true
	return self
}

func (self *MethodDeclaration) Body() *StatementsDeclaration {
	body := Statements()
	self.body = body
	return body
}

func (self *MethodDeclaration) Arg(name string) *ArgDeclaration {
	param := Arg(name)
	self.AddParams(param)
	return param
}

func (self *MethodDeclaration) KeywordArg(name string) *ArgDeclaration {
	param := KeywordArg(name)
	self.AddParams(param)
	return param
}

func Method(name string) *MethodDeclaration {
	return &MethodDeclaration{
		name:         name,
		args:         []Writable{},
		body:         nil,
		paramPerLine: false,
	}
}

func (self *MethodDeclaration) WriteCode(writer CodeWriter) {
	writer.Write("def " + self.name)

	if !self.noArgs {
		writer.Write("(")
		if self.paramPerLine {
			writer.Indent()
			writer.Eol()
		}
		for i, param := range self.args {
			param.WriteCode(writer)
			if i < len(self.args)-1 {
				writer.Write(",")
			}
			if self.paramPerLine {
				writer.Eol()
			} else {
				if i < len(self.args)-1 {
					writer.Write(" ")
				}
			}
		}
		if self.paramPerLine {
			writer.UnIndent()
		}
		writer.Write(")")
	}

	writer.Begin("")

	if self.body != nil {
		self.body.WriteCode(writer)
	}

	writer.End()
}
