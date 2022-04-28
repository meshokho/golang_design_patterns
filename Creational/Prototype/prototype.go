package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

// Car Example object with different pointers fields(slice Passengers is also a pointer).
type Car struct {
	Model      string
	Passengers []Passenger
	Appearance *Appearance
}

type Passenger struct {
	Name string
	Age  int
}

type Appearance struct {
	Color        string
	WheelsNumber int
}

// DeepCopy Simple deep copying through serialization. It is possible to use other encoders/decoders e.g. JSON
// I don't handle errors here
func (c *Car) DeepCopy() *Car {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	e.Encode(c)

	d := gob.NewDecoder(&b)
	result := Car{}
	d.Decode(&result)
	return &result
}

func main() {
	// Creating prototype of an object
	tesla := Car{
		Model: "Tesla X",
		Passengers: []Passenger{
			{"Elon", 44},
			{"Jack", 23},
		},
		Appearance: &Appearance{
			Color:        "Black",
			WheelsNumber: 4,
		},
	}

	// Creating copy and change it
	whiteTesla := tesla.DeepCopy()
	whiteTesla.Appearance.Color = "white"

	fmt.Println(tesla, tesla.Appearance)
	fmt.Println(whiteTesla, whiteTesla.Appearance)
}
