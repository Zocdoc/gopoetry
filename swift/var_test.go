package swift

import "testing"

func TestVar(t *testing.T) {
	assertCode(t, Var("name", "String"), "var name: String")
}

func TestVarInUnit(t *testing.T) {
	unit := NewUnit().AddImports(
		Import("Foundation"),
		Import("UIKit"),
	).AddDeclarations(
		Var("name", "String"),
		Var("count", "Int"),
	)

	expected := `
import Foundation
import UIKit

var name: String

var count: Int
`
	assertCode(t, unit, expected)
}
