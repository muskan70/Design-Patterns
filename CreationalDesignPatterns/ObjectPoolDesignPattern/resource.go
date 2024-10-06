package main

var x = 0

type Resource struct {
	ResourceId int
}

func GetResourceInstance() *Resource {
	x++
	return &Resource{ResourceId: x}
}
