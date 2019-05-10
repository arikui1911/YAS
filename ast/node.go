package ast

type Node interface{
	Line() int
	Column() int
}

type intLiteral struct {
	line int
	column int
	value int
}

func NewIntLiteral(line int, column int, v int) Node {
	return &intLiteral{
		line: line,
		column: column,
		value: v,
	}
}

func (n *intLiteral) Line() int { return n.line }

func (n *intLiteral) Column() int { return n.column }

func (n *intLiteral) Value() int { return n.value }
