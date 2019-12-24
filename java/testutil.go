package java

import (
	"gotest.tools/assert"
	"strings"
	"testing"
)

func AssertCode(t *testing.T, value Writable, expected string) {
	writer := CreateWriter(2)
	value.WriteCode(&writer)
	code := writer.Code()
	assert.Equal(t, strings.TrimSpace(code), strings.TrimSpace(expected))
}
