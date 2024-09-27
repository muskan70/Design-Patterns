package main

type IPizza interface {
	GetCost() int
}

type MargheritaPizza struct {
}

func (mp *MargheritaPizza) GetCost() int {
	return 100
}

type VeggieDelightPizza struct {
}

func (vg *VeggieDelightPizza) GetCost() int {
	return 150
}
