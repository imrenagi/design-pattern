package template

import "fmt"

func Execute()  {
  smsOTP := &sms{
  }
  o := otp{
    iOtp: smsOTP,
  }
  o.genAndSendOTP(4)

  fmt.Println("")
  emailOTP := &email{}
  o = otp{
    iOtp: emailOTP,
  }
  o.genAndSendOTP(4)

  fmt.Println("")
  phoneOTP := &phone{}
  o = otp{
    iOtp: phoneOTP,
  }
  o.genAndSendOTP(4)
}
