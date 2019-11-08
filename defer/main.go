package main

import (
	"fmt"
)

func main() {
	a := 10
	defer fmt.Println("world", a+10)
	fmt.Println("Hello", a)

	//deferrered function calls are pushed in stach and hence called as LIFO order
	fmt.Println("counting")
	for i := 0; i <= 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("Done")
}
