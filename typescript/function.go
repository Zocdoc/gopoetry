package typescript

type Function struct {
	Name            string
	body            BlockDeclaration
	params          []ParamDeclaration
	returnType      Writable
	typeConstructor Writable
	exported        bool
}

func (f *Function) Export() *Function {
	f.exported = true
	return f
}

func (f *Function) AddParams(params ...ParamDeclaration) *Function {
	f.params = append(f.params, params...)
	return f
}

// WriteCode implements Writable
func (f *Function) WriteCode(writer CodeWriter) {
	if f.exported {
		writer.Write("export ")
	}

	writer.Write("function " + f.Name)

	if f.typeConstructor != nil {
		writer.Write("<")
		f.typeConstructor.WriteCode(writer)
		writer.Write(">")
	}

	writer.Write("(")
	for i, param := range f.params {
		param.WriteCode(writer)
		if i < len(f.params)-1 {
			writer.Write(", ")
		}
	}
	writer.Write(")")

	if f.returnType != nil {
		writer.Write(": ")
		f.returnType.WriteCode(writer)
	}

	writer.Write(" ")
	writer.OpenBlock()
	for _, ln := range f.body.lines {
		ln.WriteCode(writer)
		writer.Eol()
	}
	writer.End()
}
