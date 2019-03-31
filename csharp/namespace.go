package csharp

type NamespaceDeclaration struct {
	namespace string
	usings    []Writable
	classes   []Writable
}

func (self *NamespaceDeclaration) AddUsings(usings ...Writable) *NamespaceDeclaration {
	self.usings = append(self.usings, usings...)
	return self
}

func (self *NamespaceDeclaration) Using(namespace string) *NamespaceDeclaration {
	self.AddUsings(Using(namespace))
	return self
}

func (self *NamespaceDeclaration) UsingStatic(namespace string) *NamespaceDeclaration {
	self.AddUsings(Using(namespace).Static())
	return self
}

func (self *NamespaceDeclaration) AddClasses(classes ...Writable) *NamespaceDeclaration {
	self.classes = append(self.classes, classes...)
	return self
}

func Namespace(namespace string) *NamespaceDeclaration {
	return &NamespaceDeclaration{
		namespace: namespace,
		usings:    []Writable{},
		classes:   []Writable{},
	}
}

func (self *NamespaceDeclaration) WriteCode(writer CodeWriter) {
	writer.Write("namespace " + self.namespace)
	writer.Begin()
	for _, import_ := range self.usings {
		import_.WriteCode(writer)
		writer.Eol()
	}
	for _, class := range self.classes {
		writer.Eol()
		class.WriteCode(writer)
	}
	writer.End()
}
