package csharp

import (
	"fmt"
	"strings"
)

type RecordDeclaration struct {
	name       string
	inherits   []string
	modifiers  []string
	attributes []Writable
	members    []Writable
	summary    SummaryDeclaration
}

func (record *RecordDeclaration) addModifier(modifier string) *RecordDeclaration {
	record.modifiers = append(record.modifiers, modifier)
	return record
}

func (record *RecordDeclaration) Private() *RecordDeclaration {
	return record.addModifier("private")
}

func (record *RecordDeclaration) Public() *RecordDeclaration {
	return record.addModifier("public")
}

func (record *RecordDeclaration) Abstract() *RecordDeclaration {
	return record.addModifier("abstract")
}

func (record *RecordDeclaration) Sealed() *RecordDeclaration {
	return record.addModifier("sealed")
}

func (record *RecordDeclaration) Inherits(types ...string) *RecordDeclaration {
	record.inherits = append(record.inherits, types...)
	return record
}

func (record *RecordDeclaration) AddMembers(members ...Writable) *RecordDeclaration {
	record.members = append(record.members, members...)
	return record
}

func (record *RecordDeclaration) AddAttributes(attributes ...Writable) *RecordDeclaration {
	record.attributes = append(record.attributes, attributes...)
	return record
}

func (record *RecordDeclaration) WithAttribute(code string) *RecordDeclaration {
	return record.AddAttributes(Attribute(code))
}

func (record *RecordDeclaration) Method(name string) *MethodDeclaration {
	method := Method(name)
	record.AddMembers(method)
	return method
}

func (record *RecordDeclaration) Constructor() *MethodDeclaration {
	ctor := Constructor(record.name)
	record.AddMembers(ctor)
	return ctor
}

func (record *RecordDeclaration) Field(type_ string, name string) *FieldDeclaration {
	field := Field(type_, name)
	record.AddMembers(field)
	return field
}

func (record *RecordDeclaration) Property(type_ string, name string) *PropertyDeclaration {
	property := Property(type_, name)
	record.AddMembers(property)
	return property
}

func (record *RecordDeclaration) Summary(summary string) *RecordDeclaration {
	record.summary.AddDescription(summary)
	return record
}

func Record(name string) *RecordDeclaration {
	return &RecordDeclaration{
		name:       name,
		modifiers:  []string{},
		attributes: []Writable{},
		members:    []Writable{},
		summary:    SummaryDeclaration{},
	}
}

func (record *RecordDeclaration) WriteCode(writer CodeWriter) {
	record.summary.WriteCode(writer)

	declaration := fmt.Sprintf("record %s", record.name)
	if len(record.modifiers) > 0 {
		declaration = strings.Join(record.modifiers, " ") + " " + declaration
	}

	if len(record.inherits) > 0 {
		declaration += " : " + strings.Join(record.inherits, ", ")
	}

	if len(record.attributes) > 0 {
		writer.Write("[")
		for i, attribute := range record.attributes {
			if i > 0 {
				writer.Write(", ")
			}
			attribute.WriteCode(writer)
		}
		writer.Write("]")
		writer.Eol()
	}

	writer.Write(declaration)
	writer.Begin()
	for index, member := range record.members {
		if index > 0 {
			writer.Eol()
		}
		member.WriteCode(writer)
	}
	writer.End()
}
