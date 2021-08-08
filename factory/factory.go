package factory

import (
  "fmt"
  "time"

  "github.com/imrenagi/design-pattern/singleton"
)

func Execute() {
  db := singleton.GetDBInstanceWithLock()
  db.GetUser("nagi")
}

type Content interface {
  Play()
}

type CloudContent struct {}

func (c *CloudContent) Play() {
  fmt.Println("this is cloud content about kubernetes")
}

type CodingContent struct {}

func (c *CodingContent) Play() {
  fmt.Println("this is coding content for teaching design pattern")
}

type GrebekRumah struct {}

func (g *GrebekRumah) Play() {
  fmt.Println("ini adalah konten yang sangat diminati para remaja")
}

type ContentType string

const (
  Hiburan ContentType = "hiburan"
  Edukasi ContentType = "edukasi"
)

type ContentCreator interface {
  Produce(time time.Time) Content
  Type() ContentType
}

type Imre struct {
}

func (i *Imre) Produce(time time.Time) Content {
  if time.Weekday().String() == "Tuesday" {
    return &CloudContent{}
  } else if time.Weekday().String() == "Thursday" {
    return &CodingContent{}
  } else {
    return nil
  }
}

func (i Imre) Type() ContentType {
  return Edukasi
}

type Atta struct {
}

func (a *Atta) Produce(time time.Time) Content {
  return &GrebekRumah{}
}

func (a *Atta) Type() ContentType {
  return Hiburan
}
