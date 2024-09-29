package main

import (
	command "comm/Command"
	invoker "comm/Invoker"
	receiver "comm/Receiver"
)

func main() {
	remote := invoker.RemoteControl{}

	tv := &receiver.TV{}

	onTvCommand := command.OnCommand{Device: tv}
	remote.SetCommand(&onTvCommand)
	remote.PressButton()

	setChannelCommand1 := command.SetChannelCommand{Device: tv, ChannelNo: 42}
	remote.SetCommand(&setChannelCommand1)
	remote.PressButton()

	setChannelCommand2 := command.SetChannelCommand{Device: tv, ChannelNo: 105}
	remote.SetCommand(&setChannelCommand2)
	remote.PressButton()

	remote.Undo()

	offTvCommand := command.OffCommand{Device: tv}
	remote.SetCommand(&offTvCommand)
	remote.PressButton()

	ac := &receiver.AirConditioner{}

	onAcCommand := command.OnCommand{Device: ac}
	remote.SetCommand(&onAcCommand)
	remote.PressButton()

	setTemperatureCommand1 := command.SetTemperatureCommand{Device: ac, Temperature: 24}
	remote.SetCommand(&setTemperatureCommand1)
	remote.PressButton()

	setTemperatureCommand2 := command.SetTemperatureCommand{Device: ac, Temperature: 27}
	remote.SetCommand(&setTemperatureCommand2)
	remote.PressButton()

	remote.Undo()

	offAcCommand := command.OffCommand{Device: ac}
	remote.SetCommand(&offAcCommand)
	remote.PressButton()
}
