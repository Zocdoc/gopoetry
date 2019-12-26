package scala

func WriteMembers(self CodeWriter, members []Writable) {
	if len(members) > 0 {
		self.Write("{")
		self.Eol()
		self.Indent()
		for _, member := range members {
			member.WriteCode(self)
		}
		self.UnIndent()
		self.Write("}")
		self.Eol()
	}
}
