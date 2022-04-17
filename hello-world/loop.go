package main

import (
	"fmt"
)

func main() {
	for timer := 10; timer >= 0; timer-- {

		if timer == 0 {
			fmt.Println("Boom!!")
			break
		}

		fmt.Println(timer)
	}

	coursesInProg := []string{
		"A course",
		"B course",
		"C course",
		"D course"}

	for _, course := range coursesInProg {
		fmt.Println("course name is ", course)
	}
}
