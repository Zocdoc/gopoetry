package csharp

import "testing"

func TestSummary(t *testing.T) {
	expected := `
/// <summary>
/// my summary
/// </summary>
`
	summary := Summary("my summary")
	assertCode(t, summary, expected)
}

func TestSummaryWithParams(t *testing.T) {
	expected := `
/// <summary>
/// my summary
/// </summary>
/// <param name="intParam">A simple int param</param>
/// <param name="stringParam"></param>
`
	summary := Summary("my summary")
	summary.AddParam("intParam", "A simple int param")
	summary.AddParam("stringParam", "")
	assertCode(t, summary, expected)
}

func TestSummaryWithReturnType(t *testing.T) {
	expected := `
/// <summary>
/// my summary
/// </summary>
/// <returns>Result</returns>
`
	summary := Summary("my summary")
	summary.AddReturnType("Result")
	assertCode(t, summary, expected)
}

func TestSummaryWithParamsAndReturnType(t *testing.T) {
	expected := `
/// <summary>
/// my summary
/// </summary>
/// <param name="intParam">A simple int param</param>
/// <param name="stringParam"></param>
/// <returns>Result</returns>
`
	summary := Summary("my summary")
	summary.AddParam("intParam", "A simple int param")
	summary.AddParam("stringParam", "")
	summary.AddReturnType("Result")
	assertCode(t, summary, expected)
}
