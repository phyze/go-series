package main

import (
	"fmt"
	"github.com/go-series/oop/abstraction/animal/pet"
)

func main() {
	dog := pet.NewDog()
	fmt.Println(dog.Sound())
}