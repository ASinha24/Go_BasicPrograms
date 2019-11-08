package main

import (
	"fmt"
)

func main() {
	var grade string = "B"
	var marks int = 90

	//passed the int value
	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 50, 60, 70:
		grade = "C"
	default:
		grade = "D"
	}
	//passing 0 value
	switch {
	case grade == "A":
		fmt.Printf("Excellent!\n")
	case grade == "B", grade == "C":
		fmt.Printf("Well done\n")
	case grade == "D":
		fmt.Printf("You are passed\n")
	case grade == "F":
		fmt.Printf("Better Try again\n")
	default:
		fmt.Printf("Invalid Grade\n")
	}
	fmt.Printf("Your grade is %s", grade)

}
