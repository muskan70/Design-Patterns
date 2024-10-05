package main

import "fmt"

type MedicalExamination struct {
	next Department
}

func (m *MedicalExamination) Execute(p *Patient) {
	if p.examinationDone {
		fmt.Println("Patient already examined")
		m.next.Execute(p)
		return
	}
	fmt.Println("Patient's medical examination going on.")
	p.examinationDone = true
	m.next.Execute(p)
}

func (m *MedicalExamination) SetNext(next Department) {
	m.next = next
}
