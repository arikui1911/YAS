package parser

import (
	"github.com/arikui1911/YAS/ast"
	"testing"
)

func testIntLiteralNode(t *testing.T, subject ast.Node, expectValue int) {
	il, ok := subject.(*ast.IntLiteral)
	if !ok {
		t.Errorf("expect *ast.IntLiteral but %T", subject)
		return
	}
	if il.Value() != expectValue {
		t.Errorf("expect %v but %v", expectValue, il.Value())
	}
}

func TestParseIntLiteral(t *testing.T) {
	tree, err := ParseString("123", "(test)")
	if err != nil {
		t.Error(err)
		return
	}
	testIntLiteralNode(t, tree, 123)
}

func testFloatLiteralNode(t *testing.T, subject ast.Node, expectValue float64) {
	fl, ok := subject.(*ast.FloatLiteral)
	if !ok {
		t.Errorf("expect *ast.FloatLiteral but %T", subject)
		return
	}
	if fl.Value() != expectValue {
		t.Errorf("expect %v but %v", expectValue, fl.Value())
	}
}

func TestParseFloatLiteral(t *testing.T) {
	tree, err := ParseString("12.3", "(test)")
	if err != nil {
		t.Error(err)
		return
	}
	testFloatLiteralNode(t, tree, 12.3)
}

func testStringLiteralNode(t *testing.T, subject ast.Node, expectValue string) {
	sl, ok := subject.(*ast.StringLiteral)
	if !ok {
		t.Errorf("expect *ast.StringLiteral but %T", subject)
		return
	}
	if sl.Value() != expectValue {
		t.Errorf("expect %v but %v", expectValue, sl.Value())
	}
}

func TestParseStringLiteral(t *testing.T) {
	tree, err := ParseString(`"Hello"`, "(test)")
	if err != nil {
		t.Error(err)
		return
	}
	testStringLiteralNode(t, tree, "Hello")
}

func testVarRefNode(t *testing.T, subject ast.Node, expectIdentifier string) {
	vr, ok := subject.(*ast.VarRef)
	if !ok {
		t.Errorf("expect *ast.VarRef but %T", subject)
		return
	}
	if vr.Identifier() != expectIdentifier {
		t.Errorf("expect %v but %v", expectIdentifier, vr.Identifier())
	}
}

func TestParseVarRef(t *testing.T) {
	tree, err := ParseString(`hoge`, "(test)")
	if err != nil {
		t.Error(err)
		return
	}
	testVarRefNode(t, tree, "hoge")
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
	testIntLiteralNode(t, me.Operand(), 666)
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
	testVarRefNode(t, ne.Operand(), "hoge")
}

func TestParseAdditionExpression(t *testing.T) {
	tree, err := ParseString(`12 + 34`, "(test)")
	if err != nil {
		t.Error(err)
		return
	}
	me, ok := tree.(*ast.AdditionExpression)
	if !ok {
		t.Errorf("expect *ast.AdditionExpression but %T", tree)
		return
	}
	testIntLiteralNode(t, me.Left(), 12)
	testIntLiteralNode(t, me.Right(), 34)
}

func TestParseSubtractionExpression(t *testing.T) {
	tree, err := ParseString(`56 - 78`, "(test)")
	if err != nil {
		t.Error(err)
		return
	}
	me, ok := tree.(*ast.SubtractionExpression)
	if !ok {
		t.Errorf("expect *ast.SubtractionExpression but %T", tree)
		return
	}
	testIntLiteralNode(t, me.Left(), 56)
	testIntLiteralNode(t, me.Right(), 78)
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
	testIntLiteralNode(t, me.Left(), 123)
	testIntLiteralNode(t, me.Right(), 456)
}

func TestParseDivisionExpression(t *testing.T) {
	tree, err := ParseString(`666 / 3`, "(test)")
	if err != nil {
		t.Error(err)
		return
	}
	de, ok := tree.(*ast.DivisionExpression)
	if !ok {
		t.Errorf("expect *ast.DivisionExpression but %T", tree)
		return
	}
	testIntLiteralNode(t, de.Left(), 666)
	testIntLiteralNode(t, de.Right(), 3)
}

func TestParseModuloExpression(t *testing.T) {
	tree, err := ParseString(`111 % 222`, "(test)")
	if err != nil {
		t.Error(err)
		return
	}
	me, ok := tree.(*ast.ModuloExpression)
	if !ok {
		t.Errorf("expect *ast.ModuloExpression but %T", tree)
		return
	}
	testIntLiteralNode(t, me.Left(), 111)
	testIntLiteralNode(t, me.Right(), 222)
}
