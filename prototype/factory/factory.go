package factory

import "github.com/imrenagi/design-pattern/prototype"

type Factory struct {}

func (f Factory) CopyRobot(r *prototype.Robot) *prototype.Robot {
  return r.Clone()
}
