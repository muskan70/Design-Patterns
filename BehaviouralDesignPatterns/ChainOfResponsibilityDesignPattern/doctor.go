package main

import "fmt"

type Doctor struct {
	next Department
}

func (d *Doctor) Execute(p *Patient) {
	if p.doctorExamined {
		fmt.Println("Patient already visited the doctor")
		d.next.Execute(p)
		return
	}
	fmt.Println("Patient visiting the doctor")
	p.doctorExamined = true
	d.next.Execute(p)
}

func (d *Doctor) SetNext(next Department) {
	d.next = next
}
