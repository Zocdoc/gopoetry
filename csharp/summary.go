package csharp

import (
	"fmt"
	"strings"
)

type SummaryParam struct {
	name, description string
	isEscaped         bool
}

type SummaryDeclaration struct {
	description string
	params      []SummaryParam
	returns     string
}

func Summary(description string) *SummaryDeclaration {
	return &SummaryDeclaration{
		description: description,
	}
}

func (self *SummaryDeclaration) AddDescription(description string) *SummaryDeclaration {
	self.description = description
	return self
}

func (self *SummaryDeclaration) AddParam(name string, description string) *SummaryDeclaration {
	self.params = append(self.params, SummaryParam{name, description, false})
	return self
}

func (self *SummaryDeclaration) AddParamAlreadyEscaped(name string, description string) *SummaryDeclaration {
	self.params = append(self.params, SummaryParam{name, description, true})
	return self
}

func (self *SummaryDeclaration) AddReturns(returns string) *SummaryDeclaration {
	self.returns = returns
	return self
}

func (self *SummaryDeclaration) writeMultiLine(writer CodeWriter, description string) {
	for _, line := range strings.Split(strings.Replace(description, "\r\n", "\n", -1), "\n") {
		if line != "" {
			writer.Write(fmt.Sprintf("/// %s", line))
			writer.Eol()
		}
	}
}

func (self *SummaryDeclaration) WriteCode(writer CodeWriter) {
	if self.description == "" {
		return
	}

	writer.Write("/// <summary>")
	writer.Eol()

	self.writeMultiLine(writer, xmlEncode(self.description))

	writer.Write("/// </summary>")
	writer.Eol()

	for _, param := range self.params {
		desc := param.description
		if !param.isEscaped {
			desc = xmlEncode(desc)
		}

		if strings.Contains(desc, "\n") {
			writer.Write(fmt.Sprintf("/// <param name=\"%s\">", param.name))
			writer.Eol()

			self.writeMultiLine(writer, desc)

			writer.Write("/// </param>")
		} else {
			writer.Write(fmt.Sprintf("/// <param name=\"%s\">%s</param>", param.name, desc))
		}

		writer.Eol()
	}

	if self.returns != "" {
		writer.Write(fmt.Sprintf("/// <returns>%s</returns>", xmlEncode(self.returns)))
		writer.Eol()
	}
}

func xmlEncode(s string) string {
	s = strings.Replace(s, "&", "&amp;", -1)
	s = strings.Replace(s, "<", "&lt;", -1)
	s = strings.Replace(s, ">", "&gt;", -1)
	s = strings.Replace(s, "\"", "&quot;", -1)
	s = strings.Replace(s, "'", "&apos;", -1)
	return s
}
