package door

import "fmt"

type Opener interface {
  Open()
  TriggerAlarm()
}

type Door struct {
}

func (d Door) Open() {
  fmt.Println("door is open")
  fmt.Println("ada suara ketika buka pintu")
}

