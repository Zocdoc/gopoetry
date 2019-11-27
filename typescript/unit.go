package typescript

type UnitDeclaration struct {
	imports      []Writable
	namespaces   []Writable
	declarations []Writable
}

func (unit *UnitDeclaration) AddImports(imports ...Writable) *UnitDeclaration {
	unit.imports = append(unit.imports, imports...)
	return unit
}

func (unit *UnitDeclaration) NamedImport(module string, params ...string) *UnitDeclaration {
	unit.AddImports(NamedImport(module, params...))
	return unit
}

func (unit *UnitDeclaration) DefaultImport(module string, as string) *UnitDeclaration {
	unit.AddImports(DefaultImport(module, as))
	return unit
}

func (unit *UnitDeclaration) AddNamespaces(namespaces ...Writable) *UnitDeclaration {
	unit.namespaces = append(unit.namespaces, namespaces...)
	return unit
}

func (unit *UnitDeclaration) Namespace(namespace string) *NamespaceDeclaration {
	namespace_ := Namespace(namespace)
	unit.AddNamespaces(namespace_)
	return namespace_
}

// AddDeclarations adds declarations to the unit
func (unit *UnitDeclaration) AddDeclarations(declarations ...Writable) *UnitDeclaration {
	unit.declarations = append(unit.declarations, declarations...)
	return unit
}

func (unit *UnitDeclaration) Code() string {
	writer := CreateWriter()
	unit.WriteCode(&writer)
	return writer.Code()
}

func Unit() *UnitDeclaration {
	return &UnitDeclaration{
		imports:    []Writable{},
		namespaces: []Writable{},
	}
}

func (unit *UnitDeclaration) WriteCode(writer CodeWriter) {
	for _, import_ := range unit.imports {
		import_.WriteCode(writer)
		writer.Eol()
	}
	for _, namespace := range unit.namespaces {
		writer.Eol()
		namespace.WriteCode(writer)
	}

	for _, class := range unit.declarations {
		writer.Eol()
		class.WriteCode(writer)
	}
}
