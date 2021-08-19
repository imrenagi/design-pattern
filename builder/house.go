package builder

import (
	"fmt"
	"log"
)

type House struct {
	numOfWindows int
	numOfDoors int
	hasGarage bool
	hasSwimmingPool bool
}

func NewHouse(numOfWindows int, numOfDoors int, hasGarage bool, hasSwimmingPool bool) *House {

	if numOfWindows > 5 {
		numOfWindows = 5
	}

	if numOfDoors > 2 {
		numOfWindows = 2
	}

	if hasGarage && numOfWindows == 5 {
		hasGarage = true
	} else {
		hasGarage = false
	}

	return &House{
		numOfWindows: numOfWindows,
		numOfDoors: numOfDoors,
		hasGarage: hasGarage,
		hasSwimmingPool: hasSwimmingPool,
	}
}

type HouseBuilder struct {
	house House
}

func (b *HouseBuilder) SetWindow(window int) *HouseBuilder {
	if b.house.numOfWindows < window {
		b.house.numOfWindows = window
	}
	return b
}

func (b *HouseBuilder) BuildDoor() *HouseBuilder {
	b.house.numOfDoors++
	return b
}

func (b *HouseBuilder) SetGarage(t string) *HouseBuilder {
	if t == "large" {
		b.house.hasGarage = true
	}
	return b
}

func (b *HouseBuilder) GetResult() (*House, error) {
	if !b.house.hasGarage {
		return nil, fmt.Errorf("a house must have a garage")
	}
	return &b.house, nil
}

func Execute() {
	hb := HouseBuilder{}
	house, err := hb.
		BuildDoor().
		SetGarage("large").GetResult()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(house)
}
