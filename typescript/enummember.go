package typescript

// EnumMemberDeclaration declares an enum member
type EnumMemberDeclaration struct {
	name  string
	value *WritableCode
}

// Value sets the enums value
func (enum *EnumMemberDeclaration) Value(value *WritableCode) *EnumMemberDeclaration {
	enum.value = value
	return enum
}

// EnumMember creates a new enum member
func EnumMember(name string) *EnumMemberDeclaration {
	return &EnumMemberDeclaration{
		name:  name,
		value: nil,
	}
}

// WriteCode writes the enum to the writer
func (enum *EnumMemberDeclaration) WriteCode(writer CodeWriter) {
	writer.Write(enum.name)
	if enum.value != nil {
		writer.Write(" = ")
		writer.Write(enum.value.code)
	}
}
