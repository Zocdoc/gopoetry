package scala

import (
	"fmt"
	"strconv"
)

type WritableCode struct {
	code string
	eol  int
}

func Code(codeFormat string, args ...interface{}) *WritableCode {
	return &WritableCode{code: fmt.Sprintf(codeFormat, args...), eol: 0}
}

func Line(codeFormat string, args ...interface{}) *WritableCode {
	return Code(codeFormat, args...).Eol()
}

func CodeIf(condition bool, codeFormat string, args ...interface{}) *WritableCode {
	if condition {
		return Code(codeFormat, args...)
	} else {
		return NoCode()
	}
}

func LineIf(condition bool, codeFormat string, args ...interface{}) *WritableCode {
	if condition {
		return Line(codeFormat, args...)
	} else {
		return NoCode()
	}
}

func (self *WritableCode) Eol() *WritableCode {
	self.eol = self.eol + 1
	return self
}

func (self *WritableCode) WriteCode(writer CodeWriter) {
	if self.code != "" {
		writer.Write(self.code)
	}
	for i := 0; i < self.eol; i++ {
		writer.Eol()
	}
}

func NoCode() *WritableCode {
	return &WritableCode{code: "", eol: 0}
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