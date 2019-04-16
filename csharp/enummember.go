package csharp

type EnumMemberDeclaration struct {
	name       string
	attributes []Writable
	value       *string
}

func (self *EnumMemberDeclaration) Value(value string) *EnumMemberDeclaration {
	self.value = &value
	return self
}

func (self *EnumMemberDeclaration) AddAttributes(attributes ...Writable) *EnumMemberDeclaration {
	self.attributes = append(self.attributes, attributes...)
	return self
}

func (self *EnumMemberDeclaration) WithAttribute(code string) *EnumMemberDeclaration {
	return self.AddAttributes(Attribute(code))
}

func EnumMember(name string) *EnumMemberDeclaration {
	return &EnumMemberDeclaration{
		name:       name,
		attributes: []Writable{},
		value:       nil,
	}
}

func (self *EnumMemberDeclaration) WriteCode(writer CodeWriter) {
	for _, attribute := range self.attributes {
		attribute.WriteCode(writer)
		writer.Eol()
	}
	writer.Write(self.name)
	if self.value != nil {
		writer.Write(" = ")
		writer.Write(*self.value)
	}
}