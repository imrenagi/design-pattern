package ikea

import (
	"github.com/imrenagi/design-pattern/abstract_factory"
	"github.com/imrenagi/design-pattern/abstract_factory/ikea/product"
)

type Ikea struct {
}

func (i *Ikea) CreateChair() abstract_factory.Chair {
	return &product.Leifarne{}
}

func (i *Ikea) CreateCoffeeTable() abstract_factory.CoffeeTable {
	return &product.VITTSJÖ{}
}

func (i *Ikea) CreateSofa() abstract_factory.Sofa {
	return &product.HEMLINGBY{}
}
