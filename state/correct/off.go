package correct


type Off struct {
	LightSwitch *LightSwitch
}

func (o Off) Press() {
	if !o.LightSwitch.IsElectricityOn() {
		return
	}
	o.LightSwitch.ChangeState(
		&On{LightSwitch: o.LightSwitch})
}

func (o Off) CanTurnOnLight() bool {
	return false
}
