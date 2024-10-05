package main

import "fmt"

func main() {
	smsOTP := &SMS{phoneNo: 7208509381}
	o := Otp{
		iotp: smsOTP,
	}
	o.generateNsendOTP()

	fmt.Println("")

	emailOTP := &Email{email: "muskanmla7d@gmail.com"}
	o = Otp{
		iotp: emailOTP,
	}
	o.generateNsendOTP()
}
