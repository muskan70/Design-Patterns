package houseBuilder

type IHouseBuilder interface {
	setNoofRooms(rooms int)
	setWalls()
	setDoor()
	setRoof()
	setFloor()
	buildHouse() *House
}

func GetHouseBuilder(houseType string) IHouseBuilder {
	if houseType == "wooden" {
		return &WoodenHouseBuilder{houseType: houseType}
	} else if houseType == "igloo" {
		return &IglooHouseBuilder{houseType: houseType}
	}
	return &NormalHouseBuilder{houseType: houseType}
}
