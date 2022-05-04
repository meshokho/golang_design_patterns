package main

import "fmt"

type Adaptee struct {
	energyAmount int
}

func (a *Adaptee) SendEnergy() int {
	toSend := a.energyAmount
	a.energyAmount = 0
	return toSend
}

type Target interface {
	SendAllEnergy() string
	GetEnergy() int
}

type Adapter struct {
	*Adaptee
}

func (a Adapter) GetEnergy() int {
	return a.energyAmount
}

// SendAllEnergy uses existing method from adaptee
func (a Adapter) SendAllEnergy() string {
	a.SendEnergy()
	return "Energy sent"
}

func NewAdapter(adaptee *Adaptee) Target {
	return &Adapter{adaptee}
}

func main() {
	adapter := NewAdapter(&Adaptee{energyAmount: 10})

	fmt.Println(adapter.GetEnergy())
	fmt.Println(adapter.SendAllEnergy())
	fmt.Println(adapter.GetEnergy())
}
