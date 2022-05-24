package main

import "fmt"

type House interface {
	Describe(roof, wall, floor string) string
}

type HouseFactory struct {
	pool map[string]House
}

func (f *HouseFactory) GetHouse(houseType string) House {
	if f.pool == nil {
		f.pool = make(map[string]House)
	}
	if _, ok := f.pool[houseType]; !ok {
		f.pool[houseType] = &ConcreteHouse{houseType: houseType}
	}
	return f.pool[houseType]
}

type ConcreteHouse struct {
	houseType string
}

func (c *ConcreteHouse) Describe(roof, wall, floor string) string {
	return fmt.Sprintf("This is a %s. It has %s roof, %s wall, %s floor", c.houseType, roof, wall, floor)
}

func main() {
	factory := new(HouseFactory)

	fiveFloors := factory.GetHouse("five-floor")
	skyscraper := factory.GetHouse("skyscraper")
	// We do not create a new skyscraper here, but pointing to an existing one.
	skyscraper2 := factory.GetHouse("skyscraper")

	fmt.Println(fiveFloors.Describe("a", "b", "c"))
	fmt.Println(skyscraper.Describe("a", "b", "c"))
	fmt.Println(skyscraper2.Describe("a", "b", "c"))
}
