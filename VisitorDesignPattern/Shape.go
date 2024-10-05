package main

type IShape interface {
	GetShape() string
	Accept(visitor IVisitor)
}

type Shape struct {
	Area      float64
	Parameter float64
}

type Circle struct {
	Radius int
	Shape
}

func (c *Circle) GetShape() string {
	return "circle"
}

func (c *Circle) Accept(visitor IVisitor) {
	visitor.VisitForCircle(c)
}

type Square struct {
	Side int
	Shape
}

func (s *Square) GetShape() string {
	return "square"
}

func (s *Square) Accept(visitor IVisitor) {
	visitor.VisitForSquare(s)
}

type Rectangle struct {
	Length  int
	Breadth int
	Shape
}

func (r *Rectangle) GetShape() string {
	return "rectangle"
}

func (r *Rectangle) Accept(visitor IVisitor) {
	visitor.VisitForRectangle(r)
}
