package scala

import "testing"

func TestImport(t *testing.T) {
	expected := `
package com.example

import com.example.Example
`
	assertCode(t, Unit("com.example").Import("com.example.Example"), expected)
}
