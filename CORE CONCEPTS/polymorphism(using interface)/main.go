package main

import "fmt"

type Income interface {
	calculate() int
	source() string
}

type FixedBilling struct {
	projectName  string
	biddedAmount int
}

type TimeAndMaterial struct {
	projectName string
	noOfHours   int
	hourlyRate  int
}

func (fb FixedBilling) calculate() int {
	return fb.biddedAmount
}

func (fb FixedBilling) source() string {
	return fb.projectName
}

func (tm TimeAndMaterial) calculate() int {
	return tm.noOfHours * tm.hourlyRate
}

func (tm TimeAndMaterial) source() string {
	return tm.projectName
}

/*Create a function to use the struct*/
func calculateIncome(ic []Income) {
	var netincome int = 0
	for _, income := range ic {
		fmt.Printf("Income from %s=$%d\n", income.source(), income.calculate())
		netincome += income.calculate()
	}
	fmt.Printf("Net income of organization = $%d", netincome)
}

func main() {
	project1 := FixedBilling{projectName: "Porject 1", biddedAmount: 500}
	project2 := FixedBilling{projectName: "Porject 2", biddedAmount: 400}
	project3 := FixedBilling{projectName: "Porject 3", biddedAmount: 600}

	incomeStreams := []Income{project1, project2, project3}
	calculateIncome(incomeStreams)

}
