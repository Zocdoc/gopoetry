package swift

import "fmt"

type StructDecl struct {
	accessModifier string
	name           string
	members        []Declaration
}

var _ Declaration = (*StructDecl)(nil)

func (*StructDecl) Declaration() {}

func NewStruct(name string) *StructDecl {
	return &StructDecl{
		name: name,
	}
}

func (s *StructDecl) WriteCode(writer CodeWriter) {
	if s.accessModifier != "" {
		writer.Write(s.accessModifier + " ")
	}

	writer.Write(fmt.Sprintf("struct %s ", s.name))

	if len(s.members) == 0 {
		writer.Write("{}")
		return
	}

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

func (s *StructDecl) Public() *StructDecl {
	s.accessModifier = "public"
	return s
}

func (s *StructDecl) Internal() *StructDecl {
	s.accessModifier = "internal"
	return s
}

func (s *StructDecl) FilePrivate() *StructDecl {
	s.accessModifier = "fileprivate"
	return s
}

func (s *StructDecl) Private() *StructDecl {
	s.accessModifier = "private"
	return s
}
