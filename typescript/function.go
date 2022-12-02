package typescript

// FunctionDecl represents a top level declaration.
type FunctionDecl struct {
	Name     string
	exported bool
	Function
}

// Export exports the function declaration
func (f *FunctionDecl) Export() *FunctionDecl {
	f.exported = true
	return f
}

// AddParams add parameters to the function signiture
func (f *FunctionDecl) AddParams(params ...ParamDeclaration) *FunctionDecl {
	f.params = append(f.params, params...)
	return f
}

// WriteCode implements Writable
func (f *FunctionDecl) WriteCode(writer CodeWriter) {
	if f.exported {
		writer.Write("export ")
	}

	if f.async {
		writer.Write("async ")
	}

	writer.Write("function " + f.Name)
	f.writeCallSig(writer)
	writer.Write(" ")
	f.writeBody(writer)
	writer.Eol()
}

// Function represents the function as an expression
type Function struct {
	async           bool
	arrowSyntax     bool
	body            BlockDeclaration
	params          []ParamDeclaration
	returnType      Writable
	typeConstructor Writable
}

// Async marks the function as async
func (f *Function) Async() *Function {
	f.async = true
	return f
}

// Arrow use arrow syntax for function
func (f *Function) Arrow() *Function {
	f.arrowSyntax = true
	return f
}

// WriteCode implements Writable
func (f *Function) WriteCode(writer CodeWriter) {
	if f.async {
		writer.Write("async ")
	}

	if !f.arrowSyntax {
		writer.Write("function ")
	}

	f.writeCallSig(writer)

	sep := " "
	if f.arrowSyntax {
		sep = " => "
	}

	writer.Write(sep)
	f.writeBody(writer)
}

// writeCallSig writes functions call signature.
// This includes any type constructors, parameters, and return types
func (f *Function) writeCallSig(writer CodeWriter) {
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
}

// writeBody write function body
func (f *Function) writeBody(writer CodeWriter) {
	writer.OpenBlock()
	for _, ln := range f.body.lines {
		ln.WriteCode(writer)
		writer.Eol()
	}
	writer.CloseBlock()
}
