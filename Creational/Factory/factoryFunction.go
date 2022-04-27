package main

import (
	"fmt"
	"strings"
)

type Animal struct {
	name, creature string
	tailsNumber    int
}

// NewAnimal This is a Factory function
// We provided a default value
func NewAnimal(name, creature string) *Animal {
	if !strings.HasPrefix(name, "A") {
		fmt.Println(name + " is a bad name, but, mmmm, OKAY")
	}

	return &Animal{
		name:        name,
		creature:    creature,
		tailsNumber: 1,
	}
}

func (c *Animal) SetTails(num int) *Animal {
	c.tailsNumber = num
	return c
}

func main() {
	cat := NewAnimal("Kitur", "cat").SetTails(2)
	dog := NewAnimal("Austin", "dog")

	fmt.Println(cat)
	fmt.Println(dog)
}
