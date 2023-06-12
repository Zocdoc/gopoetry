package swift

import "fmt"

type VarDeclaration struct {
	name, typeName string
	initValue      Writable
}

func Var(name, typeName string) *VarDeclaration {
	return &VarDeclaration{
		name:     name,
		typeName: typeName,
	}
}

func (v *VarDeclaration) InitWith(value Writable) *VarDeclaration {
	v.initValue = value
	return v
}

// VarDeclaration implements Declaration.
var _ Declaration = (*VarDeclaration)(nil)

func (v *VarDeclaration) Declaration() {}

func (v *VarDeclaration) WriteCode(writer CodeWriter) {
	writer.Write(fmt.Sprintf("var %s: %s", v.name, v.typeName))
	if v.initValue != nil {
		writer.Write(" = ")
		v.initValue.WriteCode(writer)
	}
}
