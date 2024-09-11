package main

type ArithmeticExpression interface {
	evaluate() float64
}

type Number struct {
	value int
}

func NewNumber(val int) *Number {
	return &Number{value: val}
}

func (n *Number) evaluate() float64 {
	return float64(n.value)
}

type Expression struct {
	leftExpression  ArithmeticExpression
	rightExpression ArithmeticExpression
	operation       string
}

func NewExpression(left ArithmeticExpression, right ArithmeticExpression, op string) *Expression {
	return &Expression{
		leftExpression:  left,
		rightExpression: right,
		operation:       op,
	}
}

func (e *Expression) evaluate() float64 {
	switch e.operation {
	case "+":
		return e.leftExpression.evaluate() + e.rightExpression.evaluate()
	case "-":
		return e.leftExpression.evaluate() - e.rightExpression.evaluate()
	case "*":
		return e.leftExpression.evaluate() * e.rightExpression.evaluate()
	case "/":
		return e.leftExpression.evaluate() / e.rightExpression.evaluate()
	}
	return 0.0
}
