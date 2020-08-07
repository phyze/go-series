
type PetComposition struct {
	Cat Cat
	Dog Dog
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
		Cat: cat,
		Dog: dog,
	}
	fmt.Println("cat name's ", petCom.Cat.Name)
	fmt.Println("dog name's ", petCom.Dog.Name)
}
