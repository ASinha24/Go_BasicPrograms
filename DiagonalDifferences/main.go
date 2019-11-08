package main

import (
	"fmt"
	"math"
)

// Complete the diagonalDifference function below.
func diagonalDifference(arr [][]int32) int32 {
	var sum1 int32 = 0
	var sum2 int32 = 0
	for i := 0; i < len(arr); i++ {
		sum1 = sum1 + arr[i][i]
	}
	for j := 0; j < len(arr); j++ {
		sum2 = sum2 + arr[j][len(arr)-j-1]
	}
	return int32(math.Abs(float64(sum1 - sum2)))
}
func plusMinus(arr []int32) {
	var pos, neg, zeros float32
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			zeros = zeros + 1
		} else if arr[i] < 0 {
			neg = neg + 1
		} else {
			pos = pos + 1
		}
	}
	n := float32(len(arr))
	fmt.Println(pos / n)
	fmt.Println(neg / n)
	fmt.Println(zeros / n)

}

func main() {
	arr := [][]int32{
		{4, 5, 7},
		{2, 3, 6},
		{10, 7, 8},
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println()
	}
	result := diagonalDifference(arr)
	fmt.Println(result)
	ar := []int32{2, 3, 4, 6, 0, -2}
	plusMinus(ar)

}
