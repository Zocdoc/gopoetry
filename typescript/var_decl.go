package typescript

// declType represents the different types of variable declarations.
type declType string

const (
	varDecl   declType = "var"
	letDecl   declType = "let"
	constDecl declType = "const"
)

var (
	// Const declares a constant
	Const = declFact(constDecl)

	// Let declares a `let` variable
	Let = declFact(letDecl)

	// Var declares a `var` variable
	Var = declFact(varDecl)
)

// VarDeclaration represents a variable declaration
type VarDeclaration struct {
	name     string
	varValue Writable
	varType  Writable
	declType declType
	export   bool
}

// Export exports the variable
func (d *VarDeclaration) Export() *VarDeclaration {
	d.export = true
	return d
}

// WriteCode implements Writable
func (d *VarDeclaration) WriteCode(writer CodeWriter) {
	if d.export {
		writer.Write("export ")
	}

	writer.Write(string(d.declType))
	writer.Write(" " + d.name)

	if d.varType != nil {
		writer.Write(": ")
		d.varType.WriteCode(writer)
	}

	if d.varValue != nil {
		writer.Write(" = ")
		d.varValue.WriteCode(writer)
	}

	writer.Write(";")
	writer.Eol()
}

// declFact is a factory function for creating variables, parameterized by declaration type.
func declFact(declType declType) func(name string, varType Writable, varInitValue Writable) *VarDeclaration {
	return func(name string, varType Writable, varInitValue Writable) *VarDeclaration {
		return &VarDeclaration{
			declType: declType,
			name:     name,
			varType:  varType,
			varValue: varInitValue,
		}
	}
}
