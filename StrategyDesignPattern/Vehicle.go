package main

type Vehicle struct {
	name      string
	speed     int
	fuel      int
	driveType DriveStrategy
}

func NewVehicle(d DriveStrategy) Vehicle {
	return Vehicle{driveType: d,
		speed: 10}
}

type SportyVehicle struct {
	Vehicle
}

func NewSportyVehicle(d DriveStrategy) SportyVehicle {
	return SportyVehicle{
		Vehicle{driveType: d,
			speed: 10,
			name:  "sports car",
			fuel:  100},
	}
}

type PassengerVehicle struct {
	Vehicle
	noofSeats int
}

func NewPassengerVehicle(d DriveStrategy) PassengerVehicle {
	return PassengerVehicle{
		Vehicle{driveType: d,
			speed: 10,
			name:  "passenger vehicle",
			fuel:  100},
		7,
	}
}

type OffRoadVehicle struct {
	Vehicle
}

func NewOffRoadVehicle(d DriveStrategy) OffRoadVehicle {
	return OffRoadVehicle{
		Vehicle{driveType: d,
			speed: 10,
			name:  "offRoad vehicle",
			fuel:  100},
	}
}
