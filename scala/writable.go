package scala

type Writable interface {
	WriteCode(writer CodeWriter)
}

func filterNils(writables []Writable) []Writable {
	result := []Writable{}
	for _, w := range writables {
		if w != nil {
			result = append(result, w)
		}
	}
	return result
}
