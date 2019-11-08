package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("reverse integer sort")
	num := []int{50, 90, 30, 10, 50}
	sort.Sort(sort.Reverse(sort.IntSlice(num)))
	fmt.Println(num)

	fmt.Println("reverse string sort")
	str := []string{"Elon Musk", "Bill Gates", "Satya Nadela", "Sundar Pichai"}
	sort.Sort(sort.Reverse(sort.StringSlice(str)))
	fmt.Println(str)
}
