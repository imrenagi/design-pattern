package visitor

type Sex string

const (
  Female Sex = "female"
  Male   Sex = "male"
)

type Dentist interface {
  GetSex() Sex
  DoScaling(p Patient)
  DoToothFilling(p Patient)
}

type Patient interface {
  Accept(d Dentist)
  CountKarangGigi() int
}
