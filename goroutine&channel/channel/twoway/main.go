package main

import (
	"fmt"
	"sync"
)

type twoWay struct {
	hello string
	world string
}

func hello(pipe chan *twoWay, wg *sync.WaitGroup) {
	defer wg.Done()
	v := <- pipe
	v.hello = "hello"
	pipe <- v
}

func world(pipe chan *twoWay, wg *sync.WaitGroup){
	defer wg.Done()
	v := <- pipe
	v.world = "world"
	pipe <- v
}

func main() {
	wg := &sync.WaitGroup{}
	pipe := make(chan *twoWay,1)

	//send twoWay for init
	pipe <- &twoWay{
		hello: "",
		world: "",
	}

	wg.Add(2)
	go hello(pipe,wg)
	go world(pipe,wg)
	wg.Wait()
	fmt.Println(<- pipe)
}
