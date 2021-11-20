package command

func Execute() {

  lordAdi := &LordAdi{}

  w1 := Waiter{}

  fc := &FriedRiceCommand{
    Chef: lordAdi,
    Param: FriedRiceParams{
      Spiciness: 1,
      Saltiness: 0,
      ExtraEgg:  true,
    },
  }

  w1.AddOrder(fc)

  sc := &SatePadangCommand{
    Chef: lordAdi,
    Params: SatePadangParams{
      KetupatCount: 1,
      SateCount:    20,
      Tounge:       false,
    },
  }

  w1.AddOrder(sc)

  w1.Finalize()
}