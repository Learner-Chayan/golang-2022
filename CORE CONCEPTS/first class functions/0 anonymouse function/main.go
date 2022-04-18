package main

import "fmt"

//user define function type
type mul func(a int, b int) int

func main() {

	//anonymouse function
	sum := func(a int, b int) int {
		return a + b
	}
	fmt.Println("Summation of 20 and 30 is = ", sum(20, 30))
	fmt.Printf("%T\n", sum)

	//function calling itself
	func(text string) { //receive parameter
		fmt.Println(text)
	}("-Hello first class function , I am calling myself") //arguments

	//user define function
	userDefineFunction()
}

func userDefineFunction() {
	//here mul is user define function type
	var f mul = func(a, b int) int {
		return a * b
	}

	fmt.Println("Multiplication of 20 and 30 is = ", f(20, 30))
}
