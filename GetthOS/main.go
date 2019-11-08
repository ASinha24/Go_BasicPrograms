package main

import (
	"fmt"
	"runtime"
)

func main() {
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
}
