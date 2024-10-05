package main

type Expression interface {
	Interpret(c *Context) int
}

type NumberExpression struct {
	numStr string
}

func (n *NumberExpression) Interpret(c *Context) int {
	return c.get(n.numStr)
}

type BinaryExpression struct {
	left     Expression
	right    Expression
	operator string
}

func (b *BinaryExpression) Interpret(c *Context) int {
	switch b.operator {
	case "+":
		return b.left.Interpret(c) + b.right.Interpret(c)
	case "-":
		return b.left.Interpret(c) - b.right.Interpret(c)
	case "*":
		return b.left.Interpret(c) * b.right.Interpret(c)
	}
	return 0
}
