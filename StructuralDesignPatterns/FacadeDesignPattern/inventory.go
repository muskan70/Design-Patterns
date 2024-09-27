package main

import (
	"errors"
	"fmt"
)

type Product struct {
	ProductId int
	Name      string
	Price     float64
	Quantity  int
}

var inventory map[int]Product
var i int

func NewInventory() {
	inventory = make(map[int]Product)
	i = 0
}

func AddItemInInventory(pname string, price float64, qty int) {
	i++
	inventory[i] = Product{
		ProductId: i,
		Name:      pname,
		Price:     price,
		Quantity:  qty,
	}
}

func GetProduct(itemId int, qty int) (*Product, error) {
	if inventory[itemId].Quantity > qty {
		p := inventory[itemId]
		p.Quantity -= qty
		inventory[itemId] = p
		fmt.Println("Product added in order basket:", p.Name)
		return &Product{
			ProductId: itemId, Name: p.Name, Price: p.Price, Quantity: qty,
		}, nil
	}
	return nil, errors.New("Product not available")
}
