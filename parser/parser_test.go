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

func TestParseMinusExpression(t *testing.T) {
	tree, err := ParseString(`-666`, "(test)")
	if err != nil {
		t.Error(err)
		return
	}
	me, ok := tree.(*ast.MinusExpression)
	if !ok {
		t.Errorf("expect *ast.MinusExpression but %T", tree)
		return
	}
	il, ok := me.Operand().(*ast.IntLiteral)
	if !ok {
		t.Errorf("expect *ast.IntLiteral but %T", me.Operand())
		return
	}
	if il.Value() != 666 {
		t.Errorf("expect 666 but %v", il.Value())
	}
}

func TestParseNotExpression(t *testing.T) {
	tree, err := ParseString(`!hoge`, "(test)")
	if err != nil {
		t.Error(err)
		return
	}
	ne, ok := tree.(*ast.NotExpression)
	if !ok {
		t.Errorf("expect *ast.NotExpression but %T", tree)
		return
	}
	vr, ok := ne.Operand().(*ast.VarRef)
	if !ok {
		t.Errorf("expect *ast.VarRef but %T", ne.Operand())
		return
	}
	if vr.Identifier() != "hoge" {
		t.Errorf("expect hoge but %v", vr.Identifier())
	}
}

func TestParseMultiplicationExpression(t *testing.T) {
	tree, err := ParseString(`123 * 456`, "(test)")
	if err != nil {
		t.Error(err)
		return
	}
	me, ok := tree.(*ast.MultiplicationExpression)
	if !ok {
		t.Errorf("expect *ast.MultiplicationExpression but %T", tree)
		return
	}
	lil, ok := me.Left().(*ast.IntLiteral)
	if !ok {
		t.Errorf("expect *ast.IntLiteral but %T", me.Left())
		return
	}
	if lil.Value() != 123 {
		t.Errorf("expect 123 but %v", lil.Value())
	}
	ril, ok := me.Right().(*ast.IntLiteral)
	if !ok {
		t.Errorf("expect *ast.IntLiteral but %T", me.Right())
		return
	}
	if ril.Value() != 456 {
		t.Errorf("expect 456 but %v", ril.Value())
	}
}
