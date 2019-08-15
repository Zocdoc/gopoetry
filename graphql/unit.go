package graphql

type UnitDeclaration struct {
	declarations []Writable
}

func Unit() *UnitDeclaration {
	return &UnitDeclaration{
		[]Writable{},
	}
}

func (self *UnitDeclaration) AddDeclarations(declarations ...Writable) *UnitDeclaration {
	self.declarations = append(self.declarations, declarations...)
	return self
}

func (self *UnitDeclaration) WriteCode(writer CodeWriter) {
	for index, declaration := range self.declarations {
		if index > 0 {
			writer.Eol()
		}
		declaration.WriteCode(writer)
	}
}

func (self *UnitDeclaration) Code() string {
	writer := CreateWriter()
	self.WriteCode(&writer)
	return writer.Code()
}
