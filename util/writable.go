package util

type Writable interface {
	WriteCode(writer CodeWriter)
}
