package door

import "fmt"

type Opener interface {
  Open()
}

type Door struct {
}

func (d Door) Open() {
  fmt.Println("door is open")
}

