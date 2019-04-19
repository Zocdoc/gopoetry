package csharp

import "testing"

func TestParamSimple(t *testing.T) {
	assertCode(t, Param("int", "intParam"), "int intParam")
}

func TestParamDefault(t *testing.T) {
	assertCode(t, Param("int", "intParam").Default(Null), "int intParam = null")
}

func TestParamAttributed(t *testing.T) {
	assertCode(t, Param("int", "intParam").Default(Null).WithAttribute("MyAttribute"), "[MyAttribute] int intParam = null")
}
