package main

import "fmt"

func main() {
	//make(map[type of key]type of value)
	employeeSalary := make(map[string]int)
	employeeSalary["Ashikur"] = 10000
	employeeSalary["Manik"] = 10000
	employeeSalary["Chayan"] = 1000
	fmt.Println(employeeSalary)

	//another map
	employeeAge := map[string]int{
		"chayan":  25,
		"Ashikur": 23,
		"Manik":   23,
	}
	employeeAge["Shankar"] = 28
	fmt.Println(employeeAge)

	//Zero value of a map
	var empDetails map[int]string
	//empDetails[0] = "Name : ddd , roll :343" // error
	fmt.Println(empDetails)

	//Checking if a key exists
	value, ok := employeeSalary["shankar"]
	if ok {
		fmt.Println("value is ", value)
	}
	fmt.Println("key doesn't exist")

	//Iterate over all elements in a map
	for key, value := range employeeSalary {
		fmt.Printf("Name is %s and salary is %d \n", key, value)
	}

	//Deleting items from a map
	delete(employeeAge, "chayan")
	fmt.Println("After delete chayan from employeeAge")
	fmt.Println(employeeAge)

	//Map of structs
	type employee struct {
		height string
		id     int
	}
	emp1 := employee{
		height: "5.7",
		id:     1,
	}
	emp2 := employee{
		height: "5.6",
		id:     2,
	}

	empInfo := map[string]employee{
		"chayan": emp1,
		"manik":  emp2,
	}

	for name, info := range empInfo {
		fmt.Printf("Name is %s and id is %d and height is %s \n", name, info.id, info.height)
	}

	//Length of map
	fmt.Println("Length of the map is = ", len(employeeSalary))

	//Maps are reference types ---- important
	employeeSalary2 := map[string]int{
		"steve": 12000,
		"jamie": 15000,
		"mike":  9000,
	}
	fmt.Println("Original employee salary", employeeSalary2)
	modified := employeeSalary2
	modified["mike"] = 18000
	fmt.Println("Employee salary changed", employeeSalary2)
}
