package main

import "fmt"

type House struct {
	Color       string
	FloorNumber int
	Address     string
}

const (
	SimpleHouse = iota
	BigHouse
	Skyscraper
)

func NewHouse(houseType int) *House {
	switch houseType {
	case SimpleHouse:
		return &House{
			Color:       "white",
			FloorNumber: 1,
			Address:     "",
		}
	case BigHouse:
		return &House{
			Color:       "gray",
			FloorNumber: 10,
			Address:     "",
		}
	case Skyscraper:
		return &House{
			Color:       "black",
			FloorNumber: 100,
			Address:     "",
		}
	default:
		panic("House type not supported")
	}
}

func main() {
	simple := NewHouse(SimpleHouse)
	simple.Address = "Moscow"
	fmt.Println(simple)

	scraper := NewHouse(Skyscraper)
	scraper.Address = "New York"
	fmt.Println(scraper)
}
