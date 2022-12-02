package typescript

import "testing"

func TestFunction(t *testing.T) {
	expected := `
function myFunction() {
}
  `

	fn := DeclareFunction("myFunction")
	assertCode(t, fn, expected)
}

func TestFunctionWithParams(t *testing.T) {
	expected := `
function add(x: number, y: number) {
}
  `

	fn := DeclareFunction("add")
	fn.GetExpression().AddParams(
		*Param("number", "x"),
		*Param("number", "y"),
	)

	assertCode(t, fn, expected)
}

func TestReturnType(t *testing.T) {
	expected := `
function add(): void {
}
  `

	fn := DeclareFunction("add")
	fn.GetExpression().ReturnType(C("void"))

	assertCode(t, fn, expected)
}

func TestAsyncFuncDecl(t *testing.T) {
	expected := `
async function asyncFunc() {
}
`

	fn := DeclareFunction("asyncFunc")
	fn.GetExpression().Async()
	assertCode(t, fn, expected)
}

func TestExportedFn(t *testing.T) {
	expected := `
export function add(): void {
}
  `

	fn := &FunctionDecl{
		Name: "add",
		FunctionExpression: FunctionExpression{
			returnType: Code("void"),
		},
	}

	fn = DeclareFunction("add").Export()
	fn.GetExpression().ReturnType(C("void"))

	assertCode(t, fn, expected)
}

func TestBody(t *testing.T) {
	expected := `
export function add(x: number, y: number): number {
    return x + y;
}
  `

	fn := DeclareFunction("add").Export()
	fn.GetExpression().
		AddParams(
			*Param("number", "x"),
			*Param("number", "y"),
		).
		ReturnType(C("number")).
		AppendToBody(C("return x + y;"))

	assertCode(t, fn, expected)
}

func TestTypeConstructor(t *testing.T) {
	expected := `
function id<T>(thing: T): T {
    return thing;
}
  `
	fn := DeclareFunction("id")
	fn.GetExpression().
		TypeConstructor(C("T")).
		AddParams(*Param("T", "thing")).
		ReturnType(C("T")).
		AppendToBody(C("return thing;"))

	assertCode(t, fn, expected)
}

func TestFunctionExpression(t *testing.T) {
	expected := `
const myFunc = function () {
    return 42;
};
`

	functionExperssion := NewFunctionExpression().
		AppendToBody(C("return 42;"))

	funcConst := Const("myFunc", nil, functionExperssion)
	assertCode(t, funcConst, expected)
}

func TestFunctionExpressionWithParams(t *testing.T) {
	expected := `
const add = function (a: number, b: number) {
    return a + b;
};
`
	functionExperssion := NewFunctionExpression().
		AddParams(
			*Param("number", "a"),
			*Param("number", "b"),
		).
		AppendToBody(
			C("return a + b;"),
		)

	funcConst := Const("add", nil, functionExperssion)
	assertCode(t, funcConst, expected)
}

func TestFunctionExpressionWithReturnValue(t *testing.T) {
	expected := `
const add = function (a: number, b: number): number {
    return a + b;
};
`

	functionExperssion := NewFunctionExpression().
		AddParams(
			*Param("number", "a"),
			*Param("number", "b"),
		).
		ReturnType(C("number")).
		AppendToBody(C("return a + b;"))

	funcConst := Const("add", nil, functionExperssion)
	assertCode(t, funcConst, expected)
}

func TestFunctionExpressionWithTypeConstructor(t *testing.T) {
	expected := `
const id = function <T>(x: T): T {
    return x;
};
`

	functionExperssion := NewFunctionExpression().
		TypeConstructor(C("T")).
		AddParams(*Param("T", "x")).
		ReturnType(C("T")).
		AppendToBody(C("return x;"))

	funcConst := Const("id", nil, functionExperssion)
	assertCode(t, funcConst, expected)
}

func TestAddArrowFunction(t *testing.T) {
	expected := `
const f = <T>(_x: T): number => {
    return 2;
};
`

	functionExperssion := NewFunctionExpression().
		TypeConstructor(C("T")).
		Arrow().
		AddParams(*Param("T", "_x")).
		ReturnType(C("number")).
		AppendToBody(C("return 2;"))

	funcConst := Const("f", nil, functionExperssion)
	assertCode(t, funcConst, expected)
}

func TestAddArrowFuncEmptyThunk(t *testing.T) {
	expected := `
const f = () => {
};
`

	functionExperssion := NewFunctionExpression().Arrow()

	funcConst := Const("f", nil, functionExperssion)
	assertCode(t, funcConst, expected)
}

func TestAsyncArrowFuncEmptyThunk(t *testing.T) {
	expected := `
const asyncF = async () => {
};
`

	functionExperssion := NewFunctionExpression().
		Arrow().
		Async()

	funcConst := Const("asyncF", nil, functionExperssion)
	assertCode(t, funcConst, expected)
}

func TestAsyncArrowFuncArg(t *testing.T) {
	expected := `
const f = async function <T>(f: () => T): T {
    return f();
};
`

	functionExperssion := NewFunctionExpression().
		Async().
		TypeConstructor(C("T")).
		AddParams(*Param("() => T", "f")).
		ReturnType(C("T")).
		AppendToBody(C("return f();"))

	funcConst := Const("f", nil, functionExperssion)
	assertCode(t, funcConst, expected)
}
