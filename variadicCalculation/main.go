package main

import (
	"fmt"
)

//Normal function parameter with variadic function parameter
func calculation(str string, y ...int) int64 {
	area := 1
	for _, val := range y {
		if str == "Rect" {
			area *= val
		} else if str == "square" {
			area = val * val
		}
	}
	return int64(area)
}

func main() {
	fmt.Println(calculation("Rect", 20, 30))
	fmt.Println(calculation("square", 20))
}
