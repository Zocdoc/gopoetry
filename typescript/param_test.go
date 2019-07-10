package typescript

import "testing"

func TestParamSimple(t *testing.T) {
	assertCode(t, Param("int", "intParam"), "intParam: int")
}

func TestParamDefault(t *testing.T) {
	assertCode(t, Param("int", "intParam").Default(Null), "intParam: int = null")
}
