package receiver

import "fmt"

type AirConditioner struct {
	turnOn      bool
	temperature int
}

func (a *AirConditioner) On() {
	fmt.Println("Turning AC On")
	a.turnOn = true
}

func (a *AirConditioner) Off() {
	fmt.Println("Turning AC Off")
	a.turnOn = false
}

func (a *AirConditioner) SetTemperature(temp int) {
	fmt.Println("Setting AC Temperature:", temp)
	a.temperature = temp
}

func (a *AirConditioner) SetChannelNo(channelNo int) {
	fmt.Println("Invalid command")
}
