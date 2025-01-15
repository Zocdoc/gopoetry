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

func TestSummaryMultiLined(t *testing.T) {
	expected := `
/// <summary>
/// my summary
/// my second line
/// my third line
/// </summary>
`
	summary := Summary("my summary\nmy second line\r\nmy third line\n")
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

func TestSummaryWithParamsMultiLined(t *testing.T) {
	expected := `
/// <summary>
/// my summary
/// </summary>
/// <param name="intParam">
/// A simple int param
/// my second param line
/// my third param line
/// </param>
/// <param name="stringParam"></param>
`
	summary := Summary("my summary")
	summary.AddParam("intParam", "A simple int param\nmy second param line\r\nmy third param line\n")
	summary.AddParam("stringParam", "")
	assertCode(t, summary, expected)
}

func TestSummaryWithReturns(t *testing.T) {
	expected := `
/// <summary>
/// my summary
/// </summary>
/// <returns>my return summary</returns>
`
	summary := Summary("my summary")
	summary.AddReturns("my return summary")
	assertCode(t, summary, expected)
}

func TestSummaryWithParamsAndReturns(t *testing.T) {
	expected := `
/// <summary>
/// my summary
/// </summary>
/// <param name="intParam">A simple int param</param>
/// <param name="stringParam"></param>
/// <returns>my return summary</returns>
`
	summary := Summary("my summary")
	summary.AddParam("intParam", "A simple int param")
	summary.AddParam("stringParam", "")
	summary.AddReturns("my return summary")
	assertCode(t, summary, expected)
}

func TestEmptySummary(t *testing.T) {
	expected := ``
	summary := Summary("")

	assertCode(t, summary, expected)
}

func TestSummaryWithOnlyParamDescription(t *testing.T) {
	expected := `
/// <summary>
/// </summary>
/// <param name="myParam">param description</param>
`
	summary := Summary("")
	summary.AddParam("myParam", "param description")
	assertCode(t, summary, expected)
}
