package main

import (
	"fmt"
)

func greet(pipe chan<- string) {
	pipe <- "Hi Puge"
}

func hiPuge(pipe <-chan string) {
	if s := <-pipe; s == "Hi Puge" {
		fmt.Println(s)
	}
}

func main() {
	pipe := make(chan string)
	go greet(pipe)
	hiPuge(pipe)
}
