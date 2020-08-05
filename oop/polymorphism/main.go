package main

import "fmt"

type PetComposition struct {
	cat Cat
	dog Dog
}

type Dog struct {
	Name string
}

type Cat struct {
	Name string
}

func main() {
	dog := Dog{Name: "Ivory"}
	cat := Cat{Name: "Ebony"}
	petCom := PetComposition{
		cat: cat,
		dog: dog,
	}
	fmt.Println("cat name's ", petCom.cat.Name)
	fmt.Println("dog name's ", petCom.dog.Name)
}
