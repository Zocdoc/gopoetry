package ruby

import "testing"

func TestArg(t *testing.T) {
	expected := `theArg`
	assertCode(t, Arg("theArg"), expected)
}

func TestArgDefault(t *testing.T) {
	expected := `theArg = 'the value'`
	assertCode(t, Arg("theArg").Default(Str("the value")), expected)
}

func TestKeywordArg(t *testing.T) {
	expected := `theArg:`
	assertCode(t, KeywordArg("theArg"), expected)
}

func TestKeywordArgDefault(t *testing.T) {
	expected := `theArg: 'the value'`
	assertCode(t, KeywordArg("theArg").Default(Str("the value")), expected)
}
