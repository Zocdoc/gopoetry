package csharp

type BaseStatement struct {
	arguments []string
}

func Base(args []string) *BaseStatement {
	return &BaseStatement{
		arguments: args,
	}
}

func (self *BaseStatement) WriteCode(writer CodeWriter) {
	writer.Write("base(")
	for i, arg := range self.arguments {
		writer.Write(arg)
		if i < len(self.arguments)-1 {
			writer.Write(", ")
		}
	}
	writer.Write(")")
}
