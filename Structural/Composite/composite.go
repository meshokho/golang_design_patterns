package main

import "fmt"

type Worker interface {
	Name() string
	AddSub(worker Worker)
	Workers() []Worker
	Print() string
}

type Owner struct {
	name    string
	workers []Worker
}

func (o Owner) Name() string {
	return o.name
}

func (o *Owner) AddSub(worker Worker) {
	o.workers = append(o.workers, worker)
}

func (o Owner) Workers() []Worker {
	return o.workers
}

func (o Owner) Print() string {
	result := "My name is " + o.Name() + " and here are my workers:\n"
	for _, val := range o.Workers() {
		result += val.Print()
	}

	return result
}

type Employee struct {
	name string
}

func (e Employee) Name() string {
	return e.name
}

func (e Employee) AddSub(worker Worker) {}

func (e Employee) Workers() []Worker {
	return []Worker{}
}

func (e Employee) Print() string {
	return "    My name is " + e.Name() + "\n"
}

func NewOwner(name string) *Owner {
	return &Owner{
		name: name,
	}
}

func NewEmployee(name string) *Employee {
	return &Employee{name: name}
}

func main() {
	alex := NewOwner("Alex")

	sam := NewEmployee("Sam")
	matt := NewEmployee("Matt")
	jane := NewEmployee("Jane")

	alex.AddSub(sam)
	alex.AddSub(matt)
	alex.AddSub(jane)

	// We use one interface for different structs - for branch and for leafs
	fmt.Println(alex.Print())
	fmt.Println(sam.Print())
}
