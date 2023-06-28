package swift

import (
	"fmt"
	"testing"
)

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

func TestValAccessModifiers(t *testing.T) {
	myVal := Var("myVal", "Any")

	modifiers := []func() *VarDeclaration{
		myVal.Public,
		myVal.FilePrivate,
		myVal.Private,
		myVal.Internal,
	}

	expectedMod := []string{
		"public",
		"fileprivate",
		"private",
		"internal",
	}

	for i, modify := range modifiers {
		modify()
		expected := fmt.Sprintf("%s var myVal: Any", expectedMod[i])
		assertCode(t, myVal, expected)
	}
}

func TestValAttribute(t *testing.T) {
	myVal := Var("myVal", "Any").SimpleAttributes("@IBOutlet", "@Indirect")
	assertCode(t, myVal, "@IBOutlet @Indirect var myVal: Any")
}
