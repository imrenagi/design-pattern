package adapter

import "fmt"

type AudioPlayer interface {
	Play(m Mp3)
}

type Mp3 struct {
	Data []byte
}

func (m Mp3) PlayAudio() {
	fmt.Println("playing mp3 with data" + string(m.Data))
}

type Kaset struct {
	PitaMusik string
}

// Walkman is the legacy
type Walkman struct {}

func (w Walkman) Play(c Kaset) {
	fmt.Println(c.PitaMusik)
}

type Mp3ToKasetAdapter struct {
	Adaptee Walkman
}

func (a Mp3ToKasetAdapter) Play(m Mp3) {
	k := Kaset{}
	k.PitaMusik = string(m.Data)
	a.Adaptee.Play(k)
}

func Execute() {
	mp3 := Mp3{Data: []byte(`this is taylor swift song`)}
	walkman := Walkman{}

	ad := Mp3ToKasetAdapter{Adaptee: walkman}
	ad.Play(mp3)
}