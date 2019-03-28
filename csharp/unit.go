package csharp

type UnitDeclaration struct {
	imports    []Writable
	namespaces []Writable
}

func (self *UnitDeclaration) Imports(imports ...Writable) *UnitDeclaration {
	self.imports = append(self.imports, imports...)
	return self
}

func (self *UnitDeclaration) Import(namespace string) *UnitDeclaration {
	self.Imports(Import(namespace))
	return self
}

func (self *UnitDeclaration) ImportStatic(namespace string) *UnitDeclaration {
	self.Imports(Import(namespace).Static())
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
