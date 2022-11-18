package typescript

var _ Writable = &ObjectType{}

type ObjectType struct {
	typeMembers []Writable
}

func NewObjectType(name string, typeAnnotation Writable) *ObjectType {
	ot := ObjectType{}
	ot.AddMember(name, typeAnnotation)
	return &ot
}

func (o *ObjectType) AddMember(name string, typeAnnotation Writable) *ObjectType {
	o.typeMembers = append(o.typeMembers, &PropertySig{
		name:           name,
		typeAnnotation: typeAnnotation,
	})

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
	typeAnnotation Writable
}

func (ps *PropertySig) WriteCode(writer CodeWriter) {
	writer.Write(ps.name + ": ")
	ps.typeAnnotation.WriteCode(writer)
	writer.Write(";")
	writer.Eol()
}
