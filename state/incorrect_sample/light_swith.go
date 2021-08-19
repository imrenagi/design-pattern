package incorrect_sample

type LightState string

const (
	On  LightState = "on"
	Off LightState = "off"
	Dim LightState = "dim"
)

type LightSwitch struct {
	State LightState
}

func (ls LightSwitch) Press() {
	if ls.State == Off {
		ls.State = On
	} else if ls.State == On {
		ls.State = Off
	} else {
		ls.State = Dim
	}
}

func (ls LightSwitch) CanTurnOnLight() bool {
	return ls.State == On
}
