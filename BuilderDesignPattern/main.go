package main

import "builder/houseBuilder"

func main() {
	woodenHouseBuilder := houseBuilder.GetHouseBuilder("wooden")

	director := houseBuilder.NewDirector(woodenHouseBuilder)
	h1 := director.MakeWoodenHouse()

	iglooHouseBuilder := houseBuilder.GetHouseBuilder("igloo")
	director.SetBuilder(iglooHouseBuilder)
	h2 := director.MakeIglooHouse()

	normalHouseBuilder := houseBuilder.GetHouseBuilder("normal")
	director.SetBuilder(normalHouseBuilder)
	h3 := director.MakeNormalHouse()

	h1.GetHouseDetails()
	h2.GetHouseDetails()
	h3.GetHouseDetails()
}
