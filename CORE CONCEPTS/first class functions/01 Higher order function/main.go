package main

import "fmt"

// takes one or more functions as arguments
// returns a function as its result

func main() {
	a := func(x int, y int) int {
		return x + y
	}

	higherOrder(a)
}

func higherOrder(a func(x int, y int) int) {
	fmt.Println("Summation output ", a(10, 20))
}
