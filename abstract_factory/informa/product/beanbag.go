package product

type BeanBag struct {
	price float64
	SoftnessLevel int
}

func (b BeanBag) Price() float64 {
	return b.price
}

func (b BeanBag) IsIoTEnabled() bool {
	return false
}

func (b BeanBag) IsSoft() bool {
	return b.SoftnessLevel > 10
}

