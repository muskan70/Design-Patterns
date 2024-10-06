package main

import (
	"fmt"
	"sync"
)

type TV struct {
	status     TVState
	curChannel int
	curVolume  int
}

var tv *TV
var lock = &sync.Mutex{}

func getTVInstance() *TV {
	if tv == nil {
		lock.Lock()
		defer lock.Unlock()
		if tv == nil {
			tv = &TV{
				curChannel: 0,
				curVolume:  0,
			}
			tv.status = &TVOffState{tv: tv}
		}
	}
	return tv
}

func (t *TV) pressPowerButton(state TVState) {
	t.status = state
	t.curChannel = 0
	t.curVolume = 0
}

func (t *TV) setChannel(ch int) {
	fmt.Println("Setting TV channel No:", ch)
	t.curChannel = ch
}

func (t *TV) setVolume(v int) {
	fmt.Println("Setting TV Volume:", v)
	t.curVolume = v
}

func (t *TV) getStatus() TVState {
	return t.status
}

func (t *TV) getChannel() int {
	return t.curChannel
}

func (t *TV) getVolume() int {
	return t.curVolume
}
