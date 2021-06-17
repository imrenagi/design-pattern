package informa

import (
	"github.com/imrenagi/design-pattern/abstract_factory"
	"github.com/imrenagi/design-pattern/abstract_factory/informa/product"
)

type Informa struct {
}

func (i *Informa) CreateChair() abstract_factory.Chair {
	return &product.BeanBag{
		// price: product.Price,
		// SoftnessLevel: product.SoftnessLevel,
	}
}

func (i *Informa) CreateCoffeeTable() abstract_factory.CoffeeTable {
	return nil
}

func (i *Informa) CreateSofa() abstract_factory.Sofa {
	return nil
}
