package ast

import (
	"fmt"
	"io"
)

type binaryExpression struct {
	pos
	left  Node
	right Node
}

func (n *binaryExpression) initialize(line int, column int, left Node, right Node) {
	n.line = line
	n.column = column
	n.left = left
	n.right = right
}

func (n *binaryExpression) Left() Node {
	return n.left
}

func (n *binaryExpression) Right() Node {
	return n.right
}

func (n *binaryExpression) dump(f io.Writer, nNesting int) {
	indent(f, nNesting)
	fmt.Fprintf(f, "%T(%d:%d):\n", n, n.Line(), n.Column())
	n.Left().dump(f, nNesting+1)
	n.Right().dump(f, nNesting+1)
}

type AdditionExpression struct {
	binaryExpression
}

func NewAdditionExpression(line int, column int, left Node, right Node) (ret *AdditionExpression) {
	ret = &AdditionExpression{}
	ret.binaryExpression.initialize(line, column, left, right)
	return
}

type SubtractionExpression struct {
	binaryExpression
}

func NewSubtractionExpression(line int, column int, left Node, right Node) (ret *SubtractionExpression) {
	ret = &SubtractionExpression{}
	ret.binaryExpression.initialize(line, column, left, right)
	return
}

type MultiplicationExpression struct {
	binaryExpression
}

func NewMultiplicationExpression(line int, column int, left Node, right Node) (ret *MultiplicationExpression) {
	ret = &MultiplicationExpression{}
	ret.binaryExpression.initialize(line, column, left, right)
	return
}

type DivisionExpression struct {
	binaryExpression
}

func NewDivisionExpression(line int, column int, left Node, right Node) (ret *DivisionExpression) {
	ret = &DivisionExpression{}
	ret.binaryExpression.initialize(line, column, left, right)
	return
}

type ModuloExpression struct {
	binaryExpression
}

func NewModuloExpression(line int, column int, left Node, right Node) (ret *ModuloExpression) {
	ret = &ModuloExpression{}
	ret.binaryExpression.initialize(line, column, left, right)
	return
}
