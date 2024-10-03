package main

func main() {
	stationManager := newStationManger()

	passengerTrain := NewPassengerTrain(stationManager)
	freightTrain := NewFreightTrain(stationManager)

	passengerTrain.arrive()
	freightTrain.arrive()
	passengerTrain.depart()
}
