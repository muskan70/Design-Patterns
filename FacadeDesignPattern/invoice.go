package main

import "fmt"

type Invoice struct {
	InvoiceNo  int
	OrderNo    int
	BillAmount float64
	ItemList   []*Product
}

var inv = 0

func NewInvoice(orderNo int, amt float64, items []*Product) *Invoice {
	inv++
	return &Invoice{
		InvoiceNo:  inv,
		OrderNo:    orderNo,
		BillAmount: amt,
		ItemList:   items,
	}
}

func (i *Invoice) GenerateInvoice() {
	fmt.Println("Invoice generated successfully for OrderNo:", i.OrderNo)
	fmt.Println("Invoice details:")
	fmt.Println("bill amount:", i.BillAmount)
	for itm := range i.ItemList {
		fmt.Println("ItemName:", i.ItemList[itm].Name, "ItemPrice:", i.ItemList[itm].Price, "ItemQuantity:", i.ItemList[itm].Quantity, "TotalItemPrice:", i.ItemList[itm].Price*float64(i.ItemList[itm].Quantity))
	}

}
