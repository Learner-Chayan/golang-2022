package main

import "fmt"

func main() {
	//fmt.Println("Hello world")
	count := map[int]string{1:"Apple",2:"Banana",3:"Watermelon",4:"Jackgfruit"}
	for index,value := range count{
		//fmt.Println(index)
		fmt.Println(value)
		
			for i := 0; i < 10; i++ {
				fmt.Print(i)
				if index %2 ==0 && i > 4 { 
					break
				}
			}
		fmt.Println(" ")
	}
}
