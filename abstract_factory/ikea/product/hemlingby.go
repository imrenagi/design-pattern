package product

import "github.com/imrenagi/design-pattern/abstract_factory"

type HEMLINGBY struct {
}

func (h HEMLINGBY) Price() float64 {
	return 1795000
}

func (h HEMLINGBY) Style() abstract_factory.SofaStyle {
	return abstract_factory.EuropeanStyle
}
