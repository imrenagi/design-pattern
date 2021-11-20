package dentist

import (
  "fmt"

  "github.com/imrenagi/design-pattern/visitor"
)

type Spesialis struct {}

func (s Spesialis) GetSex() visitor.Sex {
  return visitor.Male
}

func (s Spesialis) DoScaling(p visitor.Patient) {
  fmt.Println("mohon maaf tidak melayani scaling")
}

func (s Spesialis) DoToothFilling(p visitor.Patient) {
    fmt.Println("lakukan tambal gigi")
}

type Drg struct {
}

func (d Drg) GetSex() visitor.Sex {
  return visitor.Female
}

func (d Drg) DoScaling(p visitor.Patient) {
      d.ManualScaling()
}

func (d Drg) ManualScaling() {
  fmt.Println("do manual scaling")
}

func (d Drg) AutomaticScaling() {
  fmt.Println("do scaling with machine")
}

func (d Drg) DoToothFilling(p visitor.Patient) {
  fmt.Println("mohon maaf alat tambal gigi lagi rusak")
}

