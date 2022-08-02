package ruby

import (
	"fmt"
)

type ArgDeclaration struct {
	name      string
	isKeyword bool
	default_  Writable
}

func (self *ArgDeclaration) Keyword() *ArgDeclaration {
	self.isKeyword = true
	return self
}

func (self *ArgDeclaration) Default(default_ Writable) *ArgDeclaration {
	self.default_ = default_
	return self
}

func (self *ArgDeclaration) WriteCode(writer CodeWriter) {
	if self.isKeyword {
		writer.Write(fmt.Sprintf("%s:", self.name))
	} else {
		writer.Write(self.name)
	}

	if self.default_ != nil {
		if !self.isKeyword {
			writer.Write(" = ")
		} else {
			writer.Write(" ")
		}
		self.default_.WriteCode(writer)
	}
}

func NewArg(name string) *ArgDeclaration {
	return &ArgDeclaration{
		name:      name,
		isKeyword: false,
		default_:  nil,
	}
}

func KeywordArg(name string) *ArgDeclaration {
	return NewArg(name).Keyword()
}

func Arg(name string) *ArgDeclaration {
	return NewArg(name)
}
