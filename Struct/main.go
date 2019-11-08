package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipcode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipcode: 9400,
		},
	}
	//jimpointer := &jim
	//jimpointer.updateName("Jimmy")
	jim.updateName("Jimmy") //pointer shortcut
	jim.print()

}

func (pointerToPerson *person) updateName(newfirstName string) {
	(*pointerToPerson).firstName = newfirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
