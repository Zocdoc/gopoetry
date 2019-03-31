package csharp

type UnitDeclaration struct {
	usings     []Writable
	namespaces []Writable
}

func (self *UnitDeclaration) AddUsings(usings ...Writable) *UnitDeclaration {
	self.usings = append(self.usings, usings...)
	return self
}

func (self *UnitDeclaration) Using(namespace string) *UnitDeclaration {
	self.AddUsings(Using(namespace))
	return self
}

func (self *UnitDeclaration) UsingStatic(namespace string) *UnitDeclaration {
	self.AddUsings(Using(namespace).Static())
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
		usings:     []Writable{},
		namespaces: []Writable{},
	}
}

func (self *UnitDeclaration) WriteCode(writer CodeWriter) {
	for _, import_ := range self.usings {
		import_.WriteCode(writer)
		writer.Eol()
	}
	for _, namespace := range self.namespaces {
		writer.Eol()
		namespace.WriteCode(writer)
	}
}
