package observable

import (
	"observerPattern/display"
)

type WeatherStation struct {
	displayList []display.IDisplay
	temperature int
}

func NewWeatherStation() IStation {
	return &WeatherStation{
		displayList: []display.IDisplay{},
		temperature: 37}
}

func (ws *WeatherStation) AddDisplay(d display.IDisplay) {
	ws.displayList = append(ws.displayList, d)
}

func (ws *WeatherStation) RemoveDisplay(displayName string) {
	i := 0
	for displayIdx := range ws.displayList {
		if ws.displayList[displayIdx].GetName() != displayName {
			ws.displayList[i] = ws.displayList[displayIdx]
			i++
		}
	}
	ws.displayList = ws.displayList[:i-1]
}

func (ws *WeatherStation) GetUpdates() int {
	return ws.temperature
}

func (ws *WeatherStation) SetUpdates(update int) {
	ws.temperature = update
	ws.NotifyDisplay()
}

func (ws *WeatherStation) NotifyDisplay() {
	for displayIdx := range ws.displayList {
		ws.displayList[displayIdx].Update(ws.temperature)
	}
}
