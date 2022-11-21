package typescript

var _ Writable = &ObjectType{}

type ObjectType struct {
	typeMembers []Writable
}

func NewObjectType() *ObjectType {
	return &ObjectType{}
}

func (o *ObjectType) AddProp(prop *PropertySig) *ObjectType {
	o.typeMembers = append(o.typeMembers, prop)
	return o
}

// WriteCode implements Writable
func (t *ObjectType) WriteCode(writer CodeWriter) {
	writer.OpenBlock()
	for _, member := range t.typeMembers {
		member.WriteCode(writer)
	}
	writer.CloseBlock()
}

var _ Writable = &PropertySig{}

type PropertySig struct {
	Name           string
	Optional       bool
	TypeAnnotation Writable
}

func (ps *PropertySig) WriteCode(writer CodeWriter) {
	name := ps.Name
	if ps.Optional {
		name += "?"
	}

	writer.Write(name + ": ")
	ps.TypeAnnotation.WriteCode(writer)
	writer.Write(";")
	writer.Eol()
}
