package main

import "fmt"

type A struct{}

func (a *A) Println() {
	fmt.Println("A")
}

type B struct {
	A
}

func (b *B) Println() {
	fmt.Println("B")
}

func main() {
	b := B{}
	b.Println()
}
