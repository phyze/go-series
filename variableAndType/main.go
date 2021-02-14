package main

import (
	"fmt"
)

var age int

func main() {

	//static type
	age = 15
	fmt.Println("age : ", age)

	// dynamic type
	myName := "gram"
	fmt.Println("my name is : ", myName)
}
