package graphql

type Writable interface {
	WriteCode(writer CodeWriter)
}
