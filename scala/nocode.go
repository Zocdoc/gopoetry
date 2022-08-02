package scala

type NoCodeDeclaration struct {
}

func (self *NoCodeDeclaration) WriteCode(writer CodeWriter) {
}

var NoCode = &NoCodeDeclaration{}
