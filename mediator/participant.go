package mediator

import "fmt"

type Participant interface {
  GetName() string
  Answer() string
}

type Student struct {
  Moderator Moderator
  Name string
}

func (h Student) GetName() string {
  return h.Name
}

func (h Student) PressButton() {
  fmt.Println(h.Name, "pressed button")
  h.Moderator.Notify(h)
}

func (h Student) Answer() string {
  return fmt.Sprintf("%s menjawab pertanyaan", h.Name)
}

