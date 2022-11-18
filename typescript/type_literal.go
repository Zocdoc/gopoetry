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
	name           string
	optional       bool
	typeAnnotation Writable
}

func (ps *PropertySig) WriteCode(writer CodeWriter) {
	name := ps.name
	if ps.optional {
		name += "?"
	}

	writer.Write(name + ": ")
	ps.typeAnnotation.WriteCode(writer)
	writer.Write(";")
	writer.Eol()
}
