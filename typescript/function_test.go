package typescript

import "testing"

func TestFunction(t *testing.T) {
	expected := `
function myFunction() {
}
	`

	fn := &Function{
		Name: "myFunction",
	}

	assertCode(t, fn, expected)
}

func TestFunctionWithParams(t *testing.T) {
	expected := `
function add(x: number, y: number) {
}
	`

	fn := &Function{
		Name: "add",
		params: []ParamDeclaration{
			*Param("number", "x"),
			*Param("number", "y"),
		},
	}

	assertCode(t, fn, expected)
}

func TestReturnType(t *testing.T) {
	expected := `
function add(): void {
}
	`

	fn := &Function{
		Name:       "add",
		returnType: Code("void"),
	}

	assertCode(t, fn, expected)
}

func TestExportedFn(t *testing.T) {
	expected := `
export function add(): void {
}
	`

	fn := &Function{
		Name:       "add",
		returnType: Code("void"),
	}

	fn.Export()

	assertCode(t, fn, expected)
}

func TestBody(t *testing.T) {
	expected := `
export function add(x: number, y: number): number {
    return x + y;
}
	`

	fn := &Function{
		Name: "add",
		params: []ParamDeclaration{
			*Param("number", "x"),
			*Param("number", "y"),
		},
		returnType: Code("number"),
		body: BlockDeclaration{
			lines: []Writable{
				Code("return x + y;"),
			},
		},
	}

	fn.Export()

	assertCode(t, fn, expected)
}

func TestTypeConstructor(t *testing.T) {
	expected := `
function id<T>(thing: T): T {
    return thing;
}
	`

	fn := &Function{
		Name:            "id",
		typeConstructor: Code("T"),
		params: []ParamDeclaration{
			*Param("T", "thing"),
		},
		returnType: Code("T"),
		body: BlockDeclaration{
			lines: []Writable{
				Code("return thing;"),
			},
		},
	}
	assertCode(t, fn, expected)
}
