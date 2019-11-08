package main

import (
	"fmt"
)

func intPallindrom(n int) int {
	var rev, remainder int
	for i := 0; n > 0; i++ {
		remainder = n % 10
		rev = rev*10 + remainder
		n = n / 10
	}
	fmt.Println(rev)
	return rev
}

func isPallindrome(s string) {
	rev := ""

	fmt.Println(len(s))
	for i := len(s) - 1; i >= 0; i-- {
		rev += string(s[i])
	}
	if rev == s {
		fmt.Println("Pallindrome")
	} else {
		fmt.Println("not Pallindrome")
	}
}

func main() {
	num := 125
	if intPallindrom(num) == num {
		fmt.Println("Pallindrome")
	} else {
		fmt.Println("not pallindrome")
	}
	isPallindrome("CIVIC")
}
