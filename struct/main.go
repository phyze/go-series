package main

import "fmt"

type Cat struct {
	name *string
}

func (c Cat) Run() {
	fmt.Println(*c.name, ".. run")
	name := "alter"
	c.name = &name
}

func main() {
	cat := Cat{}
	name := "hello"
	cat.name = &name
	cat.Run()
	cat.Run()
}
