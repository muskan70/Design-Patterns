package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

type SMS struct {
	phoneNo int
	otpNo   int
	message string
}

func (s *SMS) generateOtp() {
	s.otpNo = rand.Intn(1000) + 1000
	fmt.Println("SMS: OTP generation done : ", s.otpNo)
}

func (s *SMS) saveToCache() {
	fmt.Printf("SMS: Saving otp %d wrt to phone %d\n", s.otpNo, s.phoneNo)
}

func (s *SMS) getMessage() {
	s.message = "Your OTP for verification is " + strconv.Itoa(s.otpNo)
	fmt.Println("SMS: OTP message created : ", s.message)
}

func (s *SMS) sendNotification() {
	fmt.Printf("SMS: OTP notification sent to phoneNo : %d", s.phoneNo)
}
