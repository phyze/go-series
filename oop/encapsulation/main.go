package main

import (
	"fmt"

	coinPrivate "github.com/go-series/oop/encapsulation/private/coin"
	"github.com/go-series/oop/encapsulation/public/coin"
)

func main() {
	fmt.Println("------------- public -----------")
	fmt.Println("before value change : ", coin.Value)
	coin.Value = 6
	fmt.Println("after value changed : ", coin.Value)
	fmt.Println()

	fmt.Println("------------- private ----------")
	fmt.Println("before  value change : ", coinPrivate.GetValue())
	coinPrivate.SetValue(8)
	coinPrivate.SetValue(81)
	fmt.Println("after value change : ", coinPrivate.GetValue())
	fmt.Println("version had change : ", coinPrivate.Version())
}
