package csharp

import (
	"testing"
)

func TestRecordBasic(t *testing.T) {
	expected := `
record MyRecord
{
}
`
	assertCode(t, Record("MyRecord"), expected)
}

func TestRecordInherits(t *testing.T) {
	expected := `
record MyRecord : MyParent
{
}
`
	assertCode(t, Record("MyRecord").Inherits("MyParent"), expected)
}

func TestRecordModifier(t *testing.T) {
	expected := `
private record MyRecord
{
}
`
	assertCode(t, Record("MyRecord").Private(), expected)
}

func TestMultipleRecordModifier(t *testing.T) {
	expected := `
private abstract record MyRecord
{
}
`
	assertCode(t, Record("MyRecord").Private().Abstract(), expected)
}

func TestRecordAttributed(t *testing.T) {
	expected := `
[MyAttribute]
record MyRecord
{
}
`
	assertCode(t, Record("MyRecord").WithAttribute("MyAttribute"), expected)
}

func TestRecordMultipleAttributes(t *testing.T) {
	expected := `
[MyAttribute1, MyAttribute2]
record MyRecord
{
}
`
	assertCode(t, Record("MyRecord").WithAttribute("MyAttribute1").WithAttribute("MyAttribute2"), expected)
}

func TestRecordMethod(t *testing.T) {
	expected := `
record MyRecord
{
    void MyMethod()
    {
    }
}
`
	record := Record("MyRecord")
	record.Method("MyMethod").Body()
	assertCode(t, record, expected)
}

func TestRecordProperty(t *testing.T) {
	expected := `
record MyRecord
{
    Result MyProperty
    {
        get;
        set;
    }
}
`
	record := Record("MyRecord")
	property := record.Property("Result", "MyProperty")
	property.Get()
	property.Set()
	assertCode(t, record, expected)
}

func TestRecordPropertyWithInitializer(t *testing.T) {
	expected := `
record MyRecord
{
    Result MyProperty
    {
        get;
        set;
    }
     = "bar";
}
`
	record := Record("MyRecord")
	property := record.Property("Result", "MyProperty")
	property.Get()
	property.Set()
	property.Initialize(Str("bar"))
	assertCode(t, record, expected)
}

func TestRecordWithSummary(t *testing.T) {
	expected := `
/// <summary>
/// my record summary
/// </summary>
record MyRecord
{
}
`
	assertCode(t, Record("MyRecord").Summary("my record summary"), expected)
}
