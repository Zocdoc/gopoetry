package typescript

// FunctionDecl represents a top level declaration.
type FunctionDecl struct {
	Name     string
	exported bool
	FunctionExpression
}

// DeclareFunction returns a new function with the specified name
func DeclareFunction(name string) *FunctionDecl {
	return &FunctionDecl{Name: name}
}

// GetExpression returns the associated function expression
func (fd *FunctionDecl) GetExpression() *FunctionExpression {
	return &fd.FunctionExpression
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
	f.body.WriteCode(writer)
	writer.Eol()
}

// FunctionExpression represents the function as an expression
type FunctionExpression struct {
	async           bool
	arrow           bool
	body            OTBSBlock
	params          []ParamDeclaration
	returnType      Writable
	typeConstructor Writable
}

// NewFunctionExpression crates a new function expression struct
func NewFunctionExpression() *FunctionExpression {
	return &FunctionExpression{}
}

// AddParams add parameters to function expression call signature
func (f *FunctionExpression) AddParams(params ...ParamDeclaration) *FunctionExpression {
	f.params = append(f.params, params...)
	return f
}

// Async configures function expression to be async
func (f *FunctionExpression) Async() *FunctionExpression {
	f.async = true
	return f
}

// SetAsync configures function expression use arrow syntax
func (f *FunctionExpression) Arrow() *FunctionExpression {
	f.arrow = true
	return f
}

// AppendToBody appends lines of code to the body of the function expression
func (f *FunctionExpression) AppendToBody(lines ...Writable) *FunctionExpression {
	f.body.AppendCode(lines...)
	return f
}

// TypeConstructor sets the type constructor for the function expression
func (f *FunctionExpression) TypeConstructor(typeConstructor Writable) *FunctionExpression {
	f.typeConstructor = typeConstructor
	return f
}

// ReturnType sets the type constructor for the function expression
func (f *FunctionExpression) ReturnType(returnType Writable) *FunctionExpression {
	f.returnType = returnType
	return f
}

// WriteCode implements Writable
func (f *FunctionExpression) WriteCode(writer CodeWriter) {
	if f.async {
		writer.Write("async ")
	}

	if !f.arrow {
		writer.Write("function ")
	}

	f.writeCallSig(writer)

	sep := " "
	if f.arrow {
		sep = " => "
	}

	writer.Write(sep)
	f.body.WriteCode(writer)
}

// writeCallSig writes functions call signature.
// This includes any type constructors, parameters, and return types
func (f *FunctionExpression) writeCallSig(writer CodeWriter) {
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
