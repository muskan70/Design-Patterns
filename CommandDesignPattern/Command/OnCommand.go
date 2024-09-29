package command

import receiver "comm/Receiver"

type OnCommand struct {
	Device receiver.Device
}

func (o *OnCommand) Execute() {
	o.Device.On()
}
