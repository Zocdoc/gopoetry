package scala

type LazyWritable struct {
	writableGetter func() Writable
}

func Lazy(writableGetter func() Writable) *LazyWritable {
	return &LazyWritable{ writableGetter: writableGetter }
}

func (self *LazyWritable) WriteCode(writer CodeWriter) {
	writable := self.writableGetter()
	writable.WriteCode(writer)
}