package main

// Если мы не хотим, чтобы пользователь напрямую взаимодействовал с объектом.

import "fmt"

type house struct {
	fundament, walls, roof string
}

type HouseBuilder struct {
	house house
}

func (b *HouseBuilder) BuildFundament(material string) *HouseBuilder {
	b.house.fundament = material
	return b
}

func (b *HouseBuilder) BuildWalls(material string) *HouseBuilder {
	b.house.walls = material
	return b
}

func (b *HouseBuilder) PaintRoof(color string) *HouseBuilder {
	b.house.roof = color
	return b
}

func describeHouse(house *house) {
	fmt.Println("House have " + house.fundament + " fundament.")
	fmt.Println("House have " + house.walls + " walls.")
	fmt.Println("House have " + house.roof + " roof.")
}

type build func(*HouseBuilder)

func BuildAHouse(action build) {
	builder := HouseBuilder{}
	action(&builder)
	describeHouse(&builder.house)
}

func main() {
	BuildAHouse(func(b *HouseBuilder) {
		b.
			BuildFundament("marble").
			BuildWalls("wooden").
			PaintRoof("red")
	})
}
