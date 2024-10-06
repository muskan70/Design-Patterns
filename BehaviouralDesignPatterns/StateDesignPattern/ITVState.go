package main

import "fmt"

type TVState interface {
	turnOnTV()
	turnOffTV()
	setChannelNo(ch int)
	setVolume(v int)
}

type TVOffState struct {
	tv *TV
}

func (s *TVOffState) turnOnTV() {
	fmt.Println("Turning TV ON!")
	s.tv.pressPowerButton(&TVOnState{tv: s.tv})
}
func (s *TVOffState) turnOffTV() {
	fmt.Println("TV is already off")
}
func (s *TVOffState) setChannelNo(ch int) {
	fmt.Println("TV is off, first Turn On TV")
}
func (s *TVOffState) setVolume(v int) {
	fmt.Println("TV is off, first Turn On TV")
}

type TVOnState struct {
	tv *TV
}

func (s *TVOnState) turnOnTV() {
	fmt.Println("TV is already on")
}
func (s *TVOnState) turnOffTV() {
	fmt.Println("Turning TV Off!")
	s.tv.pressPowerButton(&TVOffState{tv: s.tv})
}
func (s *TVOnState) setChannelNo(ch int) {
	s.tv.setChannel(ch)
}
func (s *TVOnState) setVolume(v int) {
	s.tv.setVolume(v)
}
