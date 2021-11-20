package patient

import (
  "github.com/imrenagi/design-pattern/visitor"
)

type KarangGigiPatient struct {
  SexPreference visitor.Sex
}

func (kgp *KarangGigiPatient) Accept(d visitor.Dentist) {
  if d.GetSex() == kgp.SexPreference {
    d.DoScaling(kgp)
  }
}

func (kgp KarangGigiPatient) CountKarangGigi() int {
  return 0
}

type GigiBerlubangPatient struct {
}

func (gbp *GigiBerlubangPatient) Accept(d visitor.Dentist) {
  d.DoToothFilling(gbp)
}

func (gbp *GigiBerlubangPatient) CountKarangGigi() int {
  return 0
}