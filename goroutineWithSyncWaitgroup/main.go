package main

import (
	"fmt"
	"sync"
	"time"
)

type testConcurrency struct {
	min     int
	max     int
	country string
}

func printCountry(test *testConcurrency, groupTest *sync.WaitGroup) {
	for i := test.max; i > test.min; i-- {
		time.Sleep(2 * time.Millisecond)
		fmt.Println(test.country)
	}
	fmt.Println()
	groupTest.Done()
}

func main() {
	groupTest := new(sync.WaitGroup)

	Japan := new(testConcurrency)
	India := new(testConcurrency)
	China := new(testConcurrency)

	Japan.country = "Japan"
	Japan.min = 0
	Japan.max = 5

	India.country = "India"
	India.min = 0
	India.max = 5

	China.country = "China"
	China.min = 0
	China.max = 5

	go printCountry(Japan, groupTest)
	go printCountry(India, groupTest)
	go printCountry(China, groupTest)

	groupTest.Add(3)
	groupTest.Wait()

}
