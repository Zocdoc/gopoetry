package graphql

import "fmt"

type FieldDeclaration struct {
	name  string
	type_ string
}

func Field(name string, type_ string) *FieldDeclaration {
	return &FieldDeclaration{
		name:  name,
		type_: type_,
	}
}

func (self *FieldDeclaration) WriteCode(writer CodeWriter) {
	writer.Write(fmt.Sprintf("%s: %s", self.name, self.type_))
}
