package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	perimeter() float64
}

type square struct {
	side float64
}

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (s *square) area() float64 {
	return s.side * s.side
}

func (s *square) perimeter() float64 {
	return 4 * s.side
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.height * r.width)
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func main() {

	r := rectangle{height: 3, width: 4}
	fmt.Println("Area of rectangle ", r.area())
	fmt.Println("perimeter of rectangle ", r.perimeter())

	c := &circle{radius: 5}
	fmt.Println("Area of circle ", c.area())
	fmt.Println("Perimeter of cicle ", c.perimeter())

	s := &square{side: 5}
	fmt.Println("Area of Square ", s.area())
	fmt.Println("Perimeter of Square ", s.perimeter())

}
