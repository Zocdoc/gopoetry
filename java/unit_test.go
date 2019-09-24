package java

import (
	"testing"
)

func TestImport(t *testing.T) {
	expected := `
package com.example;

import com.example.Example;
`
	AssertCode(t, Unit("com.example").Import("com.example.Example"), expected)
}
