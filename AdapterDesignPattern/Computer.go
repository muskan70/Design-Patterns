package main

import "fmt"

type Computer interface {
	InsertIntoLighteningPort()
}

type Mac struct{}

func (m *Mac) InsertIntoLighteningPort() {
	fmt.Println("Lightening port is plugged to mac machine")
}

type Windows struct{}

func (w *Windows) InsertIntoUSBPort() {
	fmt.Println("USB port is plugged to windows machine")
}
