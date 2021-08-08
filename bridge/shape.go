package bridge

import "fmt"

type Shape interface {
	Draw()
	Area() float64
}

type Circle struct {
	Color Color
	Radius int
}

func (c Circle) Draw() {
	if c.Color != nil {
		fmt.Println("drawing a circle with color " + c.Color.Hex())
	} else {
		fmt.Println("drawing a circle")
	}
}

func (c Circle) Area() float64 {
	return 3.14 * float64(c.Radius) * float64(c.Radius)
}

func (c Circle) Clone() Circle {
	return Circle{
		Color:  c.Color,
		Radius: c.Radius,
	}
}

type Square struct {
	Color Color
	Length int
}

func (s Square) Draw() {
	fmt.Println("drawing a square" + s.Color.Hex())
}

func (s Square) Area() float64 {
	area := s.Length * s.Length
	return float64(area)
}

func Execute() {
	c := Circle{
		Radius: 10,
	}
	c.Draw() // drawing a circle

	greenCircle2 := c.Clone()
	greenCircle2.Color = Green{}
	greenCircle2.Draw() // drawing a circle with color #57D658

	redCircle := Circle{
		Radius: 12,
		Color: Red{},
	}
	redCircle.Draw() // drawing a circle with color #E22E5D

	greenSquare := Square{
		Length: 5,
		Color: Green{},
	}
	greenSquare.Draw() // drawing a square #57D658
}



