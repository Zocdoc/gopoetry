package scala

type EolDefinition struct {}

func Eol() *EolDefinition {
	return &EolDefinition{}
}

func (self *EolDefinition) WriteCode(writer CodeWriter) {
	writer.Eol()
}

