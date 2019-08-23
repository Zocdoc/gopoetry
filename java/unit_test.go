package java

import (
	"gopoetry/util"
	"testing"
)

func TestImport(t *testing.T) {
	expected := `
package com.example;

import com.example.Example;
`
	util.AssertCode(t, Unit("com.example").Import("com.example.Example"), expected)
}
