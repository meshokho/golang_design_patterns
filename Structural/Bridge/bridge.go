package main

import "fmt"

type HouseInfo interface {
	MeasureTemperature() string
}

type Floorer interface {
	GetTemperature() string
}

type House struct {
	floor Floorer
}

func (h *House) MeasureTemperature() string {
	return h.floor.GetTemperature()
}

func NewHouse(floor Floorer) HouseInfo {
	return &House{floor: floor}
}

type WoodenFloor struct{}

func (w *WoodenFloor) GetTemperature() string {
	return "I am warm"
}

type MetalFloor struct{}

func (m *MetalFloor) GetTemperature() string {
	return "I am cold"
}

type GlassFloor struct{}

func (g *GlassFloor) GetTemperature() string {
	return "I don't know"
}

func main() {
	houseWithWoodenFloor := NewHouse(&WoodenFloor{})

	fmt.Println(houseWithWoodenFloor.MeasureTemperature())
}
