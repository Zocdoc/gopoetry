package java

import (
	"fmt"
	"github.com/zocdoc/gopoetry/util"
)

type ParamDeclaration struct {
	name  string
	type_ string
}

func (self *ParamDeclaration) WriteCode(writer util.CodeWriter) {
	writer.Write(fmt.Sprintf("%s %s", self.type_, self.name))
}

func Param(name string, type_ string) *ParamDeclaration {
	return &ParamDeclaration{
		name:  name,
		type_: type_,
	}
}
