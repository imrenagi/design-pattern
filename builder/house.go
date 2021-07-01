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

func (b *HouseBuilder) BuildWindow() *HouseBuilder {
	if b.house.numOfWindows < 5 {
		b.house.numOfWindows++
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

func (b *HouseBuilder) CreateHouse() (*House, error) {
	if !b.house.hasGarage {
		return nil, fmt.Errorf("a house must have a garage")
	}
	return &b.house, nil
}

func init() {
	hb := HouseBuilder{}
	house, err := hb.
		BuildDoor().
		SetGarage("large").CreateHouse()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(house)
}
