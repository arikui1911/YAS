package parser

import (
	"github.com/arikui1911/YAS/ast"
	"testing"
)

func TestParseIntLiteral(t *testing.T) {
	tree, err := ParseString("123", "(test)")
	if err != nil {
		t.Error(err)
		return
	}
	node, ok := tree.(*ast.IntLiteral)
	if !ok {
		t.Errorf("expect *ast.IntLiteral but %T", tree)
		return
	}
	if node.Value() != 123 {
		t.Errorf("expect 123 but %v", node.Value())
	}
}

func TestParseFloatLiteral(t *testing.T) {
	tree, err := ParseString("12.3", "(test)")
	if err != nil {
		t.Error(err)
		return
	}
	node, ok := tree.(*ast.FloatLiteral)
	if !ok {
		t.Errorf("expect *ast.FloatLiteral but %T", tree)
		return
	}
	if node.Value() != 12.3 {
		t.Errorf("expect 12.3 but %v", node.Value())
	}
}

func TestParseStringLiteral(t *testing.T) {
	tree, err := ParseString(`"Hello"`, "(test)")
	if err != nil {
		t.Error(err)
		return
	}
	node, ok := tree.(*ast.StringLiteral)
	if !ok {
		t.Errorf("expect *ast.StringLiteral but %T", tree)
		return
	}
	if node.Value() != "Hello" {
		t.Errorf("expect Hello but %v", node.Value())
	}
}

func TestParseVarRef(t *testing.T) {
	tree, err := ParseString(`hoge`, "(test)")
	if err != nil {
		t.Error(err)
		return
	}
	node, ok := tree.(*ast.VarRef)
	if !ok {
		t.Errorf("expect *ast.VarRef but %T", tree)
		return
	}
	if node.Identifier() != "hoge" {
		t.Errorf("expect hoge but %v", node.Identifier())
	}
}
