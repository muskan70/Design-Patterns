package main

func main() {
	v1 := NewSportyVehicle(&SpecialDriveStrategy{})
	v1.driveType.Drive()

	v2 := NewPassengerVehicle(&NormalDriveStrategy{})
	v2.driveType.Drive()

	v3 := NewOffRoadVehicle(&SpecialDriveStrategy{})
	v3.driveType.Drive()

}
