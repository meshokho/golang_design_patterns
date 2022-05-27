package main

import "fmt"

type Describer interface {
	Describe()
}

type House struct {
	NumberOfFloors int
}

func NewHouse(numberOfFloors int) *House {
	return &House{NumberOfFloors: numberOfFloors}
}

func (h House) Describe() {
	fmt.Printf("This is a house with %d floors\n", h.NumberOfFloors)
}

type BeautifulHouse struct {
	house House
}

func NewBeautifulHouse(house House) *BeautifulHouse {
	return &BeautifulHouse{house: house}
}

func (b BeautifulHouse) Describe() {
	fmt.Printf("This a very beautiful amazing omg house with %d floors\n", b.house.NumberOfFloors)
}

func main() {
	h := NewHouse(5)
	bh := NewBeautifulHouse(House{NumberOfFloors: 5})

	h.Describe()
	bh.Describe()
}
