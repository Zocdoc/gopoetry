package swift

import (
	"strings"
)

type CodeWriter interface {
	Begin()
	End()
	Eol()
	Write(code string)
	OpenBlock()
	CloseBlock()
}

type codeWriter struct {
	code        strings.Builder
	indentation int
	newLine     bool
}

var _ CodeWriter = (*codeWriter)(nil)

func prefix(indentation int) string {
	tab := strings.Repeat(" ", 4)
	return strings.Repeat(tab, indentation)
}

func (cw *codeWriter) Begin() {
	cw.Eol()
	cw.OpenBlock()
}

func (cw *codeWriter) OpenBlock() {
	cw.Write("{")
	cw.Eol()
	cw.indentation += 1
}

func (cw *codeWriter) CloseBlock() {
	cw.indentation -= 1
	cw.Write("}")
}

func (cw *codeWriter) End() {
	cw.CloseBlock()
	cw.Eol()
}

func (cw *codeWriter) Eol() {
	cw.Write("\n")
	cw.newLine = true
}

func (cw *codeWriter) Write(code string) {
	if cw.newLine {
		cw.code.WriteString(prefix(cw.indentation))
		cw.newLine = false
	}
	cw.code.WriteString(code)
}

func (cw *codeWriter) Code() string {
	return cw.code.String()
}

func CreateWriter() codeWriter {
	return codeWriter{code: strings.Builder{}, indentation: 0, newLine: true}
}
