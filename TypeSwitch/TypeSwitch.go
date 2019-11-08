package main

import (
	"fmt"
)

func main() {
	var x interface{}

	switch i := x.(type) {
	case nil:
		fmt.Printf("type of x is %T", i)
	case int:
		fmt.Printf("Type is int")
	case float64:
		fmt.Printf("Type is float64")
	case bool, string:
		fmt.Printf("Type is bool or string")
	default:
		fmt.Println("don't know the type")
	}
}
