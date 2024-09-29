package command

import receiver "comm/Receiver"

type OffCommand struct {
	Device receiver.Device
}

func (o *OffCommand) Execute() {
	o.Device.Off()
}
