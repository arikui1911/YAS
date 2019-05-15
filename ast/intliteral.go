package ast

import (
	"fmt"
	"io"
)

type IntLiteral struct {
	pos
	value int
}

func NewIntLiteral(line int, column int, v int) Node {
	return &IntLiteral{
		pos: pos{
			line:   line,
			column: column,
		},
		value: v,
	}
}

func (n *IntLiteral) Value() int { return n.value }

func (n *IntLiteral) dump(f io.Writer, nNesting int) {
	indent(f, nNesting)
	fmt.Fprintf(f, "%T(%d:%d): %v\n", n, n.Line(), n.Column(), n.Value())
}
