package typescript

// ObjectType represents the object literal type
type ObjectType struct {
	BlockDeclaration
}

var _ Writable = &ObjectType{}

func NewObjectType() *ObjectType {
	return &ObjectType{}
}

// AddProp adds a property to the object type
func (o *ObjectType) AddProp(prop *PropertySig) *ObjectType {
	o.AppendCode(prop)
	return o
}

// WriteCode implements Writable
func (t *ObjectType) WriteCode(writer CodeWriter) {
	writer.OpenBlock()
	for _, member := range t.lines {
		member.WriteCode(writer)
	}
	writer.CloseBlock()
}

// PropertySig represents named type properties of object literal types
type PropertySig struct {
	Name           string
	Optional       bool
	TypeAnnotation Writable
}

var _ Writable = &PropertySig{}

// WriteCode implements Writable
func (ps *PropertySig) WriteCode(writer CodeWriter) {
	ps.writeIdAndTypeAnnotation(writer)
	writer.Write(";")
	writer.Eol()
}

func (ps *PropertySig) writeIdAndTypeAnnotation(writer CodeWriter) {
	name := ps.Name
	if ps.Optional {
		name += "?"
	}

	writer.Write(name + ": ")
	ps.TypeAnnotation.WriteCode(writer)
}
