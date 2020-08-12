package pet

import "github.com/go-series/oop/abstraction/animal"

type IPet interface {
	animal.IMammal
}

func NewDog() IPet {
	return &dog{}
}

type dog struct {
	animal.Mammal
}

func (d *dog) Eat() {  }

func (d *dog) Sound() string {
	return d.MakeSound("pet")
}
