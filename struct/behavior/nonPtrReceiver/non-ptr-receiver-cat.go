package nonPtrReceive

import "fmt"

type Cat struct {
	Name string
}

func (c Cat) Run() {
	fmt.Println(c.Name, "...run")
	c.Name = "alter"
}
