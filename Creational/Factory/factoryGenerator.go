package main

import "fmt"

type Person struct {
	Name, Gender string
	Age          int
}

type PersonFactory struct {
	Gender string
}

// NewPersonFactory functional approach
func NewPersonFactory(gender string) func(name string, age int) *Person {
	return func(name string, age int) *Person {
		return &Person{
			Name:   name,
			Gender: gender,
			Age:    age,
		}
	}
}

func NewPersonStructuralFactory(gender string) *PersonFactory {
	return &PersonFactory{Gender: gender}
}

func (f *PersonFactory) SetNameAndAge(name string, age int) *Person {
	return &Person{
		Name:   name,
		Gender: f.Gender,
		Age:    age,
	}
}

func main() {
	menFactory := NewPersonFactory("men")
	womenFactory := NewPersonFactory("women")

	alex := menFactory("Alex", 30)
	jane := womenFactory("Jane", 25)

	fmt.Println(*alex)
	fmt.Println(*jane)

	helicopterFactory := NewPersonStructuralFactory("helicopter")

	rhino := helicopterFactory.SetNameAndAge("Rhino", 3)
	fmt.Println(*rhino)
}
