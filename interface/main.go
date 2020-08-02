package main

import "fmt"

type Walking interface {
	Walk()
}

type Running interface {
	Run()
}

type IDuck interface {
	Walking
	Running
}

type Duck struct {
}

func (d *Duck) Run() {
	fmt.Println("duck run..")
}


//ลอง comment method นี้จะขึ้น error ว่ายังไม่ได้ทำการ impliment 
func (d *Duck) Walk() {
	fmt.Println("duck walk..")
}

func main() {
	var duck IDuck
	duck = &Duck{}
	duck.Run()
	duck.Walk()
}
