package main

import "fmt"

func main() {
	b := 255
	var a *int = &b
	fmt.Printf("Type of a is %T ", a)
	fmt.Println("Address of b is ", a)

	//Returning pointer from a function
	value := hello()
	fmt.Println("The value is = ", *value)
}

func hello() *int {
	a := 256
	return &a
}
