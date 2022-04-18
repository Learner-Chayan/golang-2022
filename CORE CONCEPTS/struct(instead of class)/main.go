package main

import (
	"fmt"
	"oop/employee"
)

func main() {
	fmt.Println("Hello")

	// e := employee.Employee{
	// 	FirstName:   "Chayan",
	// 	LastName:    "Roy",
	// 	TotalLeaves: 44,
	// 	LeavesTake:  34,
	// }

	e := employee.New("Chayan", "Roy", 44, 34)
	e.LeavesRemaining()
}
