package houseBuilder

type NormalHouseBuilder struct {
	houseType string
	rooms     int
	wallType  string
	doorType  string
	roofType  string
	floorType string
}

func (n *NormalHouseBuilder) setNoofRooms(rooms int) {
	n.rooms = rooms
}
func (n *NormalHouseBuilder) setWalls() {
	n.wallType = "brick walls"
}
func (n *NormalHouseBuilder) setDoor() {
	n.doorType = "steel gate"
}
func (n *NormalHouseBuilder) setRoof() {
	n.roofType = "concrete roof"
}
func (n *NormalHouseBuilder) setFloor() {
	n.floorType = "marble floor"
}
func (n *NormalHouseBuilder) buildHouse() *House {
	return &House{
		houseType: n.houseType,
		rooms:     n.rooms,
		wallType:  n.wallType,
		doorType:  n.doorType,
		roofType:  n.roofType,
		floorType: n.floorType,
	}
}
