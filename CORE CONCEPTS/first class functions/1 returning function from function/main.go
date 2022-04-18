package main

import "fmt"

func myFunc() func(a int, b int) int {
	f := func(a int, b int) int {
		return a + b
	}
	return f
}

func main() {

	s := myFunc()
	fmt.Println("The summation is : ", s(23, 43))
}
