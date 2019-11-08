package main

import (
	"fmt"
)

func staircase(n int32) {
	var i int32
	var space int32
	var k int32
	for i = 1; i <= n; i++ {
		for space = 1; space <= n-i; space++ {
			fmt.Print(" ")
		}
		for k = 1; k <= i; k++ {
			fmt.Print("#")
		}
		fmt.Println()
	}

}

func main() {
	var n int32 = 4
	staircase(n)
}
