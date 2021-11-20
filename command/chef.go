package command

import "fmt"

type FriedRiceParams struct {
  Spiciness int
  Saltiness int
  ExtraEgg  bool
}

type Chef interface {
  CookFriedRice(params FriedRiceParams)
  PrepareSatePadang(params SatePadangParams)
}

type LordAdi struct{}

func (la *LordAdi) CookFriedRice(params FriedRiceParams) {
  fmt.Println("lord adi cooks fried rice %v", params)
}

func (la *LordAdi) PrepareSatePadang(params SatePadangParams) {
  fmt.Println("lord adi cooks sate oadang %v", params)
}

type SatePadangParams struct {
  KetupatCount int
  SateCount    int
  Tounge       bool
}
