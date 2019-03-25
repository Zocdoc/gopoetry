package poetrycs

import "testing"

func TestParamSimple(t *testing.T) {
	assertCode(t, Param("int", "intParam"), "int intParam")
}