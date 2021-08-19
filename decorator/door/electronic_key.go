package door

import "fmt"

func NewElectronicKeyDoor(opener Opener) Opener {
  return &ElectronicKey{
    opener: opener,
  }
}

type ElectronicKey struct {
  opener Opener
}

func (e ElectronicKey) Open() {
  fmt.Println("this is electronic key")
  ok := e.ConnectToWifi()
  if ok {
    e.opener.Open()
  }
}

func (e ElectronicKey) ConnectToWifi() bool {
  fmt.Println("connecting to wifi")
  // e.opener.Open()
  return false
}
