package scala

type UnitDeclaration struct {
}

func Unit(packageName string) *UnitDeclaration {
	return &UnitDeclaration{}
}

func (self *UnitDeclaration) WriteCode(writer CodeWriter) {
}
