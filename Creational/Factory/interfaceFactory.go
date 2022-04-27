package main

import "fmt"

// В этом случае фабрика предоставляет не объект, а интерфейс, с которым можно взаимодейстовать.

type Animal1 interface {
	SayName()
	CountTails()
}

type animal struct {
	name, creature string
	tailsNumber    int
}

type dog struct {
	name, creature string
	tailsNumber    int
}

func (d dog) SayName() {
	fmt.Println("Dog says: 'RRROORORORRO'. Sadly, it cannot talk :(")
}

func (d dog) CountTails() {
	fmt.Println("Dog says: 'RRROORORORRO'. Sadly, it cannot talk :(")
}

func (a animal) SayName() {
	fmt.Printf("Hello, I'm %s, I'm a %s\n", a.name, a.creature)
}

func (a animal) CountTails() {
	fmt.Printf("Umm, I have... Wait, wait... Oh, i have %d tails!\n", a.tailsNumber)
}

func NewAnimal1(name, creature string, num int) Animal1 {
	if creature == "dog" {
		return &dog{
			name:        name,
			creature:    creature,
			tailsNumber: num,
		}
	}

	return &animal{
		name:        name,
		creature:    creature,
		tailsNumber: num,
	}
}

func main() {
	cat := NewAnimal1("Utty", "cat", 1)
	cat.SayName()
	cat.CountTails()

	dragon := NewAnimal1("Richard", "dragon", 3)
	dragon.CountTails()

	dog := NewAnimal1("Bax", "dog", 1)
	dog.SayName()
	dog.CountTails()
}
