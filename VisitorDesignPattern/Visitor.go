package main

import "math"

type IVisitor interface {
	VisitForCircle(c *Circle)
	VisitForSquare(s *Square)
	VisitForRectangle(r *Rectangle)
}

type AreaCalculator struct {
}

func (a *AreaCalculator) VisitForCircle(c *Circle) {
	c.Area = math.Floor((3.14*float64(c.Radius)*float64(c.Radius))*100) / 100
}

func (a *AreaCalculator) VisitForSquare(s *Square) {
	s.Area = float64(s.Side) * float64(s.Side)
}

func (a *AreaCalculator) VisitForRectangle(r *Rectangle) {
	r.Area = float64(r.Length) * float64(r.Breadth)
}

type ParameterCalculator struct {
}

func (p *ParameterCalculator) VisitForCircle(c *Circle) {
	c.Parameter = 2 * 3.14 * float64(c.Radius)
}

func (p *ParameterCalculator) VisitForSquare(s *Square) {
	s.Parameter = float64(s.Side) * 4
}

func (p *ParameterCalculator) VisitForRectangle(r *Rectangle) {
	r.Parameter = (float64(r.Length) + float64(r.Breadth)) * 2
}
