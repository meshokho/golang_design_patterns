package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Prototyper interface {
	DeepCopy() *Person
	GetName() string
	GetAddress() PersonAddress
	SetAddress(city, street string, house int)
}

type Person struct {
	Name    string
	Address *PersonAddress
}

func (c *Person) SetAddress(city, street string, house int) {
	address := PersonAddress{
		City:   city,
		Street: street,
		House:  house,
	}

	c.Address = &address
}

type PersonAddress struct {
	City   string
	Street string
	House  int
}

func (c *Person) DeepCopy() *Person {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	e.Encode(c)

	d := gob.NewDecoder(&b)
	result := Person{}
	d.Decode(&result)
	return &result
}

func (c *Person) GetName() string {
	return c.Name
}

func (c *Person) GetAddress() PersonAddress {
	return *c.Address
}

func NewPerson(name string) Prototyper {
	return &Person{
		Name:    name,
		Address: nil,
	}
}

func main() {
	mike := NewPerson("Mike")
	mike.SetAddress("Moscow", "Pushkin", 45)

	fmt.Println(mike.GetName())
	fmt.Println(mike.GetAddress())
}
