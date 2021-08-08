package bridge

type Color interface {
	Hex() string
}

type Red struct {}

func (r Red) Hex() string {
	return "#E22E5D"
}

type Green struct {}

func (g Green) Hex() string {
	return "#57D658"
}

