package csharp

type ImportDeclaration struct {
	namespace string
	isStatic  bool
}

func Import(namespace string) *ImportDeclaration {
	return &ImportDeclaration{
		namespace: namespace,
		isStatic:  false,
	}
}

func (self *ImportDeclaration) Static() *ImportDeclaration {
	self.isStatic = true
	return self
}

func (self *ImportDeclaration) WriteCode(writer CodeWriter) {
	line := "import "
	if self.isStatic {
		line = line + "static "
	}
	line = line + self.namespace
	line = line + ";"
	writer.Write(line)
}
