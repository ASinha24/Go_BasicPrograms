package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("******************Contains function********************")
	//Contains function
	fmt.Println(strings.Contains("Australia", "lia"))
	fmt.Println(strings.Contains("Australia", "tia"))

	fmt.Println("******************Count Function***********************")
	fmt.Println(strings.Count("Australia", "a"))
	fmt.Println(strings.Count("Australia", "A"))

	fmt.Println("******************HasPrefix Function***********************")
	fmt.Println(strings.HasPrefix("Australia", "Aus"))
	fmt.Println(strings.HasPrefix("Australia", "aus"))

	fmt.Println("******************HasSuffix Function***********************")
	fmt.Println(strings.HasSuffix("Australia", "lia"))
	fmt.Println(strings.HasSuffix("Australia", "tia"))

	fmt.Println("******************Index Function***********************")
	fmt.Println(strings.Index("Australia", "lia"))
	fmt.Println(strings.Index("Australia", "u"))

	fmt.Println("******************Join Function***********************")
	fmt.Println(strings.Join([]string{"Australia", "Japan", "India"}, "-"))
	fmt.Println(strings.Join([]string{"Australia", "Germany", "Japan"}, " "))

	fmt.Println("******************Repeat Function***********************")
	fmt.Println(strings.Repeat("Australia ", 3))

	fmt.Println("******************Replace Function***********************")
	fmt.Println(strings.Replace("Australia", "A", "B", 1))
	fmt.Println(strings.Replace("Australia", "a", "b", 1))
	fmt.Println(strings.Replace("Australia", "a", "b", 2))
	fmt.Println(strings.Replace("Australia", "a", "b", 3))

	fmt.Println("******************split Function***********************")
	fmt.Println(strings.Split("Australia-India-America", "-"))

	fmt.Println("******************ToLower Function***********************")
	fmt.Println(strings.ToLower("AusTRALIA"))

	fmt.Println("******************ToUpper Function***********************")
	fmt.Println(strings.ToUpper("Australia"))
}
