package poetrycs

import "testing"

func TestNamespaceEmpty(t *testing.T) {
	expected := `
namespace Bla
{
}
`
	assertCode(t, Namespace("Bla"), expected)
}

func TestNamespaceImportsAndClasses(t *testing.T) {
	expected := `
namespace Bla
{
    import Some.Namespace;
    import Other.Namespace;
    
    class MyClass
    {
    }
}
`
	namespace :=
		Namespace("Bla").
			Import("Some.Namespace").
			Import("Other.Namespace").
			AddClasses(Class("MyClass"))
	assertCode(t, namespace, expected)
}
