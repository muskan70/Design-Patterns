package houseBuilder

type Director struct {
	builder IHouseBuilder
}

func NewDirector(b IHouseBuilder) *Director {
	return &Director{builder: b}
}

func (d *Director) SetBuilder(b IHouseBuilder) {
	d.builder = b
}

func (d *Director) MakeWoodenHouse() *House {
	d.builder.setNoofRooms(2)
	d.builder.setWalls()
	d.builder.setRoof()
	d.builder.setFloor()
	d.builder.setDoor()
	return d.builder.buildHouse()
}

func (d *Director) MakeIglooHouse() *House {
	d.builder.setNoofRooms(1)
	d.builder.setWalls()
	d.builder.setRoof()
	d.builder.setFloor()
	d.builder.setDoor()
	return d.builder.buildHouse()
}

func (d *Director) MakeNormalHouse() *House {
	d.builder.setNoofRooms(3)
	d.builder.setWalls()
	d.builder.setRoof()
	d.builder.setFloor()
	d.builder.setDoor()
	return d.builder.buildHouse()
}
