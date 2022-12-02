package typescript

import "testing"

func TestFunction(t *testing.T) {
	expected := `
function myFunction() {
}
  `

	fn := &FunctionDecl{
		Name: "myFunction",
	}

	assertCode(t, fn, expected)
}

func TestFunctionWithParams(t *testing.T) {
	expected := `
function add(x: number, y: number) {
}
  `

	fn := &FunctionDecl{
		Name: "add",
		Function: Function{
			params: []ParamDeclaration{
				*Param("number", "x"),
				*Param("number", "y"),
			},
		},
	}

	assertCode(t, fn, expected)
}

func TestReturnType(t *testing.T) {
	expected := `
function add(): void {
}
  `

	fn := &FunctionDecl{
		Name: "add",
		Function: Function{
			returnType: Code("void"),
		},
	}

	assertCode(t, fn, expected)
}

func TestAsyncFuncDecl(t *testing.T) {
	expected := `
async function asyncFunc() {
}
`

	fn := &FunctionDecl{
		Name: "asyncFunc",
		Function: Function{
			async: true,
		},
	}

	assertCode(t, fn, expected)
}

func TestExportedFn(t *testing.T) {
	expected := `
export function add(): void {
}
  `

	fn := &FunctionDecl{
		Name: "add",
		Function: Function{
			returnType: Code("void"),
		},
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

	fn := &FunctionDecl{
		Name: "add",
		Function: Function{
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

	fn := &FunctionDecl{
		Name: "id",
		Function: Function{
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
		},
	}
	assertCode(t, fn, expected)
}

func TestFunctionExpression(t *testing.T) {
	expected := `
const myFunc = function () {
    return 42;
};
`
	functionExperssion := &Function{
		body: BlockDeclaration{
			lines: []Writable{
				Code("return 42;"),
			},
		},
	}

	funcConst := Const("myFunc", nil, functionExperssion)
	assertCode(t, funcConst, expected)
}

func TestFunctionExpressionWithParams(t *testing.T) {
	expected := `
const add = function (a: number, b: number) {
    return a + b;
};
`
	functionExperssion := &Function{
		params: []ParamDeclaration{
			*Param("number", "a"),
			*Param("number", "b"),
		},
		body: BlockDeclaration{
			lines: []Writable{
				Code("return a + b;"),
			},
		},
	}

	funcConst := Const("add", nil, functionExperssion)
	assertCode(t, funcConst, expected)
}

func TestFunctionExpressionWithReturnValue(t *testing.T) {
	expected := `
const add = function (a: number, b: number): number {
    return a + b;
};
`
	functionExperssion := &Function{
		params: []ParamDeclaration{
			*Param("number", "a"),
			*Param("number", "b"),
		},
		returnType: Code("number"),
		body: BlockDeclaration{
			lines: []Writable{
				Code("return a + b;"),
			},
		},
	}

	funcConst := Const("add", nil, functionExperssion)
	assertCode(t, funcConst, expected)
}

func TestFunctionExpressionWithTypeConstructor(t *testing.T) {
	expected := `
const id = function <T>(x: T): T {
    return x;
};
`
	functionExperssion := &Function{
		typeConstructor: Code("T"),
		params: []ParamDeclaration{
			*Param("T", "x"),
		},
		returnType: Code("T"),
		body: BlockDeclaration{
			lines: []Writable{
				Code("return x;"),
			},
		},
	}

	funcConst := Const("id", nil, functionExperssion)
	assertCode(t, funcConst, expected)
}

func TestAddArrowFunction(t *testing.T) {
	expected := `
const f = <T>(_x: T): number => {
    return 2;
};
`

	functionExperssion := &Function{
		typeConstructor: Code("T"),
		arrowSyntax:     true,
		params: []ParamDeclaration{
			*Param("T", "_x"),
		},
		returnType: Code("number"),
		body: BlockDeclaration{
			lines: []Writable{
				Code("return 2;"),
			},
		},
	}

	funcConst := Const("f", nil, functionExperssion)
	assertCode(t, funcConst, expected)
}

func TestAddArrowFuncEmptyThunk(t *testing.T) {
	expected := `
const f = () => {
};
`

	functionExperssion := &Function{}
	functionExperssion.Arrow()

	funcConst := Const("f", nil, functionExperssion)
	assertCode(t, funcConst, expected)
}

func TestAsyncArrowFuncEmptyThunk(t *testing.T) {
	expected := `
const asyncF = async () => {
};
`

	functionExperssion := &Function{}
	functionExperssion.Arrow()
	functionExperssion.Async()
	funcConst := Const("asyncF", nil, functionExperssion)
	assertCode(t, funcConst, expected)
}
