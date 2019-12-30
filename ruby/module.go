package ruby

type ModuleDeclaration struct {
	name           string
	declarations   []Writable
}

func (self *ModuleDeclaration) AddDeclarations(declarations ...Writable) *ModuleDeclaration {
	self.declarations = append(self.declarations, declarations...)
	return self
}

func Module(name string) *ModuleDeclaration {
	return &ModuleDeclaration{
		name:         name,
		declarations: []Writable{},
	}
}

func (self *ModuleDeclaration) WriteCode(writer CodeWriter) {
	writer.Begin("module "+self.name)
	for index, declaration := range self.declarations {
		if index > 0 { writer.Eol() }
		declaration.WriteCode(writer)
	}
	writer.End()
}

