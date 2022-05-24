package main

import (
	"fmt"
	"strings"
)

// House provides simple interface to complex system made of simple subsystems.
type House struct {
	roof  *Roof
	floor *Floor
	walls *Walls
}

func (h *House) DescribeHouse() string {
	return strings.Join([]string{
		h.floor.DescribeFloor(),
		h.walls.DescribeWalls(),
		h.roof.DescribeRoof(),
	}, "")
}

func NewDefaultHouse() *House {
	return &House{
		roof: &Roof{
			material: "wood",
			height:   1.4,
		},
		floor: &Floor{
			material: "marble",
		},
		walls: &Walls{
			amount: 4,
			color:  "brown",
		},
	}
}

type Walls struct {
	amount int
	color  string
}

func (w *Walls) DescribeWalls() string {
	return fmt.Sprintf("There are %d walls and they are %s\n", w.amount, w.color)
}

type Floor struct {
	material string
}

func (f *Floor) DescribeFloor() string {
	return fmt.Sprintf("Floor is made of %s\n", f.material)
}

type Roof struct {
	material string
	height   float32
}

func (r *Roof) DescribeRoof() string {
	return fmt.Sprintf("Roof is made of %s and it is %f meters height\n", r.material, r.height)
}

func main() {
	defaultHouse := NewDefaultHouse()
	fmt.Println(defaultHouse.DescribeHouse())

	// We still have access to internal parts of a system.
	roof := &Roof{
		material: "wood",
		height:   5.0,
	}
	fmt.Println(roof.DescribeRoof())
}
