package main

import "fmt"

type HouseDescriber interface {
	Describe() string
}

type House struct {
	NumberOfFloors int
}

func (h House) Describe() string {
	return fmt.Sprintf("This house has %d floors", h.NumberOfFloors)
}

type ColoredHouse struct {
	House House

	Color string
}

func (c ColoredHouse) Describe() string {
	return fmt.Sprintf("%s and it is %s", c.House.Describe(), c.Color)
}

func main() {
	house := &House{NumberOfFloors: 5}
	fmt.Println(house.Describe())

	blackHouse := &ColoredHouse{
		House: *house,
		Color: "Black",
	}
	fmt.Println(blackHouse.Describe())
}
