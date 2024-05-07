package typescript

import "fmt"

// ObjectValue represents a javascript object
type ObjectValue struct {
	BlockDeclaration
	preAssignmentSpreads  []string
	postAssignmentSpreads []string
}

func NewObjectValue() *ObjectValue {
	return &ObjectValue{}
}

// WriteCode implements Writable
func (t *ObjectValue) WriteCode(writer CodeWriter) {
	if len(t.lines) == 0 {
		writer.Write("{}")
		return
	}

	writer.OpenBlock()

	for _, spread := range t.preAssignmentSpreads {
		writer.Write(fmt.Sprintf("...%s,", spread))
		writer.Eol()
	}

	for _, member := range t.lines {
		member.WriteCode(writer)
	}

	for _, spread := range t.postAssignmentSpreads {
		writer.Write(fmt.Sprintf("...%s,", spread))
		writer.Eol()
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

func (o *ObjectValue) AddSpread(objectNameToSpread string, postPropertyAssignment bool) *ObjectValue {
	if postPropertyAssignment {
		o.postAssignmentSpreads = append(o.postAssignmentSpreads, objectNameToSpread)
	} else {
		o.preAssignmentSpreads = append(o.preAssignmentSpreads, objectNameToSpread)
	}
	return o
}

// ObjProp represent a named property with a value on an object
type ObjProp struct {
	Name  string
	Value Writable
}

func (ps *ObjProp) WriteCode(writer CodeWriter) {
	writer.Write(ps.Name)

	if ps.Value != nil {
		writer.Write(": ")
		ps.Value.WriteCode(writer)
	}
	writer.Write(",")
	writer.Eol()
}
