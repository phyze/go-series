package main

import (
	"fmt"
)

func main() {
	fmt.Println("--- pointer ---")
	printAddress(5)
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
