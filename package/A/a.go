package A

import (
	"fmt"

	"github.com/go-series/package/temp"
)


func SetKeep() {
	s := "hello"
	temp.Keep = s
}


func Print() {
	fmt.Println(temp.Keep)
}
