package typescript

// ObjectValue represents a javascript object
type ObjectValue struct {
	BlockDeclaration
}

// WriteCode implements Writable
func (t *ObjectValue) WriteCode(writer CodeWriter) {
	writer.OpenBlock()
	for _, member := range t.lines {
		member.WriteCode(writer)
	}
	writer.CloseBlock()
}

// AddProp add a property
func (o *ObjectValue) AddProp(name string, value Writable) *ObjectValue {
	o.AppendCode(&ObjProp{
		Name:  name,
		Value: value,
	})
	return o
}

// ObjProp represent a named property with a value on an object
type ObjProp struct {
	Name  string
	Value Writable
}

func (ps *ObjProp) WriteCode(writer CodeWriter) {
	writer.Write(ps.Name + ": ")
	ps.Value.WriteCode(writer)
	writer.Write(",")
	writer.Eol()
}
