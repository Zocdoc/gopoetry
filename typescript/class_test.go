package typescript

import (
	"testing"
)

func TestClassBasic(t *testing.T) {
	expected := `
class MyClass
{
}
`
	assertCode(t, Class("MyClass"), expected)
}

func TestClassExtends(t *testing.T) {
	expected := `
class MyClass extends MyParent
{
}
`
	assertCode(t, Class("MyClass").Extends("MyParent"), expected)
}

func TestClassImplements(t *testing.T) {
	expected := `
class MyClass implements MyParent
{
}
`
	assertCode(t, Class("MyClass").Implements("MyParent"), expected)
}

func TestClassImplementsAndExtends(t *testing.T) {
	expected := `
class MyClass extends Single implements Double, Triple
{
}
`
	assertCode(t, Class("MyClass").Extends("Single").Implements("Double", "Triple"), expected)
}

func TestClassModifier(t *testing.T) {
	expected := `
private class MyClass
{
}
`
	assertCode(t, Class("MyClass").Private(), expected)
}

func TestClassMethod(t *testing.T) {
	expected := `
class MyClass
{
    MyMethod(): void
    {
    }
}
`
	class := Class("MyClass")
	class.Method("MyMethod").Body()
	assertCode(t, class, expected)
}

func TestClassProperty(t *testing.T) {
	expected := `
class MyClass
{
    MyProperty: Result;
}
`
	class := Class("MyClass")
	_ = class.Property("Result", "MyProperty")
	assertCode(t, class, expected)
}

func TestClassConstructor(t *testing.T) {
	expected := `
export class MyClass
{
    constructor(private name: string = 'bar')
    {
    }
}
`
	class := Class("MyClass").Export()
	class.Constructor().
		Param("string", "name").Default(Str("bar")).Private()
	assertCode(t, class, expected)
}

func TestClassConstructorWithBody(t *testing.T) {
	expected := `
export class MyClass
{
    constructor(private name: string = 'bar')
    {
        let foo = 'bar';
    }
}
`
	class := Class("MyClass").Export()
	ctor := class.Constructor()
	ctor.Param("string", "name").Default(Str("bar")).Private()

	ctor.Body("let foo = 'bar';")
	assertCode(t, class, expected)
}
