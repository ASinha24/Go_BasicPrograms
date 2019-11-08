package main

import (
	"fmt"
)

func main() {
	variadicFunc(1, "red", true, 10.5, []string{"foo", "bar", "baz"}, map[string]int{"apple": 23, "tomato": 13})
}

func variadicFunc(i ...interface{}) {
	for _, v := range i {
		//fmt.Println(v, "--", reflect.ValueOf(v).Kind()) = fmt.Printf("%v -- %T \n", v, v) both is equla , we can use either of these
		fmt.Printf("%v -- %T \n", v, v)
	}
}
