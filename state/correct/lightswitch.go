package correct

import "fmt"

type State interface {
	Press()
	CanTurnOnLight() bool
}

func NewLightSwitch() *LightSwitch {
	ls := &LightSwitch{}
	ls.ChangeState(&Off{ls})
	return ls
}

type LightSwitch struct {
	State State
}

func (ls LightSwitch) Press() {
	ls.State.Press()
	fmt.Println("udah selesai")
}

func (ls LightSwitch) CanTurnOnLight() bool {
	return ls.State.CanTurnOnLight()
}

func (ls *LightSwitch) ChangeState(state State) {
	ls.State = state
}

func (ls LightSwitch) IsElectricityOn() bool {
	return true
}
