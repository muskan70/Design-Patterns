package main

type QuadrilateralFactory struct {
}

func (q *QuadrilateralFactory) GetShape(edgeType string, edgeLen []int) IShape {
	shape := Shape{
		sides:    4,
		edgeType: edgeType,
		edgeLen:  edgeLen,
	}
	switch edgeType {
	case EQUAL_ALL_Angle_90:
		return &Square{Shape: shape}
	case EQUAL_TWO_SIDES_ANGLE_90:
		return &Rectangle{Shape: shape}
	case EQUAL_TWO_SIDES:
		return &Parallelogram{Shape: shape}
	case EQUAL_ALL:
		return &Rhombus{Shape: shape}
	case EQUAL_NOT:
		return &Quadrilateral{Shape: shape}
	}
	return nil
}
