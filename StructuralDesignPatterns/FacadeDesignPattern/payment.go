package main

import (
	"errors"
	"fmt"
)

type IPayment interface {
	MakePayment(amount float64) error
}

type CreditCard struct {
	CardNo      int
	CreditLimit float64
}

func (c *CreditCard) MakePayment(amount float64) error {
	if c.CreditLimit < amount {
		return errors.New("insufficient credit limit")
	}
	c.CreditLimit -= amount
	fmt.Println("Payment done through credit card")
	return nil
}

type Wallet struct {
	Balance float64
}

func (w *Wallet) MakePayment(amount float64) error {
	if w.Balance < amount {
		return errors.New("insufficient wallet balance")
	}
	w.Balance -= amount
	fmt.Println("Payment done through wallet")
	return nil
}
