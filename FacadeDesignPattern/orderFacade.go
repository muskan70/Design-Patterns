package main

type OrderCreationFacade struct {
	OrderID       int
	Items         []*Product
	PaymentMethod IPayment
	Notification  INotification
	Invoice       *Invoice
}

var ord = 0

func NewOrderFacade(payMethod IPayment, notificationType INotification) OrderCreationFacade {
	ord++
	return OrderCreationFacade{
		OrderID:       ord,
		PaymentMethod: payMethod,
		Notification:  notificationType,
	}
}

func (o *OrderCreationFacade) AddProduct(pId int, qty int) error {
	product, err := GetProduct(pId, qty)
	if err != nil {
		return err
	}
	o.Items = append(o.Items, product)
	return nil
}

func (o *OrderCreationFacade) CreateOrder() error {
	billAmount := 0.0
	for i := range o.Items {
		billAmount += o.Items[i].Price * float64(o.Items[i].Quantity)
	}
	o.PaymentMethod.MakePayment(billAmount)
	o.Invoice = NewInvoice(o.OrderID, billAmount, o.Items)
	o.Invoice.GenerateInvoice()
	o.Notification.SendNotification("Order created Successfully")
	return nil
}
