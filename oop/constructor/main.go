package main

import (
	"fmt"

	"github.com/go-series/oop/constructor/box"
)

func main() {
	box := box.NewBox(50, 50)
	fmt.Println(box)
}
