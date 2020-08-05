package main

import (
	"fmt"
)

func main() {
	fmt.Println("-- represent pointer how it work")
	printAddress(5)
	fmt.Println()

	fmt.Println("-- how to use pointer with struct")
	howToUsePtrWithStrcut()
	fmt.Println()
	
	fmt.Println("-- dereference pointer")
	dereference()

}

func printAddress(n int) {
	var a int
	var b *int
	a = n
	b = &a
	fmt.Printf("a pointed to %v , address is %p \n", a, &a)
	fmt.Printf("value of b that kept address of a , value is %p  \n", b)
	fmt.Printf("b pointed to  %p  , address is %p \n", b, &b)
}

func howToUsePtrWithStrcut() {
	type Person struct {
		Name string
		Age  uint8
	}

	type Employee struct {
		ID     string
		Person *Person
	}

	person := &Person{
		Name: "nanao chan",
		Age:  18,
	}

	employee := &Employee{
		Person: person,
		ID:     "ACG10101",
	}
	fmt.Println(employee)
}

func dereference() {
	var number *int
	a := 5
	number = &a
	fmt.Println("before dereference pointer : ", number)
	fmt.Println("after dereference pointer : ", *number)
}
