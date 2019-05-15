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
