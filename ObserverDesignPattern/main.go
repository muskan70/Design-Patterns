package main

import (
	observable "observerPattern/Observable"
	"observerPattern/display"
)

func main() {
	ws := observable.NewWeatherStation()

	tv := display.NewTVDisplay()
	mobile := display.NewMobileDisplay()

	ws.AddDisplay(tv)
	ws.AddDisplay(mobile)

	ws.SetUpdates(45)
	ws.RemoveDisplay("tv")
	ws.SetUpdates(42)
}
