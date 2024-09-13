package houseBuilder

type WoodenHouseBuilder struct {
	houseType string
	rooms     int
	wallType  string
	doorType  string
	roofType  string
	floorType string
}

func (w *WoodenHouseBuilder) setNoofRooms(rooms int) {
	w.rooms = rooms
}
func (w *WoodenHouseBuilder) setWalls() {
	w.wallType = "wooden walls"
}
func (w *WoodenHouseBuilder) setDoor() {
	w.doorType = "wooden door"
}
func (w *WoodenHouseBuilder) setRoof() {
	w.roofType = "wooden roof"
}
func (w *WoodenHouseBuilder) setFloor() {
	w.floorType = "wooden floor"
}
func (w *WoodenHouseBuilder) buildHouse() *House {
	return &House{
		houseType: w.houseType,
		rooms:     w.rooms,
		wallType:  w.wallType,
		doorType:  w.doorType,
		roofType:  w.roofType,
		floorType: w.floorType,
	}
}
