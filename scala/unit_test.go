package scala

import "testing"

func TestImport(t *testing.T) {
	expected := `
package com.example

import com.example.Example`
	assertCode(t, Unit("com.example").Import("com.example.Example"), expected)
}

func TestTwoDeclarations(t *testing.T) {
	expected := `
package com.example

import com.example.Example

trait IExample

object IExample`
	unit := Unit("com.example")
	unit.Import("com.example.Example")
	unit.AddDeclarations(Trait("IExample"), Object("IExample"))
	assertCode(t, unit, expected)
}
