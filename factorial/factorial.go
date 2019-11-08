package main

import (
	"fmt"
)

var fact uint64 = 1
var number int

func factorial(n int) uint64 {
	if n < 0 {
		fmt.Println("value should be positive")
	} else {
		for i := 1; i <= n; i++ {
			fact *= uint64(i)
		}
	}
	return fact
}
func main() {
	fmt.Println("Please enter a number")
	fmt.Scan(&number)
	fmt.Println("factorial of a number is: ", factorial(number))
}
