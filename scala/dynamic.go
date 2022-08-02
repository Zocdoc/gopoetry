package scala

type WritableList struct {
	list []Writable
}

func (self *WritableList) Add(writable ...Writable) {
	self.list = append(self.list, writable...)
}

func Dynamic(dynamic func(code *WritableList)) []Writable {
	code := &WritableList{list: []Writable{}}
	dynamic(code)
	return code.list
}
