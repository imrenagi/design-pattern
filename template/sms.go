package template

import "fmt"

type sms struct {
  otp
}

func (s *sms) getMessage(otp string) string {
  return "SMS OTP for login is " + otp
}

func (s *sms) sendNotification(message string) error {
  fmt.Printf("SMS: sending sms: %s\n", message)
  return nil
}

func (s *sms) publishMetric() {
  fmt.Printf("SMS: publishing metrics\n")
}
