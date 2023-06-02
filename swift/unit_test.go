package swift

import "testing"

func TestUnitEmpty(t *testing.T) {
	assertCode(t, NewUnit(), "")
}

func TestUnitUsingWithNamespace(t *testing.T) {
	unit := NewUnit().AddImports(
		Import("Foundation"),
		Import("UIKit"),
	)

	expected := `
import Foundation
import UIKit
`
	assertCode(t, unit, expected)
}
