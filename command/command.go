package command

type Interface interface {
  Execute()
}

type FriedRiceCommand struct {
  Chef Chef //receiver
  Param FriedRiceParams
}

func (frc FriedRiceCommand) Execute() {
  frc.Chef.CookFriedRice(frc.Param)
}

type SatePadangCommand struct {
  Chef Chef
  Params SatePadangParams
}

func (spc SatePadangCommand) Execute() {
  spc.Chef.PrepareSatePadang(spc.Params)
}