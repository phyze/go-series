package a

import "fmt"

var I = 5

func PrintAddrsI() {
	fmt.Printf("i addrs is : %p , value is : %v", &I, I)
}
