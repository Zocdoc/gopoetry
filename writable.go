package poetrycs

type Writable interface {
	WriteCode(writer CodeWriter)
}