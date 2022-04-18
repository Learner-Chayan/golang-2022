package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"log"
)

func main() {
	fmt.Println("Hello world")

	cr := &Person{
		Name: "Chayan",
		Age:  25,
	}

	data, err := proto.Marshal(cr)
	if err != nil {
		log.Fatal("Error", err)
	}

	fmt.Println(data)
	newCr := &Person{}

	err = proto.Unmarshal(data, newCr)
	if err != nil {
		log.Fatal("Error", err)
	}

	fmt.Println(newCr.GetAge())
}
