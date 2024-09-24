package main

type IShapeFactory interface {
	GetShape(edgeType string, edgeLen []int) IShape
}

func GetShapeFactory(sides int) IShapeFactory {
	if sides == 3 {
		return &TriangleFactory{}
	} else if sides == 4 {
		return &QuadrilateralFactory{}
	} else {
		return nil
	}
}
