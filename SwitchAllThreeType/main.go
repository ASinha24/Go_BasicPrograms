package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("when is Saturday")
	today := time.Now().Weekday() // it returns weekday type which is ultimately struct of int type

	switch time.Saturday { //case result should equal to the switch result
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	fmt.Print("go is running on")
	switch os := runtime.GOOS; os { //short assignment "os := runtime.GOOS;" to get the os ang goarch value run command "go tool dist list"
	case "android":
		fmt.Print("android")
	case "darwin":
		fmt.Print("darwin")
	case "freebsd":
		fmt.Print("freebsd")
	default:
		fmt.Printf(" %s. \n", os)
	}
	fmt.Println(runtime.GOARCH) //GOARCH is the running program's architecture target: one of 386, amd64, arm, s390x, and so on

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good Morning")
	case t.Hour() < 17:
		fmt.Println("Good Afternoon")
	default:
		fmt.Println("Good Evening")
	}
}
