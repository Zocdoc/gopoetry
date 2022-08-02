package csharp

import (
	"strconv"
)

type WritableCode struct {
	code string
}

func (self *WritableCode) WriteCode(writer CodeWriter) {
	writer.Write(self.code)
}

func Code(code string) *WritableCode {
	return &WritableCode{code: code}
}

func C(code string) *WritableCode {
	return Code(code)
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
