package scala

type Writable interface {
	WriteCode(writer CodeWriter)
}
