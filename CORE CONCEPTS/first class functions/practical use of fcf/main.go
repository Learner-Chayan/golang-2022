package main

import "fmt"

//example 1
type student struct {
	firstName string
	lastName  string
	grade     string
	country   string
}

func filter(s []student, f func(student) bool) []student {
	var r []student

	for _, stu := range s {
		if f(stu) {
			r = append(r, stu)
		}
	}

	return r
}

//example 2

func iMap(n []int, f func(a int) int) []int {
	var r []int
	for _, value := range n {
		r = append(r, f(value))
	}
	return r
}

func main() {
	//example 1......................................
	s1 := student{
		firstName: "Chayan",
		lastName:  "Roy",
		grade:     "A",
		country:   "Bangladesh",
	}

	s2 := student{
		firstName: "Amrit",
		lastName:  "Hasan",
		grade:     "B",
		country:   "Pakistan",
	}

	s := []student{s1, s2}
	f := filter(s, func(stu student) bool {
		return stu.grade == "B"
	})

	fmt.Println(f)

	//example 2................................
	numbers := []int{2, 3, 4, 5, 6}
	result := iMap(numbers, func(a int) int {
		return a * 5
	})
	fmt.Println(result)
}
