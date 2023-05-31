package csharp

import "fmt"

type SummaryParam struct {
	name, description string
}

type SummaryDeclaration struct {
	description string
	params      []SummaryParam
	returnType  string
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

func (self *SummaryDeclaration) AddReturnType(returnType string) *SummaryDeclaration {
	self.returnType = returnType
	return self
}

func (self *SummaryDeclaration) WriteCode(writer CodeWriter) {
	if self.description == "" {
		return
	}

	writer.Write("/// <summary>")
	writer.Eol()
	writer.Write(fmt.Sprintf("/// %s", self.description))
	writer.Eol()
	writer.Write("/// </summary>")
	writer.Eol()

	for _, param := range self.params {
		writer.Write(fmt.Sprintf("/// <param name=\"%s\">%s</param>", param.name, param.description))
		writer.Eol()
	}

	if self.returnType != "" {
		writer.Write(fmt.Sprintf("/// <returns>%s</returns>", self.returnType))
		writer.Eol()
	}
}
