package yas

import (
	"fmt"
	"testing"
	"github.com/arikui1911/YAS/parser"
)

func TestRuntime(t *testing.T) {
	src := `123 / 0`
	tree, err := parser.ParseString(src, "(test)")
	if err != nil {
		t.Error(err)
		return
	}

	r := NewRuntime()
	v, err := r.Evaluate(tree)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("<<<%v>>>\n", v)
}