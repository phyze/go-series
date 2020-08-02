package main

import (
	"fmt"
	"time"

	"github.com/go-series/package/A"
	"github.com/go-series/package/temp"
)


func main() {
	go A.SetKeep()
	time.Sleep(time.Second)
	fmt.Println(temp.Keep)
	A.Print()
}
