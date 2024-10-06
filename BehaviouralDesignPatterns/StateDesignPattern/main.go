package main

import "fmt"

func main() {
	tvInst := getTVInstance()
	state := tvInst.getStatus()

	state.turnOnTV()
	state = tv.getStatus()
	state.setChannelNo(45)
	fmt.Println("ChannelNo:", tv.getChannel())
	state.setVolume(70)
	fmt.Println("Volume:", tv.getVolume())

	state.setChannelNo(32)
	fmt.Println("ChannelNo:", tv.getChannel())
	state.setVolume(83)
	fmt.Println("Volume:", tv.getVolume())

	state.turnOffTV()
	state = tv.getStatus()
	state.setChannelNo(45)
	fmt.Println("ChannelNo:", tv.getChannel())
	state.setVolume(70)
	fmt.Println("Volume:", tv.getVolume())

}
