package typescript

import "testing"

func TestMethodVoid(t *testing.T) {
	assertCode(t, Method("MyMethod"), "MyMethod(): void;")
}

func TestMethodReturns(t *testing.T) {
	assertCode(t, Method("MyMethod").Returns("Result"), "MyMethod(): Result;")
}

func TestMethodPublic(t *testing.T) {
	assertCode(t, Method("MyMethod").Public(), "public MyMethod(): void;")
}

func TestMethodParams(t *testing.T) {
	method := Method("MyMethod").Public()
	method.Param("int", "intParam")
	method.Param("string", "stringParam")
	assertCode(t, method, "public MyMethod(intParam: int, stringParam: string): void;")
}

func TestMethodParamsAsync(t *testing.T) {
	method := Method("MyMethod").Public().Async().Returns("Promise<void>")
	method.Param("int", "intParam")
	method.Param("string", "stringParam")
	assertCode(t, method, "public async MyMethod(intParam: int, stringParam: string): Promise<void>;")
}

func TestMethodBody(t *testing.T) {
	expected := `
MyMethod(): Result
{
}
`
	method := Method("MyMethod").Returns("Result")
	method.Body()
	assertCode(t, method, expected)
}
