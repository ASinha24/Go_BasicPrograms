package main

import (
	"fmt"
)

func variadicFunction(s ...string) {
	fmt.Println(s)
}

func main() {
	variadicFunction()
	variadicFunction("red")
	variadicFunction("red", "blue", "green", "purple")
}
