package swift

import "fmt"

type EnumDecl struct {
	name    string
	cases []EnumCaseDecl
}

var _ Declaration = (*EnumDecl)(nil)

func (*EnumDecl) Declaration() {}

func (s *EnumDecl) WriteCode(writer CodeWriter) {
	writer.Write(fmt.Sprintf("enum %s ", s.name))
	writer.OpenBlock()

	for _, c := range s.cases {
		c.WriteCode(writer)
		writer.Eol()
	}

	writer.CloseBlock()
}

func (s *EnumDecl) AddCases(cases ...EnumCaseDecl) *EnumDecl {
	s.cases = append(s.cases, cases...)
	return s
}

func NewEnum(name string) *EnumDecl {
	return &EnumDecl{
		name: name,
	}
}
