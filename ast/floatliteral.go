package ast

import (
	"fmt"
	"io"
)

type FloatLiteral struct {
	pos
	value float64
}

func NewFloatLiteral(line int, column int, v float64) *FloatLiteral {
	return &FloatLiteral{
		pos: pos{
			line:   line,
			column: column,
		},
		value: v,
	}
}

func (n *FloatLiteral) Value() float64 { return n.value }

func (n *FloatLiteral) dump(f io.Writer, nNesting int) {
	indent(f, nNesting)
	fmt.Fprintf(f, "%T(%d:%d): %v\n", n, n.Line(), n.Column(), n.Value())
}
