package java

import (
	"strings"
)

type CodeWriter interface {
	Begin()
	End()
	Eol()
	Write(code string)
}

func prefix(indentation int, numSpacesIndent int) string {
	tab := strings.Repeat(" ", numSpacesIndent)
	return strings.Repeat(tab, indentation)
}

type codeWriter struct {
	code              strings.Builder
	indentation       int
	newLine           bool
	numOfSpacesIndent int
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
	if self.newLine {
		self.newLine = false
	}
	self.Write("\n")
	self.newLine = true
}

func (self *codeWriter) Write(code string) {
	if self.newLine {
		self.code.WriteString(prefix(self.indentation, self.numOfSpacesIndent))
		self.newLine = false
	}
	self.code.WriteString(code)
}

func (self *codeWriter) Code() string {
	return self.code.String()
}

func CreateWriter(numOfSpacesIndent int) codeWriter {
	return codeWriter{code: strings.Builder{}, indentation: 0, newLine: true, numOfSpacesIndent: numOfSpacesIndent}
}
