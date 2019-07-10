package typescript

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
	expected := `import { foo, bar } from 'bar-package';
import * as buzz from 'buzz-package';

/// <reference path="foo.ts" />
namespace Bla
{
    class MyClass
    {
    }
}
`
	namespace :=
		Namespace("Bla").WithReference("foo.ts").
			NamedImport("bar-package", "foo", "bar").
			DefaultImport("buzz-package", "buzz").
			AddDeclarations(Class("MyClass"))
	assertCode(t, namespace, expected)
}
