package main

import "fmt"

func miniMaxSum(arr []int32) {

	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			arr[i+1], arr[i] = arr[i], arr[i+1]
		}
	}
	fmt.Println(arr)
}

func main() {
	arr := []int32{3, 2, 1, 4, 5}
	miniMaxSum(arr)
}
