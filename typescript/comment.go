package typescript

import (
	"fmt"
)

type CommentBlockDeclaration struct {
	comments []string
}

func CommentBlock(lines ...string) *CommentBlockDeclaration {
	return &CommentBlockDeclaration{
		comments: lines,
	}
}

func (self *CommentBlockDeclaration) AddComments(lines ...string) *CommentBlockDeclaration {
	self.comments = append(self.comments, lines...)
	return self
}

func (self *CommentBlockDeclaration) writeMultiLine(writer CodeWriter) {
	for _, line := range self.comments {
		writer.Write(fmt.Sprintf(" * %s", line))
		writer.Eol()
	}
}

func (self *CommentBlockDeclaration) WriteComments(writer CodeWriter) {
	if len(self.comments) == 0 {
		return
	}

	if len(self.comments) == 1 {
		writer.Write(fmt.Sprintf("/** %s */", self.comments[0]))
		writer.Eol()
		return
	}
	
	writer.Write("/**")
	writer.Eol()

	self.writeMultiLine(writer)

	writer.Write(" */")
	writer.Eol()
}
