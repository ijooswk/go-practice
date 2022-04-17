package main

import (
	"fmt"
	"reflect"
)

func main() {
	courses := make([]string, 5, 10)

	courses[0] = "Course A"
	courses[1] = "Course B"
	courses[2] = "Course C"

	fmt.Println("Type of", reflect.TypeOf(courses))

	fmt.Println(len(courses), cap(courses))

	for _, course := range courses {
		fmt.Println("course is", course)
	}

	sliceofslice := courses[0:6]
	fmt.Println(sliceofslice)
}
