package main

import (
	"fmt"
	"sync"
)

func helloMike(pipe chan   string, wg *sync.WaitGroup) {
	defer wg.Done()
	if s := <- pipe ; s == "Hello Mike" {
		pipe <- "Hi"
	}
}

func main() {
	wg := &sync.WaitGroup{}
	pipe := make(chan string,1)
	wg.Add(1)
	pipe <- "Hello Mike"
	go helloMike(pipe,wg)
	wg.Wait()
	fmt.Println(<-pipe)
}
