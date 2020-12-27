package scala

import (
"strconv"
)

type WritableCode struct {
	code string
	eol bool
}

func Code(code string) *WritableCode {
	return &WritableCode{code: code, eol: false}
}

func Line(code string) *WritableCode {
	return &WritableCode{code: code, eol: true}
}

func (self *WritableCode) WriteCode(writer CodeWriter) {
	writer.Write(self.code)
	if self.eol {
		writer.Eol()
	}
}

func Int(value int) *WritableCode {
	return Code(strconv.Itoa(value))
}

func Str(value string) *WritableCode {
	return Code("\"" + value + "\"")
}

var True = Code("true")

var False = Code("false")

var Null = Code("null")