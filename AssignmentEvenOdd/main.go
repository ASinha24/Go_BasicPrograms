package main

import (
	"fmt"
)

func main() {
	number := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, n := range number {

		if n%2 != 0 {
			fmt.Println(n, "is odd")
		} else {
			fmt.Println(n, "is even")
		}

	}

}
