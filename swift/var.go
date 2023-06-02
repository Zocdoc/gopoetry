package swift

import "fmt"

type VarDeclaration struct {
	name, typeName string
}

func Var(name, typeName string) *VarDeclaration {
	return &VarDeclaration{
		name:     name,
		typeName: typeName,
	}
}

// VarDeclaration implements Declaration.
var _ Declaration = (*VarDeclaration)(nil)

func (v *VarDeclaration) Declaration() {}

func (v *VarDeclaration) WriteCode(writer CodeWriter) {
	writer.Write(fmt.Sprintf("var %s: %s", v.name, v.typeName))
}
