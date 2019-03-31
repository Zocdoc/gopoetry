package csharp

import "testing"

func TestNamespaceEmpty(t *testing.T) {
	expected := `
namespace Bla
{
}
`
	assertCode(t, Namespace("Bla"), expected)
}

func TestNamespaceUsingsAndClasses(t *testing.T) {
	expected := `
namespace Bla
{
    using Some.Namespace;
    using Other.Namespace;
    
    class MyClass
    {
    }
}
`
	namespace :=
		Namespace("Bla").
			Using("Some.Namespace").
			Using("Other.Namespace").
			AddClasses(Class("MyClass"))
	assertCode(t, namespace, expected)
}
