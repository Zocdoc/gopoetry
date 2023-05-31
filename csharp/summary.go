package csharp

import "fmt"

type SummaryDeclaration struct {
	description string
	params      []map[string]string
	returnType  *string
}

func Summary(description string) *SummaryDeclaration {
	return &SummaryDeclaration{
		description: description,
	}
}

func (self *SummaryDeclaration) AddParam(name string, description string) *SummaryDeclaration {
	self.params = append(self.params, map[string]string{name: description})
	return self
}

func (self *SummaryDeclaration) AddReturnType(returnType string) *SummaryDeclaration {
	self.returnType = &returnType
	return self
}

func (self *SummaryDeclaration) WriteCode(writer CodeWriter) {
	writer.Write("/// <summary>")
	writer.Eol()
	writer.Write(fmt.Sprintf("/// %s", self.description))
	writer.Eol()
	writer.Write("/// </summary>")
	writer.Eol()

	for _, param := range self.params {
		for name, description := range param {
			writer.Write(fmt.Sprintf("/// <param name=\"%s\">%s</param>", name, description))
			writer.Eol()
		}
	}

	if self.returnType != nil {
		writer.Write(fmt.Sprintf("/// <returns>%s</returns>", *self.returnType))
		writer.Eol()
	}
}
