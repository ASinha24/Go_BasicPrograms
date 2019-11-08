package main

import (
	"fmt"
	"math"
)

func main() {
	var num [10]float64
	var mean, sum, SD float64

	fmt.Println("enter 10 numbers: ")
	for i := 1; i <= 10; i++ {
		fmt.Printf("enter %d numbers: ", i)
		fmt.Scan(&num[i-1])
		sum += num[i-1]
	}
	mean = sum / 10

	for j := 0; j < 10; j++ {
		SD += math.Pow((num[j] - mean), 2)
	}

	SD = math.Sqrt(SD / 10)
	fmt.Println("The standard deviation is ", SD)

}
