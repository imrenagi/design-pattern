package composite

type VPEng struct {
  Subordinates []Employee
}

func (v VPEng) GetSalary() int {
  return 8
}

func (v VPEng) TotalDivisionSalary() int {
  var sum int
  for _, s := range v.Subordinates {
    sum = sum + s.GetSalary()
  }
  return v.GetSalary() + sum
}
