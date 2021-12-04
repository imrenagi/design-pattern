package mediator

func Execute() {

  t := Teacher{}

  s1 := Student{
    Moderator: &t,
    Name:      "Andi",
  }

  s2 := Student{
    Moderator: &t,
    Name:      "Tini",
  }

  t.P1 = s1
  t.P2 = s2

  s1.PressButton()
  s2.PressButton()
}
