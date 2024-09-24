package houseBuilder

type IglooHouseBuilder struct {
	houseType string
	rooms     int
	wallType  string
	doorType  string
	roofType  string
	floorType string
}

func (i *IglooHouseBuilder) setNoofRooms(rooms int) {
	i.rooms = rooms
}
func (i *IglooHouseBuilder) setWalls() {
	i.wallType = "snow walls"
}
func (i *IglooHouseBuilder) setDoor() {
	i.doorType = "wooden door"
}
func (i *IglooHouseBuilder) setRoof() {
	i.roofType = "snow roof"
}
func (i *IglooHouseBuilder) setFloor() {
	i.floorType = "snow floor"
}
func (i *IglooHouseBuilder) buildHouse() *House {
	return &House{
		houseType: i.houseType,
		rooms:     i.rooms,
		wallType:  i.wallType,
		doorType:  i.doorType,
		roofType:  i.roofType,
		floorType: i.floorType,
	}
}
