package ast

import (
	"fmt"
	"io"
)

type Node interface {
	Line() int
	Column() int
	dump(f io.Writer, nNesting int)
}

func Dump(output io.Writer, tree Node) {
	if tree == nil {
		return
	}
	tree.dump(output, 0)
}

func indent(f io.Writer, n int) {
	for i := 0; i < n; i++ {
		fmt.Fprint(f, "  ")
	}
}

type pos struct {
	line   int
	column int
}

func (p *pos) Line() int { return p.line }

func (p *pos) Column() int { return p.column }

type intLiteral struct {
	pos
	value int
}

func NewIntLiteral(line int, column int, v int) Node {
	return &intLiteral{
		pos: pos{
			line:   line,
			column: column,
		},
		value: v,
	}
}

func (n *intLiteral) Value() int { return n.value }

func (n *intLiteral) dump(f io.Writer, nNesting int) {
	indent(f, nNesting)
	fmt.Fprintf(f, "%T(%d:%d): %v\n", n, n.Line(), n.Column(), n.Value())
}

type floatLiteral struct {
	pos
	value float64
}

func NewFloatLiteral(line int, column int, v float64) Node {
	return &floatLiteral{
		pos: pos{
			line:   line,
			column: column,
		},
		value: v,
	}
}

func (n *floatLiteral) Value() float64 { return n.value }

func (n *floatLiteral) dump(f io.Writer, nNesting int) {
	indent(f, nNesting)
	fmt.Fprintf(f, "%T(%d:%d): %v\n", n, n.Line(), n.Column(), n.Value())
}

type stringLiteral struct {
	pos
	value string
}

func NewStringLiteral(line int, column int, v string) Node {
	return &stringLiteral{
		pos: pos{
			line:   line,
			column: column,
		},
		value: v,
	}
}

func (n *stringLiteral) Value() string { return n.value }

func (n *stringLiteral) dump(f io.Writer, nNesting int) {
	indent(f, nNesting)
	fmt.Fprintf(f, "%T(%d:%d): %v\n", n, n.Line(), n.Column(), n.Value())
}
