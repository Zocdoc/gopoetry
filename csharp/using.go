package csharp

type UsingDeclaration struct {
	namespace string
	isStatic  bool
}

func Using(namespace string) *UsingDeclaration {
	return &UsingDeclaration{
		namespace: namespace,
		isStatic:  false,
	}
}

func (self *UsingDeclaration) Static() *UsingDeclaration {
	self.isStatic = true
	return self
}

func (self *UsingDeclaration) WriteCode(writer CodeWriter) {
	line := "using "
	if self.isStatic {
		line = line + "static "
	}
	line = line + self.namespace
	line = line + ";"
	writer.Write(line)
}
