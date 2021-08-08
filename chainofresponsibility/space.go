package chainofresponsibility

import "strings"

type SpaceRemoval struct {
  next StringHandler
}

func (s *SpaceRemoval) SetNext(h StringHandler) {
  s.next = h
}

func (s SpaceRemoval) Process(st string) string {
  st = strings.Replace(st, " ", "", -1)
  if s.next != nil {
    st = s.next.Process(st)
  }
  return st
}
