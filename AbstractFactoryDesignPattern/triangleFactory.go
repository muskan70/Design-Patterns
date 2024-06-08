package main

type TriangleFactory struct {
}

func (t *TriangleFactory) GetShape(edgeType string, edgeLen []int) IShape {
	shape := Shape{
		sides:    3,
		edgeType: edgeType,
		edgeLen:  edgeLen,
	}
	switch edgeType {
	case EQUAL_ALL:
		return &EquilateralTriangle{Shape: shape}
	case EQUAL_TWO_SIDES:
		return &IsoselesTriangle{Shape: shape}
	case ANGLE_90:
		return &RightAngleTriangle{Shape: shape}
	case EQUAL_NOT:
		return &ScaleneTriangle{Shape: shape}
	}
	return nil
}
