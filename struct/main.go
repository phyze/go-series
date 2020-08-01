package main

import "fmt"

type Cat struct {
	Name string
}

func (c Cat) Run() {
	fmt.Println(c.Name, ".. run")
	c.Name = "alter"
}

func main() {
	cat := Cat{Name: "nano"}
	cat.Run()
	cat.Run()
}
