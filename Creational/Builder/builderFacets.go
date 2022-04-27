package main

import "fmt"

// Если у нас есть несколько Builder'ов.

type Animal struct {
	Color      string
	LegsNumber int
	Tail       bool

	Name, OwnerName string
}

type AnimalBuilder struct {
	animal *Animal
}

func NewAnimalBuilder() *AnimalBuilder {
	return &AnimalBuilder{animal: &Animal{}}
}

func (b *AnimalBuilder) Build() *Animal {
	return b.animal
}

func (b *AnimalBuilder) SetAppearance() *AnimalAppearanceBuilder {
	return &AnimalAppearanceBuilder{*b}
}

func (b *AnimalBuilder) SetOwnerInfo() *AnimalOwnerBuilder {
	return &AnimalOwnerBuilder{*b}
}

type AnimalOwnerBuilder struct {
	AnimalBuilder
}

func (aob *AnimalOwnerBuilder) GiveAName(name string) *AnimalOwnerBuilder {
	aob.animal.Name = name
	return aob
}

func (aob *AnimalOwnerBuilder) ClaimOwner(ownerName string) *AnimalOwnerBuilder {
	aob.animal.OwnerName = ownerName
	return aob
}

type AnimalAppearanceBuilder struct {
	AnimalBuilder
}

func (aab *AnimalAppearanceBuilder) SetColor(color string) *AnimalAppearanceBuilder {
	aab.animal.Color = color
	return aab
}

func (aab *AnimalAppearanceBuilder) SetLegsNumber(number int) *AnimalAppearanceBuilder {
	aab.animal.LegsNumber = number
	return aab
}

func (aab *AnimalAppearanceBuilder) SetTail(is bool) *AnimalAppearanceBuilder {
	aab.animal.Tail = is
	return aab
}

func main() {
	ab := NewAnimalBuilder()
	ab.
		SetAppearance().
		SetColor("Black").
		SetLegsNumber(4).
		SetTail(true).
		SetOwnerInfo().
		GiveAName("Bax").
		ClaimOwner("Max")

	animal := ab.Build()
	fmt.Println(*animal)
}
