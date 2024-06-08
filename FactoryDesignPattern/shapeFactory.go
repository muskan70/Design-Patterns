package main

func shapeFactory(sides int, edgeType string, sideLen []int) IShape {
	if sides == 0 {
		return &Circle{radius: sideLen[0]}
	} else if sides == 3 {
		return &Triangle{Polygon{sides: sides, edgeType: edgeType, sideLen: sideLen}}
	} else if sides == 4 {
		return &Quadrilateral{Polygon{sides: sides, edgeType: edgeType, sideLen: sideLen}}
	} else if sides == 5 {
		return &Pentagon{Polygon{sides: sides, edgeType: edgeType, sideLen: sideLen}}
	} else if sides == 6 {
		return &Hexagon{Polygon{sides: sides, edgeType: edgeType, sideLen: sideLen}}
	} else if sides == 8 {
		return &Octagon{Polygon{sides: sides, edgeType: edgeType, sideLen: sideLen}}
	} else {
		return &Polygon{sides: sides, edgeType: edgeType, sideLen: sideLen}
	}
}
