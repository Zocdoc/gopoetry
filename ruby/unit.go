package ruby

type UnitDeclaration struct {
	requirements []Writable
	declarations []Writable
}

func (self *UnitDeclaration) AddRequire(require ...Writable) *UnitDeclaration {
	self.requirements = append(self.requirements, require...)
	return self
}

func (self *UnitDeclaration) Require(filename string) *UnitDeclaration {
	self.AddRequire(Require(filename))
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

func Unit() *UnitDeclaration {
	return &UnitDeclaration{
		requirements: []Writable{},
		declarations: []Writable{},
	}
}

func (self *UnitDeclaration) WriteCode(writer CodeWriter) {
	for _, require := range self.requirements {
		require.WriteCode(writer)
	}
	for _, declaration := range self.declarations {
		declaration.WriteCode(writer)
	}
}
