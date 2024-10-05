package receiver

type Device interface {
	On()
	Off()
	SetChannelNo(channelNo int)
	SetTemperature(temp int)
}
