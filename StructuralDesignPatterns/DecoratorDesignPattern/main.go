package main

import "fmt"

func main() {
	vg := &VeggieDelightPizza{}
	fmt.Println("VeggieDelight:", vg.GetCost())

	pizza1 := &TomatoTopping{pizza: vg}
	fmt.Println("VeggieDelight with Tomato Topping:", pizza1.GetCost())

	pizza3 := &CheezeTopping{pizza: pizza1}
	fmt.Println("VeggieDelight with Tomato+Cheeze Topping:", pizza3.GetCost())

	mg := &MargheritaPizza{}
	fmt.Println("Margherita:", mg.GetCost())

	pizza2 := &CheezeTopping{pizza: mg}
	fmt.Println("Margherita with Cheeze Topping:", pizza2.GetCost())

	pizza4 := &TomatoTopping{pizza: pizza2}
	fmt.Println("Margherita with Tomato+Cheeze Topping:", pizza4.GetCost())
}
