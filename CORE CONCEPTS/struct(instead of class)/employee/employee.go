package employee

import "fmt"

type Employee struct {
	FirstName   string
	LastName    string
	TotalLeaves int
	LeavesTake  int
}

/*Firstly there was no New method */
func New(firstName string, lastName string, totalLeaves int, leavesTaken int) Employee {
	e := Employee{
		FirstName:   firstName,
		LastName:    lastName,
		TotalLeaves: totalLeaves,
		LeavesTake:  leavesTaken,
	}

	return e
}

func (e Employee) LeavesRemaining() {
	fmt.Printf("%s %s has %d leaves remaining \n", e.FirstName, e.LastName, (e.TotalLeaves - e.LeavesTake))
}
