package main

import "fmt"

func main() {
	NewInventory()
	AddItemInInventory("Coke", 40, 5)
	AddItemInInventory("Lays", 20, 10)
	AddItemInInventory("Chocolate", 100, 4)

	walletPay := &Wallet{Balance: 500}
	emailNotify := &EmailNotification{UserName: "muskan", EmailId: "muskanmangla***@gmail.com"}
	order := NewOrderFacade(walletPay, emailNotify)

	if err := order.AddProduct(1, 2); err != nil {
		fmt.Println(err)
	}

	if err := order.AddProduct(2, 3); err != nil {
		fmt.Println(err)
	}

	if err := order.AddProduct(3, 1); err != nil {
		fmt.Println(err)
	}

	if err := order.CreateOrder(); err != nil {
		fmt.Println(err)
	}

}
