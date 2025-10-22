package typescript

import (
	"fmt"
	"strings"
)

type CommentBlockDeclaration struct {
	comments []string
}

func CommentBlock(lines ...string) *CommentBlockDeclaration {
	return &CommentBlockDeclaration{
		comments: lines,
	}
}

func cleanCommentLine(line string) string {
	return strings.Replace(line, "*/", "* /", -1)
}

func (self *CommentBlockDeclaration) AddComments(lines ...string) *CommentBlockDeclaration {
	for _, line := range lines {
		for _, comment := range strings.Split(strings.Replace(line, "\r\n", "\n", -1), "\n") {
			cleanComment := cleanCommentLine(comment)
			self.comments = append(self.comments, cleanComment)
		}
	}
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
