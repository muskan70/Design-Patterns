package houseBuilder

import "fmt"

type House struct {
	houseType string
	rooms     int
	wallType  string
	doorType  string
	roofType  string
	floorType string
}

func (h *House) GetHouseDetails() {
	fmt.Println("this house is", h.houseType)
	fmt.Println("It has", h.rooms, "rooms, made of", h.wallType, ",", h.doorType, ",", h.roofType, ",", h.floorType)
}
