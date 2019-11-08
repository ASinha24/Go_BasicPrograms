package main

import (
	"fmt"
	"math"
)

func main() {
	var radius float64
	var area float64
	var perimeter float64
	const Pi = 3.14

	fmt.Print("Please enter radius   ")
	fmt.Scanln(&radius)

	area = Pi * math.Pow(radius, 2)
	perimeter = 2 * Pi * radius

	fmt.Println("area of circle :", area)
	fmt.Println("Perimeter of circle: ", perimeter)
}
