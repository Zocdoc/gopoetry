package csharp

type NamespaceDeclaration struct {
	namespace string
	imports   []Writable
	classes   []Writable
}

func (self *NamespaceDeclaration) Imports(imports ...Writable) *NamespaceDeclaration {
	self.imports = append(self.imports, imports...)
	return self
}

func (self *NamespaceDeclaration) Import(namespace string) *NamespaceDeclaration {
	self.Imports(Import(namespace))
	return self
}

func (self *NamespaceDeclaration) ImportStatic(namespace string) *NamespaceDeclaration {
	self.Imports(Import(namespace).Static())
	return self
}

func (self *NamespaceDeclaration) AddClasses(classes ...Writable) *NamespaceDeclaration {
	self.classes = append(self.classes, classes...)
	return self
}

func Namespace(namespace string) *NamespaceDeclaration {
	return &NamespaceDeclaration{
		namespace: namespace,
		imports:   []Writable{},
		classes:   []Writable{},
	}
}

func (self *NamespaceDeclaration) WriteCode(writer CodeWriter) {
	writer.Write("namespace " + self.namespace)
	writer.Begin()
	for _, import_ := range self.imports {
		import_.WriteCode(writer)
		writer.Eol()
	}
	for _, class := range self.classes {
		writer.Eol()
		class.WriteCode(writer)
	}
	writer.End()
}
