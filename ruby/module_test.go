package ruby

import "testing"

func TestModuleSimple(t *testing.T) {
	expected := `
module MyModule
end
`
	assertCode(t, Module("MyModule"), expected)
}

func TestModuleDeclaration(t *testing.T) {
	expected := `
module MyModule
  def the_function()
  end
end
`
	assertCode(t, Module("MyModule").AddDeclarations(Method("the_function")), expected)
}
