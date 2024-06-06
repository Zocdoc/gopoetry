package csharp

import "testing"

func TestMethodVoid(t *testing.T) {
	assertCode(t, Method("MyMethod"), "void MyMethod();")
}

func TestMethodReturns(t *testing.T) {
	assertCode(t, Method("MyMethod").Returns("Result"), "Result MyMethod();")
}

func TestMethodPublic(t *testing.T) {
	assertCode(t, Method("MyMethod").Public(), "public void MyMethod();")
}

func TestMethodParams(t *testing.T) {
	method := Method("MyMethod").Public()
	method.Param("int", "intParam")
	method.Param("string", "stringParam")
	assertCode(t, method, "public void MyMethod(int intParam, string stringParam);")
}

func TestMethodBody(t *testing.T) {
	expected := `
Result MyMethod()
{
}
`
	method := Method("MyMethod").Returns("Result")
	method.Body()
	assertCode(t, method, expected)
}

func TestMethodAttributed(t *testing.T) {
	expected := `
[MyAttribute]
void MyMethod();
`
	assertCode(t, Method("MyMethod").WithAttribute("MyAttribute"), expected)
}

func TestMethodSummary(t *testing.T) {
	expected := `
/// <summary>
/// my method summary
/// </summary>
void MyMethod();
`
	method := Method("MyMethod").Summary("my method summary")
	assertCode(t, method, expected)
}

func TestMethodSummaryWithParams(t *testing.T) {
	expected := `
/// <summary>
/// my method summary
/// </summary>
/// <param name="intParam">A simple int param</param>
/// <param name="stringParam"></param>
void MyMethod(int intParam, string stringParam);
`
	method := Method("MyMethod").Summary("my method summary")
	method.ParamWithDescription("int", "intParam", "A simple int param")
	method.Param("string", "stringParam")
	assertCode(t, method, expected)
}

func TestMethodSummaryWithReturns(t *testing.T) {
	expected := `
/// <summary>
/// my method summary
/// </summary>
/// <returns>my return summary</returns>
void MyMethod();
`
	method := Method("MyMethod").Summary("my method summary").ReturnsSummary("my return summary")
	assertCode(t, method, expected)
}

func TestMethodSummaryWithParamsAndReturns(t *testing.T) {
	expected := `
/// <summary>
/// my method summary
/// </summary>
/// <param name="intParam">A simple int param</param>
/// <param name="stringParam"></param>
/// <returns>my return summary</returns>
void MyMethod(int intParam, string stringParam);
`
	method := Method("MyMethod").Summary("my method summary").ReturnsSummary("my return summary")
	method.ParamWithDescription("int", "intParam", "A simple int param")
	method.Param("string", "stringParam")
	assertCode(t, method, expected)
}

func TestMethodSummaryXmlEscape(t *testing.T) {
	expected := `
/// <summary>
/// my method summary &lt;foo&gt;
/// </summary>
/// <param name="intParam">A simple int &lt;param&gt;</param>
/// <param name="intParam2"><see cref="int" /></param>
/// <param name="stringParam"></param>
/// <returns>my return summary &lt;bar&gt;</returns>
void MyMethod(int intParam, int intParam2, string stringParam);
`
	method := Method("MyMethod").Summary("my method summary <foo>").ReturnsSummary("my return summary <bar>")
	method.ParamWithDescription("int", "intParam", "A simple int <param>")
	method.ParamWithDescriptionAlreadyEscaped("int", "intParam2", "<see cref=\"int\" />")
	method.Param("string", "stringParam")
	assertCode(t, method, expected)
}

func TestMethodSummaryXmlEmbedded(t *testing.T) {
	expected := `
/// <summary>
/// my method summary &lt;foo&gt;
/// </summary>
/// <param name="intParam">A simple int &lt;param&gt;</param>
/// <param name="stringParam"></param>
/// <returns>my return summary &lt;bar&gt;</returns>
void MyMethod(int intParam, string stringParam);
`
	method := Method("MyMethod").Summary("my method summary <foo>").ReturnsSummary("my return summary <bar>")
	method.ParamWithDescription("int", "intParam", "A simple int <param>")
	method.Param("string", "stringParam")
	assertCode(t, method, expected)
}

func TestConstructor(t *testing.T) {
	expected := `
MyType(string myString)
{
}
`
	ctor := Constructor("MyType")
	ctor.Param("string", "myString")
	ctor.Body()
	assertCode(t, ctor, expected)
}

func TestConstructorWithBase(t *testing.T) {
	expected := `
MyType(string myString, bool myBool) : base(myString)
{
}
`
	ctor := Constructor("MyType")
	ctor.Param("string", "myString")
	ctor.Param("bool", "myBool")
	ctor.WithBase("myString")
	ctor.Body()
	assertCode(t, ctor, expected)
}

func TestMethodExpression(t *testing.T) {
	expected := `
Result MyMethod() => true;
`
	method := Method("MyMethod").Returns("Result")
	method.ExpressionBodiedMember("true")
	assertCode(t, method, expected)
}
