package main

import "fmt"

type Cashier struct {
	next Department
}

func (c *Cashier) Execute(p *Patient) {
	if p.paymentDone {
		fmt.Println("Patient already paid")
		return
	}
	fmt.Println("Patient paying the cashier")
	p.paymentDone = true
}

func (c *Cashier) SetNext(next Department) {
	c.next = next
}
