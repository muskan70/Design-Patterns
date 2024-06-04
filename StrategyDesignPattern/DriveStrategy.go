package main

import "fmt"

type DriveStrategy interface {
	Drive()
}

type NormalDriveStrategy struct{}

func (nd *NormalDriveStrategy) Drive() {
	fmt.Println("Normal Drive Strategy")
}

type SpecialDriveStrategy struct{}

func (sd *SpecialDriveStrategy) Drive() {
	fmt.Println("Special Drive Strategy")
}
