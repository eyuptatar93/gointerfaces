package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type rectangle struct {
	width, height float64
}
type circle struct {
	radius float64
}

// ----------------------------------------------------------------------
func (r rectangle) area() float64 {
	return r.height * r.width
}
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// -------------------------------------------------------------------------
func school(s shape) {
	fmt.Println("Area: ", s.area())
}

func main() {
	r := rectangle{width: 10, height: 6}
	school(r) // Area:  60

	c := circle{radius: 10}
	school(c) // Area:  314.1592653589793
}
