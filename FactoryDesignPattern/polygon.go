package main

import "fmt"

type IShape interface {
	GetName() string
	GetFeatures()
	GetParameter() float64
}

type Polygon struct {
	sides    int
	edgeType string
	sideLen  []int
}

func (p *Polygon) GetName() string {
	return "POLYGON"
}

func (p *Polygon) GetFeatures() {
	fmt.Println("It is a polygon of", p.sides)
}

func (p *Polygon) GetParameter() float64 {
	parameter := 0.0
	for i := range p.sideLen {
		parameter += float64(p.sideLen[i])
	}
	return parameter
}

type Triangle struct {
	Polygon
}

func (t *Triangle) GetName() string {
	return "TRIANGLE"
}

func (t *Triangle) GetFeatures() {
	switch t.edgeType {
	case EQUAL_ALL:
		fmt.Println("It is an equiltataral triangle with each angle of 60 degree + ", EQUAL_ALL)
	case EQUAL_TWO_SIDES:
		fmt.Println("It is an isoseles triangle with two equal angles + ", EQUAL_TWO_SIDES)
	case EQUAL_NOT:
		fmt.Println("It is a scalene triangle + ", EQUAL_NOT)
	case ANGLE_90:
		fmt.Println("It is a right angle triangle")
	}
}

func (t *Triangle) GetParameter() float64 {
	switch t.edgeType {
	case EQUAL_ALL:
		return 3 * float64(t.sideLen[0])
	case EQUAL_TWO_SIDES:
		return float64(2*t.sideLen[0] + t.sideLen[1])
	case EQUAL_NOT:
		return float64(t.sideLen[0] + t.sideLen[1] + t.sideLen[2])
	case ANGLE_90:
		return float64(t.sideLen[0] + t.sideLen[1] + t.sideLen[2])
	}
	return 0
}

type Quadrilateral struct {
	Polygon
}

func (q *Quadrilateral) GetName() string {
	return "QUADILATERAL"
}

func (q *Quadrilateral) GetFeatures() {
	switch q.edgeType {
	case EQUAL_ALL_Angle_90:
		fmt.Println("It is a square + ", EQUAL_ALL_Angle_90)
	case EQUAL_TWO_SIDES_ANGLE_90:
		fmt.Println("It is a rectangle + ", EQUAL_TWO_SIDES_ANGLE_90)
	case EQUAL_TWO_SIDES:
		fmt.Println("It is a parallelogram + ", EQUAL_TWO_SIDES)
	case EQUAL_ALL:
		fmt.Println("It is a rhombus + ", EQUAL_ALL)
	case EQUAL_NOT:
		fmt.Println("It is a normal quadilateral + ", EQUAL_NOT)

	}
}

func (q *Quadrilateral) GetParameter() float64 {
	switch q.edgeType {
	case EQUAL_ALL_Angle_90:
		return 4 * float64(q.sideLen[0])
	case EQUAL_TWO_SIDES_ANGLE_90:
		return float64(2*q.sideLen[0] + 2*q.sideLen[1])
	case EQUAL_TWO_SIDES:
		return float64(2*q.sideLen[0] + 2*q.sideLen[1])
	case EQUAL_ALL:
		return float64(4 * q.sideLen[0])
	case EQUAL_NOT:
		return float64(q.sideLen[0] + q.sideLen[1] + q.sideLen[2] + q.sideLen[3])
	}
	return 0
}

type Pentagon struct {
	Polygon
}

func (P *Pentagon) GetName() string {
	return "PENTAGON"
}

func (p *Pentagon) GetFeatures() {
	fmt.Println("It is a pentagon of", p.sides)
}

func (p *Pentagon) GetParameter() float64 {
	return 5 * float64(p.sideLen[0])
}

type Hexagon struct {
	Polygon
}

func (h *Hexagon) GetName() string {
	return "HEXAGON"
}

func (h *Hexagon) GetFeatures() {
	fmt.Println("It is a hexagon of", h.sides)
}

func (h *Hexagon) GetParameter() float64 {
	return 6 * float64(h.sideLen[0])
}

type Octagon struct {
	Polygon
}

func (o *Octagon) GetName() string {
	return "OCTAGON"
}

func (o *Octagon) GetFeatures() {
	fmt.Println("It is a octagon of", o.sides)
}

func (o *Octagon) GetParameter() float64 {
	return 8 * float64(o.sideLen[0])
}

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
