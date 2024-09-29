package command

import receiver "comm/Receiver"

type SetChannelCommand struct {
	Device    receiver.Device
	ChannelNo int
}

func (c *SetChannelCommand) Execute() {
	c.Device.SetChannelNo(c.ChannelNo)
}
