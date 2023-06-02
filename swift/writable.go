package swift

type Writable interface {
	WriteCode(writer CodeWriter)
}
