
package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Print("hello ")
}

func main() {
	go hello()
	time.Sleep(time.Millisecond * 50)
	fmt.Println("world")
}

// hello world

