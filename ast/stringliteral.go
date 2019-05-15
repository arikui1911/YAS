package ast

import (
	"fmt"
	"io"
)

type StringLiteral struct {
	pos
	value string
}

func NewStringLiteral(line int, column int, v string) *StringLiteral {
	return &StringLiteral{
		pos: pos{
			line:   line,
			column: column,
		},
		value: v,
	}
}

func (n *StringLiteral) Value() string { return n.value }

func (n *StringLiteral) dump(f io.Writer, nNesting int) {
	indent(f, nNesting)
	fmt.Fprintf(f, "%T(%d:%d): %v\n", n, n.Line(), n.Column(), n.Value())
}
