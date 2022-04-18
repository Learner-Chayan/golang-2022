package main

import (
	"fmt"
	practical "learnGo/practical"
)

/*
 Notes for this file -
 1. defer in function
 2. defer in method
 3. defer with arguments
 4. Stack of defers (LAST in FIRST OUT)

*/

func main() {
	//defer in function..
	nums := []int{78, 108, 3, 444, 3, 232, 4444, 442, 33, 332}
	largest(nums)

	//defer in method..
	deferInMethod()

	//defer with argument..
	deferWithArguments()

	//Stack of defers ....
	stackOfDefers()

	//practical use of defer
	practical.Calculate()
}

//defer in function..............................
func largest(nums []int) {
	defer finished()
	fmt.Println("Started finding largest")
	max := nums[0]

	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	fmt.Println("Largest number in ", nums, " is ", max)
}

func finished() {
	fmt.Println("Finished finding largest .")
}

//defer in method............................

type person struct {
	firstName string
	lastName  string
}

func deferInMethod() {
	p := person{
		firstName: "Chayan",
		lastName:  "Roy",
	}

	defer p.fullName()
	fmt.Printf("Welcome")
}

func (p person) fullName() {
	fmt.Printf(" %s %s \n", p.firstName, p.lastName)
}

// defer with arguments ..........................
func deferWithArguments() {
	a := 10
	fmt.Println("Value of a before defer use : ", a)
	defer printArgumentValue(a)

	a++ // though value of a changed but defer function get old value
	fmt.Println("Value of a after increment value ", a)
}

func printArgumentValue(a int) {
	fmt.Println("Value of a in defer function ", a)
}

// Stack of defers (multiple defers - last in first out)........
func stackOfDefers() {
	name := "Chayan"
	fmt.Printf("Original string : %s \n", string(name))
	fmt.Printf("Reversed string : ")
	for _, v := range string(name) {
		defer fmt.Printf("%c", v)
	}

	/*

	 ..............................
	 .defer fmt.Printf("%c",v)| n .
	 ..............................
	 .defer fmt.Printf("%c",v)| a .
	 ..............................
	 .defer fmt.Printf("%c",v)| y .
	 ..............................
	 .defer fmt.Printf("%c",v)| a .
	 ..............................
	 .defer fmt.Printf("%c",v)| h .
	 ..............................
	 .defer fmt.Printf("%c",v)| c .
	 ..............................

	*/

}
