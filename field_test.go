package poetrycs

import "testing"

func TestFieldPublicStatic(t *testing.T) {
	assertCode(t, Field("MyType", "myField").Public().Static(), "public static MyType myField;")
}

func TestFieldInitializer(t *testing.T) {
	assertCode(t, Field("int", "myField").Init(Int(3)), "int myField = 3;")
}