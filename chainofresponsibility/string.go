package chainofresponsibility

import "fmt"

type StringHandler interface {
  SetNext(h StringHandler)
  Process(s string) string
}


func Execute() {

  // lc -> sr

  sr := SpaceRemoval{}
  lc := LowerCaseHandler{}
  lc.SetNext(&sr)

  st := lc.Process("THE titanic")
  st2 := sr.Process("Hello      world")

  fmt.Println(st2)
  fmt.Println(st)

}
