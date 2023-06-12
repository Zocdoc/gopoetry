package swift

import "fmt"

type StructDecl struct {
	name    string
	members []Declaration
}

var _ Declaration = (*StructDecl)(nil)

func (*StructDecl) Declaration() {}

func (s *StructDecl) WriteCode(writer CodeWriter) {
	writer.Write(fmt.Sprintf("struct %s ", s.name))
	writer.OpenBlock()

	for _, member := range s.members {
		member.WriteCode(writer)
		writer.Eol()
	}

	writer.CloseBlock()
}

func (s *StructDecl) AddMembers(members ...Declaration) *StructDecl {
	s.members = append(s.members, members...)
	return s
}

func NewStruct(name string) *StructDecl {
	return &StructDecl{
		name: name,
	}
}
