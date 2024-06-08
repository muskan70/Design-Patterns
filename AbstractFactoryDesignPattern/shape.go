package main

type IShape interface {
	GetName() string
	GetFeatures()
	GetParameter() float64
}

type Shape struct {
	sides    int
	edgeType string
	edgeLen  []int
}
