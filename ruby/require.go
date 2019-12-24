package ruby

import "fmt"

// TODO: Require new lines

type RequireDeclaration struct {
	filename string
}

func Require(filename string) *RequireDeclaration {
	return &RequireDeclaration{
		filename: filename,
	}
}

func (self *RequireDeclaration) WriteCode(writer CodeWriter) {
	writer.Write(fmt.Sprintf("require \"%s\"", self.filename))
	writer.Eol()
}

