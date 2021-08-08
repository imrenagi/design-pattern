package door

import (
  "fmt"
  "math/rand"
)

func NewMagicalDoor(o Opener) Opener {
  return &MagicalKey{opener: o}
}

type MagicalKey struct {
  opener Opener
}

func (e MagicalKey) Open() {
  fmt.Println("magical key is dangerous")
  if e.CanWeUseMagic() {
    fmt.Println("magic is working")
  } else {
    fmt.Println("cant use magic. use your hand")
  }
  e.opener.Open()
}

func (e MagicalKey) CanWeUseMagic() bool {
  return rand.Int() > 0
}