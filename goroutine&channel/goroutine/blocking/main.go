package main

import (
	"fmt"
	"sync"
)

func hello(wg *sync.WaitGroup) {
	fmt.Print("hello ")
	wg.Done()
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go hello(&wg)
	wg.Wait()
	fmt.Println("world")
}

// hello world
