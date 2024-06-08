package main

import "fmt"

type Circle struct {
	radius int
}

func (c *Circle) GetName() string {
	return "CIRCLE"
}

func (c *Circle) GetFeatures() {
	fmt.Println("It is a circle of radius:", c.radius)
}

func (c *Circle) GetParameter() float64 {
	return 6.28 * float64(c.radius)
}
