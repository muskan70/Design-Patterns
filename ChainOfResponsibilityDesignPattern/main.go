package main

func main() {
	c := &Cashier{}

	e := &MedicalExamination{}
	e.SetNext(c)

	d := &Doctor{}
	d.SetNext(e)

	r := Reception{}
	r.SetNext(d)

	p := NewPatient("Naman")
	r.Execute(&p)

}
