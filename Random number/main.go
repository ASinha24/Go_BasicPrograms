package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	fmt.Println("My Favorite number is ", r.Intn(10))
}
