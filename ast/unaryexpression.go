package ast

import (
	"fmt"
	"io"
)

type unaryExpression struct {
	pos
	operand Node
}

func (n *unaryExpression) initialize(line int, column int, operand Node) {
	n.line = line
	n.column = column
	n.operand = operand
}

func (n *unaryExpression) Operand() Node {
	return n.operand
}

func (n *unaryExpression) dump(f io.Writer, nNesting int) {
	indent(f, nNesting)
	fmt.Fprintf(f, "%T(%d:%d):\n", n, n.Line(), n.Column())
	n.Operand().dump(f, nNesting+1)
}

type MinusExpression struct {
	unaryExpression
}

func NewMinusExpression(line int, column int, operand Node) (ret *MinusExpression) {
	ret = &MinusExpression{}
	ret.unaryExpression.initialize(line, column, operand)
	return
}

type NotExpression struct {
	unaryExpression
}

func NewNotExpression(line int, column int, operand Node) (ret *NotExpression) {
	ret = &NotExpression{}
	ret.unaryExpression.initialize(line, column, operand)
	return
}
