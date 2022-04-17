package main

import (
	"fmt"
	"reflect"
	"strconv"
)

//If vars are not defined, sting = " ", and int initialise 0
// var (
// 	name, course string
// 	module, clip int
// )

var (
	name, course, module = "Sehun Park", "Getting Started with K8s", "4"
	clip                 = 2
	total                int
)

func main() {
	name = "Sehun"

	fmt.Println("Hello World")
	fmt.Println("Name and course are set to", name, "and", course, ".")
	fmt.Println("Module and clip are set to", module, "and", clip, ".")
	fmt.Println("Name is of type", reflect.TypeOf(name))
	fmt.Println("Module is of type", reflect.TypeOf(module))

	fmt.Println("Total is", total)

	iModule, err := strconv.Atoi(module)
	if err == nil {
		total = iModule + clip
		fmt.Println("Module plus clip equals", total)
	}

	var ptr *string = &course
	fmt.Println("Cour name is", ptr, *ptr)
}
