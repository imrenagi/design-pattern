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
  e.ConnectToWifi()
}

func (e ElectronicKey) ConnectToWifi() {
  fmt.Println("connecting to wifi")
  // e.opener.Open()
}
