package main

import "fmt"

func main() {

	cntxt := &Context{hashmap: make(map[string]int)}
	cntxt.put("a", 4)
	cntxt.put("b", 7)

	exp1 := BinaryExpression{left: &NumberExpression{numStr: "a"}, right: &NumberExpression{numStr: "b"}, operator: "*"}
	fmt.Println(exp1.Interpret(cntxt))

	exp2 := BinaryExpression{left: &NumberExpression{numStr: "a"}, right: &NumberExpression{numStr: "b"}, operator: "+"}
	fmt.Println(exp2.Interpret(cntxt))
}
