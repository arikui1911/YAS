package ast

import (
	"fmt"
	"io"
)

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
