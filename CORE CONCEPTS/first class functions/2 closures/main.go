package main

import "fmt"

func appendStr() func(paramStr string) string {
	str := "Hello"

	f := func(paramStr string) string {
		return str + " " + paramStr
	}

	return f
}

func main() {

	f1 := appendStr()
	f2 := appendStr()

	fmt.Println(f1("World"))
	fmt.Println(f2("Everyone"))

	fmt.Println(f1("Chayan"))
	fmt.Println(f2("!"))
}
