package ruby

import "testing"

func TestClassSimple(t *testing.T) {
	expected := `
class MyClass
end
`
	assertCode(t, Class("MyClass"), expected)
}

func TestClassInherits(t *testing.T) {
	expected := `
class MyClass < Base
end
`
	assertCode(t, Class("MyClass").Inherits("Base"), expected)
}
