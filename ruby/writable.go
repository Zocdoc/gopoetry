package ruby

type Writable interface {
	WriteCode(writer CodeWriter)
}
