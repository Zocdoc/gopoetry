package typescript

type UnitDeclaration struct {
	imports    []Writable
	namespaces []Writable
}

func (self *UnitDeclaration) AddImports(imports ...Writable) *UnitDeclaration {
	self.imports = append(self.imports, imports...)
	return self
}

func (self *UnitDeclaration) NamedImport(module string, params ...string) *UnitDeclaration {
	self.AddImports(NamedImport(module, params...))
	return self
}

func (self *UnitDeclaration) DefaultImport(module string, as string) *UnitDeclaration {
	self.AddImports(DefaultImport(module, as))
	return self
}

func (self *UnitDeclaration) AddNamespaces(namespaces ...Writable) *UnitDeclaration {
	self.namespaces = append(self.namespaces, namespaces...)
	return self
}

func (self *UnitDeclaration) Namespace(namespace string) *NamespaceDeclaration {
	namespace_ := Namespace(namespace)
	self.AddNamespaces(namespace_)
	return namespace_
}

func (self *UnitDeclaration) Code() string {
	writer := CreateWriter()
	self.WriteCode(&writer)
	return writer.Code()
}

func Unit() *UnitDeclaration {
	return &UnitDeclaration{
		imports:    []Writable{},
		namespaces: []Writable{},
	}
}

func (self *UnitDeclaration) WriteCode(writer CodeWriter) {
	for _, import_ := range self.imports {
		import_.WriteCode(writer)
		writer.Eol()
	}
	for _, namespace := range self.namespaces {
		writer.Eol()
		namespace.WriteCode(writer)
	}
}
