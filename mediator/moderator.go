package mediator

import (
  "fmt"
)

type Moderator interface {
  Notify(sender Participant) error
}

type Teacher struct {
  P1, P2 Participant
}

func (t *Teacher) Notify(sender Participant) error {
  var answer string
  if t.P1.GetName() == sender.GetName() {
    answer = t.P1.Answer()

  } else {
    answer = t.P2.Answer()
  }
  fmt.Println(answer)
  return nil
}

