package composite

type CEO struct {
  Subordinates []Employee
}

func (c CEO) GetSalary() int {
  return 10
}

func (c CEO) TotalDivisionSalary() int {
  var sum int
  for _, s := range c.Subordinates {
    sum = sum + s.TotalDivisionSalary()
  }
  return c.GetSalary() + sum
}

