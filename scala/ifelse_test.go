package scala

import "testing"

func TestIfElseBasic(t *testing.T) {
	assertCode(t, If(true).Then(Code("true code")).Else(Code("false code")), "true code")
	assertCode(t, If(false).Then(Code("true code")).Else(Code("false code")), "false code")
}

func TestIfElseOmit(t *testing.T) {
	assertCode(t, If(true), "")
	assertCode(t, If(false), "")
}