package main

import "fmt"

type INotification interface {
	SendNotification(message string)
}

type EmailNotification struct {
	UserName string
	EmailId  string
}

func (e *EmailNotification) SendNotification(message string) {
	fmt.Println(e.UserName+" has recieved email notification on emailId:", e.EmailId, "with message:", message)
}

type WhatsAppNotification struct {
	UserName string
	PhoneNo  string
}

func (w *WhatsAppNotification) SendNotification(message string) {
	fmt.Println(w.UserName+" has recieved whatsapp notification on WhatsApp Number:", w.PhoneNo, "with message:", message)
}
