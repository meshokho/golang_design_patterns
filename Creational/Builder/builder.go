package main

import "fmt"

type Builder interface {
	Paint(color string)
	ConstructCar(model string)
	AddFuel(amount int)
}

type Director struct {
	builder Builder
}

func (d Director) GetCarReady() {
	d.builder.ConstructCar("Lada Granta")
	d.builder.Paint("White")
	d.builder.AddFuel(50)
}

type ConcreteBuilder struct {
	car *Car
}

func (c *ConcreteBuilder) Paint(color string) {
	c.car.Color = color
}

func (c *ConcreteBuilder) ConstructCar(model string) {
	c.car.CarModel = model
}

func (c *ConcreteBuilder) AddFuel(amount int) {
	c.car.FuelAmount = amount
}

type Car struct {
	CarModel   string
	Color      string
	FuelAmount int
}

func main() {
	car := new(Car)

	director := Director{builder: &ConcreteBuilder{car: car}}
	director.GetCarReady()

	fmt.Println(car)
}
