package csharp

import "testing"

func TestParamSimple(t *testing.T) {
	assertCode(t, Param("int", "intParam"), "int intParam")
}

func TestParamDefault(t *testing.T) {
	assertCode(t, Param("int", "intParam").Default(Null), "int intParam = null")
}
