package csharp

type Writable interface {
	WriteCode(writer CodeWriter)
}
