package main

import (
	"fmt"
	"os"
)

func main() {
	name := os.Getenv("USERNAME")
	course := "Getting Started with Kubernetes"

	fmt.Println("Hi", name, "your course is", course)
	updateCourse(course)

	fmt.Println("You're currenly watching", course)

	updateCourseReference(&course)

	fmt.Println("You're currenly watching2", course)

	// declare const
	const miles = 186000
	fmt.Println("The speed of light is", miles)

}

// this way, parameter will be copied not passed directly
func updateCourse(course string) string {
	course = "Getting Started with docker"
	fmt.Println("Updated course to", course)
	return course
}

// passing by reference
func updateCourseReference(course *string) string {
	*course = "Getting Started with docker"
	fmt.Println("Updated course to", course)
	return *course
}
