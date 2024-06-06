package main

type TomatoTopping struct {
	pizza IPizza
}

func (t *TomatoTopping) GetCost() int {
	return t.pizza.GetCost() + 20
}

type CheezeTopping struct {
	pizza IPizza
}

func (c *CheezeTopping) GetCost() int {
	return c.pizza.GetCost() + 50
}
