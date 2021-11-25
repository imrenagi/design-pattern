package template

import "fmt"

type iOtp interface {
  genRandomOTP(int) string
  saveOTPCache(string)
  getMessage(string) string
  sendNotification(string) error
  publishMetric()
}

type otp struct {
  iOtp iOtp
}

func (o *otp) genRandomOTP(len int) string {
  randomOTP := "1234"
  fmt.Printf("OTP: generating random otp %s\n", randomOTP)
  return randomOTP
}

func (o *otp) saveOTPCache(otp string) {
  fmt.Printf("OTP: saving otp: %s to cache\n", otp)

  type Car struct {
    Name string
    Color string
  }

  cars := []Car{}
  for _, it := range cars {
    it.Name
  }

}

func (o *otp) genAndSendOTP(otpLength int) error {
  otp := o.iOtp.genRandomOTP(otpLength)
  o.iOtp.saveOTPCache(otp)
  message := o.iOtp.getMessage(otp)
  err := o.iOtp.sendNotification(message)
  if err != nil {
    return err
  }
  o.iOtp.publishMetric()
  return nil
}
