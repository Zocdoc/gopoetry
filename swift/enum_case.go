package swift

import "fmt"

type EnumCaseDecl struct {
	name    string
}

func (s EnumCaseDecl) WriteCode(writer CodeWriter) {
	writer.Write(fmt.Sprintf("case %s", s.name))
}

func NewCase(name string) EnumCaseDecl {
	return EnumCaseDecl{
		name: name,
	}
}
