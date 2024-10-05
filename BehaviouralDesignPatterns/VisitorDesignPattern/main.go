package main

import "fmt"

func main() {
	square := &Square{Side: 2}
	circle := &Circle{Radius: 3}
	rectangle := &Rectangle{Length: 2, Breadth: 3}

	areaCalculator := &AreaCalculator{}

	square.Accept(areaCalculator)
	circle.Accept(areaCalculator)
	rectangle.Accept(areaCalculator)

	parameterCalculator := &ParameterCalculator{}

	square.Accept(parameterCalculator)
	circle.Accept(parameterCalculator)
	rectangle.Accept(parameterCalculator)

	fmt.Println("Circle Params-> Area:", circle.Area, "Parameter:", circle.Parameter)
	fmt.Println("Square Params-> Area:", square.Area, "Parameter:", square.Parameter)
	fmt.Println("Rectangle Params-> Area:", rectangle.Area, "Parameter:", rectangle.Parameter)

}
