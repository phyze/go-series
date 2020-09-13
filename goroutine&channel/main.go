package main

import (
	"fmt"
)

func main() {
	a := make(chan int)
	go i(a,1)
	go i(a,2)
	go i(a,3)
	a <- 3
	fmt.Println("hello")
}

func i(a <-chan int,i int) {
	fmt.Println(<-a,i)
}


