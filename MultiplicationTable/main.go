package main

import (
	"fmt"
)

func main() {
	fmt.Println("Please enter the number")
	var num int
	fmt.Scan(&num)

	fmt.Println("**************")
	for i := 1; i <= 10; i++ {
		fmt.Println(num, "X", i, "=", num*i)
	}

}
