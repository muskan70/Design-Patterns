package receiver

import "fmt"

type TV struct {
	turnOn    bool
	channelNo int
}

func (t *TV) On() {
	fmt.Println("Turning TV On")
	t.turnOn = true
}

func (t *TV) Off() {
	fmt.Println("Turning TV Off")
	t.turnOn = false
}

func (t *TV) SetChannelNo(channelNo int) {
	fmt.Println("Setting TV ChannelNo:", channelNo)
	t.channelNo = channelNo
}

func (t *TV) SetTemperature(temp int) {
	fmt.Println("Invalid command")
}
