package main

type IOTP interface {
	generateOtp()
	saveToCache()
	getMessage()
	sendNotification()
}

type Otp struct {
	iotp IOTP
}

func (o *Otp) generateNsendOTP() {
	o.iotp.generateOtp()
	o.iotp.saveToCache()
	o.iotp.getMessage()
	o.iotp.sendNotification()
}
