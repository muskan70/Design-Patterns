package main

import "fmt"

type ScaleneTriangle struct {
	Shape
}

func (q *ScaleneTriangle) GetName() string {
	return "SCALENE TRIANGLE"
}

func (q *ScaleneTriangle) GetFeatures() {
	fmt.Println("It is a scalene triangle + ", EQUAL_NOT)
}

func (q *ScaleneTriangle) GetParameter() float64 {
	return float64(q.edgeLen[0] + q.edgeLen[1] + q.edgeLen[2])
}

type EquilateralTriangle struct {
	Shape
}

func (q *EquilateralTriangle) GetName() string {
	return "EQUILATERAL TRIANGLE"
}

func (q *EquilateralTriangle) GetFeatures() {
	fmt.Println("It is an equiltataral triangle with each angle of 60 degree + ", EQUAL_ALL)
}

func (q *EquilateralTriangle) GetParameter() float64 {
	return 3 * float64(q.edgeLen[0])
}

type IsoselesTriangle struct {
	Shape
}

func (q *IsoselesTriangle) GetName() string {
	return "ISOSELES TRIANGLE"
}

func (q *IsoselesTriangle) GetFeatures() {
	fmt.Println("It is an isoseles triangle with two equal angles + ", EQUAL_TWO_SIDES)
}

func (q *IsoselesTriangle) GetParameter() float64 {
	return float64(2*q.edgeLen[0] + q.edgeLen[1])
}

type RightAngleTriangle struct {
	Shape
}

func (q *RightAngleTriangle) GetName() string {
	return "RIGHT ANGLE TRIANGLE"
}

func (q *RightAngleTriangle) GetFeatures() {
	fmt.Println("It is a right angle triangle")
}

func (q *RightAngleTriangle) GetParameter() float64 {
	return float64(q.edgeLen[0] + q.edgeLen[1] + q.edgeLen[2])
}
