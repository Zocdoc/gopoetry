package swift

// Unit represents a Swift source file
type Unit struct {
	imports      []ImportDeclaration
	declarations []Declaration
}

func NewUnit() *Unit {
	return &Unit{}
}

func (u *Unit) AddImports(imports ...ImportDeclaration) *Unit {
	u.imports = append(u.imports, imports...)
	return u
}

func (u *Unit) AddDeclarations(decls ...Declaration) *Unit {
	u.declarations = append(u.declarations, decls...)
	return u
}

// WriteCode implements Writable.
func (u *Unit) WriteCode(writer CodeWriter) {
	for _, importDecl := range u.imports {
		importDecl.WriteCode(writer)
		writer.Eol()
	}

	writer.Eol()

	for _, decl := range u.declarations {
		decl.WriteCode(writer)
		writer.Eol()
	}
}

var _ Writable = (*Unit)(nil)
