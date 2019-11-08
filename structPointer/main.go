package main

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1,2}
	v2 = Vertex{X:1}
	v3 = Vertex{}
	p = &Vertex{1,2}
)

func main() {
	// v := Vertex{1, 2}
	// p := &v
	// p.X = 1e9
	// fmt.Println(v)
	// fmt.Println(*p)
	// fmt.Println(v.X, v.Y)
	// fmt.Println((*p).X, (*p).Y)
	fmt.Println("**********", v1, p, v2, v3)
	fmt.Printf("type of p is %T",p)
}
