package coin

var value int = 5
var version int

func SetValue(n int) {
	value = n
	version = version + 1
}

func Version() int { return version }

func GetValue() int {
	return value
}
