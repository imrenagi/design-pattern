package template

import "fmt"

type email struct {
  otp
}

func (s *email) getMessage(otp string) string {
  return "EMAIL OTP for login is " + otp
}

func (s *email) sendNotification(message string) error {
  fmt.Printf("EMAIL: sending email: %s\n", message)
  return nil
}

func (s *email) publishMetric() {
  fmt.Printf("EMAIL: publishing metrics\n")
}
