package command

import receiver "comm/Receiver"

type SetTemperatureCommand struct {
	Device      receiver.Device
	Temperature int
}

func (t *SetTemperatureCommand) Execute() {
	t.Device.SetTemperature(t.Temperature)
}
