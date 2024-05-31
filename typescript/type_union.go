package typescript

import (
	"fmt"
)

// DeclareType instantiates a new type declaration
func DeclareType(name string, typeRef Writable) *TypeDeclaration {
	return &TypeDeclaration{
		name:          name,
		typeReference: typeRef,
	}
}

// TypeDeclaration declares a type
type TypeDeclaration struct {
	name          string
	typeReference Writable
	export        bool
	CommentBlockDeclaration
}

func (td *TypeDeclaration) Export() *TypeDeclaration {
	td.export = true
	return td
}

// WriteCode writes the type declaration to the writer
func (td *TypeDeclaration) WriteCode(writer CodeWriter) {
	td.WriteComments(writer)
	exportMod := ""
	if td.export {
		exportMod = "export " // trailing space is intentional
	}

	writer.Write(fmt.Sprintf("%stype %s = ", exportMod, td.name))
	td.typeReference.WriteCode(writer)
	writer.Write(";")
	writer.Eol()
}

// UnionTypeRef represents the union of a series of type refs
// see: https://www.typescriptlang.org/docs/handbook/typescript-in-5-minutes.html#unions
type UnionTypeRef struct {
	typeRefs []Writable
}

// UnionType is a constructor function for creating union types
func UnionType(refs ...string) *UnionTypeRef {
	typeRefs := []Writable{}

	for _, str := range refs {
		typeRefs = append(typeRefs, Code(str))
	}

	return &UnionTypeRef{
		typeRefs: typeRefs,
	}
}

func (utr *UnionTypeRef) Union(ref Writable) *UnionTypeRef {
	return &UnionTypeRef{
		typeRefs: append(utr.typeRefs, ref),
	}
}

// WriteCode writes the union type ref to the writer
func (utr *UnionTypeRef) WriteCode(writer CodeWriter) {
	for i, ref := range utr.typeRefs {
		ref.WriteCode(writer)
		if i < len(utr.typeRefs)-1 {
			writer.Write(" | ")
		}
	}
}
