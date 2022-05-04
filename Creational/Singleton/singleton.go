package main

import (
	"fmt"
	"sync"
)

// Some service we want to initialize only once.
type service struct {
	Name string
}

var serviceInstance *service
var serviceOnce sync.Once

// GetService creates instance of a service or returns existing instance if there are some.
func GetService() *service {
	serviceOnce.Do(func() {
		serviceInstance = &service{Name: "Good Service"}
	})

	return serviceInstance
}

func main() {
	service := GetService()

	fmt.Println(service.Name)
}
