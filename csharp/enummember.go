package csharp

type EnumMemberDeclaration struct {
	name       string
	attributes []Writable
	value      *string
	summary    SummaryDeclaration
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
		value:      nil,
		summary:    SummaryDeclaration{},
	}
}

func (self *EnumMemberDeclaration) WriteCode(writer CodeWriter) {
	self.summary.WriteCode(writer)
	if len(self.attributes) > 0 {
		writer.Write("[")
		for i, attribute := range self.attributes {
			if i > 0 {
				writer.Write(", ")
			}
			attribute.WriteCode(writer)
		}
		writer.Write("]")
		writer.Eol()
	}
	writer.Write(self.name)
	if self.value != nil {
		writer.Write(" = ")
		writer.Write(*self.value)
	}
}

func (self *EnumMemberDeclaration) Summary(summary string) {
	self.summary.AddDescription(summary)
}
