package typescript

type IfBlock struct {
	condition          Writable
	ifBlock, elseBlock BlockDeclaration
}
