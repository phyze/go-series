package operation

import (
	"github.com/go-series/pointer/money"
)

func Balance() int {
	return *money.Value
}

func WithDraw(amount int) {
	v := *money.Value
	v = v - amount
	money.Value = &v
}
