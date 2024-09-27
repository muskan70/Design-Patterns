package main

import "fmt"

func main() {
	num1 := NewNumber(5)
	num2 := NewNumber(7)
	exp1 := NewExpression(num1, num2, "*")

	num3 := NewNumber(7)
	parentExp := NewExpression(num3, exp1, "-")
	ans := parentExp.evaluate()
	fmt.Println(ans)
}
