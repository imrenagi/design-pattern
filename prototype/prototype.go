package prototype

import (
  "fmt"

  "github.com/imrenagi/design-pattern/singleton"
)

type Prototype interface {
  Clone() Prototype
}

type Robot struct {
  ID string
  Name string
  processor string // private properties
}

func NewRobot(p Prototype) Prototype {
  return p.Clone()
}

func (r Robot) Clone() Prototype {
  return &Robot{
    Name: r.Name,
    ID: r.ID,
    processor: r.processor,
  }
}

func Execute() {
  r := Robot{
    Name: "imre",
    ID: "1",
  }

  r2 := r.Clone()
  robot2, ok := r2.(*Robot)
  if ok {
    robot2.Name = "John"
  }

  fmt.Println(r.Name) // imre
  fmt.Println(robot2.Name) //John

  db := singleton.GetDBInstanceWithOnce()
  db.GetUser("imrenagi")
}