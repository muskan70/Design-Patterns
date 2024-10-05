package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

type Email struct {
	email   string
	otpNo   int
	message string
}

func (e *Email) generateOtp() {
	e.otpNo = rand.Intn(100000) + 100000
	fmt.Println("EMAIL: OTP generation done : ", e.otpNo)
}

func (e *Email) saveToCache() {
	fmt.Printf("EMAIL: Saving otp %d wrt to email %s\n", e.otpNo, e.email)
}

func (e *Email) getMessage() {
	e.message = "Your OTP for verification is " + strconv.Itoa(e.otpNo)
	fmt.Println("EMAIL: OTP message created : ", e.message)
}

func (e *Email) sendNotification() {
	fmt.Printf("EMAIL: OTP notification sent to maidId : %s", e.email)
}
