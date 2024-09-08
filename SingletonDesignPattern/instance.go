package main

import (
	"fmt"
	"sync"
)

type TV struct {
	status     string
	curChannel int
	curVolume  int
}

var tv *TV

var lock = &sync.Mutex{}

func getTvInstance() *TV {
	if tv == nil {
		lock.Lock()
		defer lock.Unlock()
		if tv == nil {
			fmt.Println("creating Instance")
			tv = &TV{status: "OFF", curChannel: 1, curVolume: 1}
		} else {
			fmt.Println("Tv Instance already present")
		}
	} else {
		fmt.Println("Tv Instance already present")
	}
	return tv
}
