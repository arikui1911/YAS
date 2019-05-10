package parser

import (
	"github.com/arikui1911/YAS/ast"
	"os"
	"testing"
)

func TestParser(t *testing.T) {
	src := `123`
	tree, err := ParseString(src, "(test)")
	if err != nil {
		t.Error(err)
	} else {
		ast.Dump(os.Stdout, tree)
	}
}
