package scala

import (
	"strings"
)

type CodeWriter interface {
	Begin()
	End()
	Eol()
	Write(code string)
}

func prefix(indentation int) string {
	tab := strings.Repeat(" ", 2)
	return strings.Repeat(tab, indentation)
}

type codeWriter struct {
	code        strings.Builder
	indentation int
	newLine     bool
}

func (self *codeWriter) Begin() {
	self.Write(" {")
	self.Eol()
	self.indentation += 1
}

func (self *codeWriter) End() {
	self.indentation -= 1
	self.Write("}")
	self.Eol()
}

func (self *codeWriter) Eol() {
	self.Write("\n")
	self.newLine = true
}

func (self *codeWriter) Write(code string) {
	if self.newLine {
		self.code.WriteString(prefix(self.indentation))
		self.newLine = false
	}
	self.code.WriteString(code)
}

func (self *codeWriter) Code() string {
	return self.code.String()
}

func CreateWriter() codeWriter {
	return codeWriter{code: strings.Builder{}, indentation: 0, newLine: true}
}
