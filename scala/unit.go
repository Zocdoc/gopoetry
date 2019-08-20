package scala

import "fmt"

type UnitDeclaration struct {
	package_     string
	imports      []string
	declarations []Writable
}

func Unit(package_ string) *UnitDeclaration {
	return &UnitDeclaration{
		package_,
		[]string{},
		[]Writable{},
	}
}

func (self *UnitDeclaration) AddImports(imports ...string) *UnitDeclaration {
	self.imports = append(self.imports, imports...)
	return self
}

func (self *UnitDeclaration) Import(package_ string) *UnitDeclaration {
	self.AddImports(package_)
	return self
}

func (self *UnitDeclaration) AddDeclarations(declarations ...Writable) *UnitDeclaration {
	self.declarations = append(self.declarations, declarations...)
	return self
}

func (self *UnitDeclaration) Code() string {
	writer := CreateWriter()
	self.WriteCode(&writer)
	return writer.Code()
}

func (self *UnitDeclaration) WriteCode(writer CodeWriter) {
	writer.Write(fmt.Sprintf("package %s", self.package_))
	writer.Eol()
	if len(self.imports) > 0 {
		writer.Eol()
		for _, import_ := range self.imports {
			writer.Write(fmt.Sprintf("import %s", import_))
			writer.Eol()
		}
	}
	for _, class := range self.declarations {
		writer.Eol()
		class.WriteCode(writer)
		writer.Eol()
	}
}
