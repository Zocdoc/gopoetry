package swift

import (
	"strings"
	"testing"

	"gotest.tools/assert"
)

func assertCode(t *testing.T, value Writable, expected string) {
	writer := CreateWriter()
	value.WriteCode(&writer)
	code := writer.Code()
	assert.Equal(t, strings.TrimSpace(code), strings.TrimSpace(expected))
}
