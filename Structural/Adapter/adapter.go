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

// Target is an interface which our system needs
type Target interface {
	SendAllEnergy()
	GetEnergy() int
}

type Adapter struct {
	*Adaptee
}

func (a Adapter) GetEnergy() int {
	return a.energyAmount
}

// SendAllEnergy uses existing method from adaptee
func (a Adapter) SendAllEnergy() {
	a.SendEnergy()
}

func NewAdapter(adaptee *Adaptee) Target {
	return &Adapter{adaptee}
}

func main() {
	// Adapter makes Adaptee to be able to use Target interface.
	adapter := NewAdapter(&Adaptee{energyAmount: 10})

	fmt.Println(adapter.GetEnergy())
	adapter.SendAllEnergy()
	fmt.Println(adapter.GetEnergy())
}
