package poetrycs

import (
	"fmt"
)

type ParamDeclaration struct {
	name string
	type_ string
}

func (self *ParamDeclaration) WriteCode(writer CodeWriter) {
	writer.Write(fmt.Sprintf("%s %s", self.type_, self.name))
}

func Param(type_ string, name string) ParamDeclaration {
	return ParamDeclaration{
		name: name,
		type_: type_,
	}
}