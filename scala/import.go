package scala

import "fmt"

type ImportDeclaration struct {
	package_ string
}

func Import(package_ string) *ImportDeclaration {
	return &ImportDeclaration{package_}
}

func (self *ImportDeclaration) WriteCode(writer CodeWriter) {
	writer.Write(fmt.Sprintf("import %s", self.package_))
	writer.Eol()
}
