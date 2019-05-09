package parser

import (
	"testing"
)

func TestParser(t *testing.T) {
	src := `123`
	err := ParseString(src, "(test)")
	if err != nil {
		t.Error(err)
	}
}
