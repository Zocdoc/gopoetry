package csharp

type SingleLineExpressionDeclaration struct {
	Expression string
}

func SingleLineExpression(expression string) *SingleLineExpressionDeclaration {
	return &SingleLineExpressionDeclaration{Expression: expression}
}

func (ed *SingleLineExpressionDeclaration) WriteCode(writer CodeWriter) {
	writer.Write(ed.Expression)
}
