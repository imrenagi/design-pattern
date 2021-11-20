package template

import "fmt"

type phone struct {
  otp
}

func (p *phone) genRandomOTP(length int) string {
  randomOTP := "12345678910"
  fmt.Println("PHONE: genRandomOTP ", randomOTP)
  return randomOTP
}

func (p *phone) getMessage(otp string) string {
  fmt.Println("PHONE: getMessage OTP for login is ", otp)
  return "This is bot, your otp code for login is " + otp
}

func (p *phone) sendNotification(msg string) error {
  fmt.Println("PHONE: creating a phone call with message ", msg)
  return nil
}

func (p *phone) publishMetric() {
  fmt.Println("PHONE: publishMetric")
}
