package main

import (
  "fmt"
  "time"

  "github.com/imrenagi/design-pattern/factory"
)


func main() {

  var contentCreator factory.ContentCreator
  contentCreator = &factory.Imre{}

  content := contentCreator.Produce(time.Now())
  content.Play()
  fmt.Println(contentCreator.Type())

}
