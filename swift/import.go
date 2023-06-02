package swift

import "fmt"

type ImportDeclaration struct {
	moduleName string
}

var _ Writable = (*ImportDeclaration)(nil)

func Import(moduleName string) ImportDeclaration {
	return ImportDeclaration{
		moduleName: moduleName,
	}
}

func (id *ImportDeclaration) WriteCode(writer CodeWriter) {
	writer.Write(fmt.Sprintf("import %s", id.moduleName))
}
