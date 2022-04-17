package main

import (
	"fmt"
	"os"
)

func main() {

	if ddLengthMins, cawLengthMins := 275, 30; ddLengthMins > cawLengthMins {
		fmt.Println("ddLengthMin is greater than cawLengthMin")
	} else {
		fmt.Println("cawLengthMin is greater than ddLengthMin")
	}

	switch "Kubernetes" {
	case "kubernetes":
		fmt.Println("Hello World")
		fallthrough
	case "Kubernetes":
		fmt.Println("Hello World2")
		fallthrough
	case "K8s":
		fmt.Println("Hello World3")
		fallthrough
	case "Is623tio":
		fmt.Println("Hello World4")
		fallthrough
	default:
		fmt.Println("Default")
	}

	_, err := os.Open("./test1.txt")

	if err != nil {
		fmt.Println("This is the error code:", err)
	}
}
