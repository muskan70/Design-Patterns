package main

import "fmt"

type Train interface {
	arrive()
	depart()
	permitArrival()
}

type PassengerTrain struct {
	havePermissionToArrive bool
	mediator               Mediator
}

func NewPassengerTrain(m Mediator) *PassengerTrain {
	return &PassengerTrain{havePermissionToArrive: false, mediator: m}
}

func (p *PassengerTrain) arrive() {
	if p.havePermissionToArrive || p.mediator.canArrive(p) {
		fmt.Println("Passenger Train Arrived")
		return
	}
	fmt.Println("Passenger Train Arrival blocked, waiting")
}

func (p *PassengerTrain) depart() {
	fmt.Println("Passenger Train Departed")
	p.havePermissionToArrive = false
	p.mediator.notifyAboutDeparture()
}

func (p *PassengerTrain) permitArrival() {
	fmt.Println("Passenger Train Arrival permitted, arriving")
	p.havePermissionToArrive = true
	p.arrive()
}

type FreightTrain struct {
	havePermissionToArrive bool
	mediator               Mediator
}

func NewFreightTrain(m Mediator) *FreightTrain {
	return &FreightTrain{havePermissionToArrive: false, mediator: m}
}

func (f *FreightTrain) arrive() {
	if f.havePermissionToArrive || f.mediator.canArrive(f) {
		fmt.Println("Freight Train Arrived")
		return
	}
	fmt.Println("Freight Train Arrival blocked, waiting")
}

func (f *FreightTrain) depart() {
	fmt.Println("Freight Train Departed")
	f.havePermissionToArrive = false
	f.mediator.notifyAboutDeparture()
}

func (f *FreightTrain) permitArrival() {
	fmt.Println("Freight Train Arrival permitted, arriving")
	f.havePermissionToArrive = true
	f.arrive()
}
