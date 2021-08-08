package bridge

import "fmt"

type OS interface {
	OpenLink()
}

type Windows struct {

}

func (w Windows) OpenLink() {
	fmt.Println("opening a link in windows")
}

type Ubuntu struct {

}

func (u Ubuntu) OpenLink() {
	fmt.Println("opening a link in ubuntu")
}