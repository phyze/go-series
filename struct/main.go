package main

import "fmt"

type Cat struct {
}

func (c Cat) Run() {
	fmt.Println("run..")
}

func main() {
	var cat1 Cat
	cat1.Run()

	cat2 := new(Cat)
	cat2.Run()

	fmt.Printf("cat1 pointer is :%v , cat2 pointer is :%v",cat1,*cat2)
	
}
