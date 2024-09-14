package main

import "fmt"

type WindowsAdaptor struct {
	WindowMachine *Windows
}

func (wa *WindowsAdaptor) InsertIntoLighteningPort() {
	fmt.Println("Lightening signal converted to USB by adapter")
	wa.WindowMachine.InsertIntoUSBPort()
}
