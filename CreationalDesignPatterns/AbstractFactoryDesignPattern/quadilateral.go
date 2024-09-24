package main

import "fmt"

type Quadrilateral struct {
	Shape
}

func (q *Quadrilateral) GetName() string {
	return "QUADILATERAL"
}

func (q *Quadrilateral) GetFeatures() {
	fmt.Println("It is a normal quadilateral + ", EQUAL_NOT)
}

func (q *Quadrilateral) GetParameter() float64 {
	return float64(q.edgeLen[0] + q.edgeLen[1] + q.edgeLen[2] + q.edgeLen[3])
}

type Square struct {
	Shape
}

func (q *Square) GetName() string {
	return "SQUARE"
}

func (q *Square) GetFeatures() {
	fmt.Println("It is a square + ", EQUAL_ALL_Angle_90)
}

func (q *Square) GetParameter() float64 {
	return 4 * float64(q.edgeLen[0])
}

type Rhombus struct {
	Shape
}

func (q *Rhombus) GetName() string {
	return "RHOMBUS"
}

func (q *Rhombus) GetFeatures() {
	fmt.Println("It is a rhombus + ", EQUAL_ALL)
}

func (q *Rhombus) GetParameter() float64 {
	return 4 * float64(q.edgeLen[0])
}

type Rectangle struct {
	Shape
}

func (q *Rectangle) GetName() string {
	return "RECTANGLE"
}

func (q *Rectangle) GetFeatures() {
	fmt.Println("It is a rectangle + ", EQUAL_TWO_SIDES_ANGLE_90)
}

func (q *Rectangle) GetParameter() float64 {
	return float64(2*q.edgeLen[0] + 2*q.edgeLen[1])
}

type Parallelogram struct {
	Shape
}

func (q *Parallelogram) GetName() string {
	return "PARALLELOGRAM"
}

func (q *Parallelogram) GetFeatures() {
	fmt.Println("It is a parallelogram + ", EQUAL_TWO_SIDES)
}

func (q *Parallelogram) GetParameter() float64 {
	return float64(2*q.edgeLen[0] + 2*q.edgeLen[1])
}
