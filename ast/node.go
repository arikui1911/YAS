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

type VarRef struct {
	pos
	identifier string
}

func NewVarRef(line int, column int, identifier string) *VarRef {
	return &VarRef{
		pos: pos{
			line:   line,
			column: column,
		},
		identifier: identifier,
	}
}

func (n *VarRef) Identifier() string { return n.identifier }

func (n *VarRef) dump(f io.Writer, nNesting int) {
	indent(f, nNesting)
	fmt.Fprintf(f, "%T(%d:%d): %v\n", n, n.Line(), n.Column(), n.Identifier())
}

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

type FloatLiteral struct {
	pos
	value float64
}

func NewFloatLiteral(line int, column int, v float64) Node {
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

type StringLiteral struct {
	pos
	value string
}

func NewStringLiteral(line int, column int, v string) Node {
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
