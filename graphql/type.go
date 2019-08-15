package graphql

type TypeDeclaration struct {
	name    string
	members []Writable
}

func Type(name string) *TypeDeclaration {
	return &TypeDeclaration{
		name:    name,
		members: []Writable{},
	}
}

func (self *TypeDeclaration) Field(name string, type_ string) *FieldDeclaration {
	field := Field(name, type_)
	self.members = append(self.members, field)
	return field
}

func (self *TypeDeclaration) WriteCode(writer CodeWriter) {
	if len(self.members) > 0 {
		writer.Write(" ")
		writer.Begin()
		for index, member := range self.members {
			if index > 0 {
				writer.Eol()
			}
			member.WriteCode(writer)
		}
		writer.End()
	}
}
