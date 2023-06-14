package swift

import "fmt"

type VarDeclaration struct {
	accessModifier string
	name           string
	typeName       string
	initValue      Writable
}

// VarDeclaration implements Declaration.
var _ Declaration = (*VarDeclaration)(nil)

func (v *VarDeclaration) Declaration() {}

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

func (v *VarDeclaration) Public() *VarDeclaration {
	v.accessModifier = "public"
	return v
}

func (v *VarDeclaration) Internal() *VarDeclaration {
	v.accessModifier = "internal"
	return v
}

func (v *VarDeclaration) FilePrivate() *VarDeclaration {
	v.accessModifier = "fileprivate"
	return v
}

func (v *VarDeclaration) Private() *VarDeclaration {
	v.accessModifier = "private"
	return v
}

func (v *VarDeclaration) WriteCode(writer CodeWriter) {
	if v.accessModifier != "" {
		writer.Write(v.accessModifier + " ")
	}

	writer.Write(fmt.Sprintf("var %s: %s", v.name, v.typeName))
	if v.initValue != nil {
		writer.Write(" = ")
		v.initValue.WriteCode(writer)
	}
}
