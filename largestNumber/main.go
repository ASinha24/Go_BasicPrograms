package main

import (
	"fmt"
)

func max(a, b, c int) {
	if a > b && a > c {
		fmt.Printf("first number %d is greatest", a)
	} else if b > a && b > c {
		fmt.Printf("Second number %d is greatest", b)
	} else if c > a && c > b {
		fmt.Printf("third number %d is greatest", c)
	} else {
		fmt.Println("please enter three diffrent number")
	}
}

func main() {
	var first, sec, third int
	fmt.Println("enter first number")
	fmt.Scan(&first)
	fmt.Println("enter second number")
	fmt.Scan(&sec)
	fmt.Println("enter third number")
	fmt.Scan(&third)
	max(first, sec, third)
}
