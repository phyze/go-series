package b

import (
	"fmt"

	"github.com/go-series/package/a"
)

func PrintAddrsI() {
	fmt.Printf("I addrs is : %p , value is : %v", &a.I, a.I)
}
