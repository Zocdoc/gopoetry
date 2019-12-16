package ruby

import "testing"

func TestRequireSimple(t *testing.T) {
	assertCode(t, Require("net/http"), `require "net/http"`)
}