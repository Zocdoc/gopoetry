package csharp

import (
	"fmt"
	"strings"
)

type SummaryParam struct {
	name, description string
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
	self.params = append(self.params, SummaryParam{name, description})
	return self
}

func (self *SummaryDeclaration) AddReturns(returns string) *SummaryDeclaration {
	self.returns = returns
	return self
}

func (self *SummaryDeclaration) WriteCode(writer CodeWriter) {
	if self.description == "" {
		return
	}

	writer.Write("/// <summary>")
	writer.Eol()

	for _, line := range strings.Split(self.description, "\n") {
		if line != "" {
			writer.Write(fmt.Sprintf("/// %s", line))
			writer.Eol()
		}
	}

	writer.Write("/// </summary>")
	writer.Eol()

	for _, param := range self.params {
		writer.Write(fmt.Sprintf("/// <param name=\"%s\">%s</param>", param.name, param.description))
		writer.Eol()
	}

	if self.returns != "" {
		writer.Write(fmt.Sprintf("/// <returns>%s</returns>", self.returns))
		writer.Eol()
	}
}
