package main

import "time"

type Patient struct {
	name             string
	lastVisited      time.Time
	registrationDone bool
	doctorExamined   bool
	examinationDone  bool
	paymentDone      bool
}

func NewPatient(name string) Patient {
	return Patient{name: name, lastVisited: time.Now(), registrationDone: false, doctorExamined: false, examinationDone: false, paymentDone: false}
}
