package facade

import "github.com/imrenagi/design-pattern/facade/gofood"

type Ticket struct {
  User string
  Type string
  Message string
}

type Response struct {
  Status string
}

func HandleTicket(t Ticket) Response {
  switch t.Type {
  case "gofood":
    response := gofood.CheckMessage(t.Message)
    return Response{Status: response}
  case "goride":
  case "gosend":
  }

  return Response{
    Status: "customer service is busy",
  }
}
