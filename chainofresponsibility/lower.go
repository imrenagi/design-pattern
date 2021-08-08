package chainofresponsibility

import "strings"

type LowerCaseHandler struct {
  upperCase StringHandler
  next StringHandler
}

func (l LowerCaseHandler) Process(s string) string {
  if len(s) < 5 {
    s = l.next.Process(s)
  } else {
    s = strings.ToLower(s)
    if l.upperCase != nil {
      s = l.upperCase.Process(s)
    }
  }
  return s
}

func (l *LowerCaseHandler) SetNext(h StringHandler) {
  l.upperCase = h
}
