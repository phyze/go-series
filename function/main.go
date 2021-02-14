package main

import "fmt"

func pick(s []string) func(pick string) []string {
	return func(p string) []string {
		temp := make([]string, 0)
		for _, v := range s {
			if v == p {
				temp = append(temp, v)
			}
		}
		return temp
	}
}

func main() {
	//anonymous function
	func() {
		fmt.Println("i'm anonymous function")
	}()
	fmt.Println()

	// 1 st class function
	anonymous := func() {
		fmt.Println("i'm 1st class function")
	}
	anonymous()
	fmt.Println()

	// 1 st class function with arg
	greet := func(name string) {
		fmt.Println("1st class function pass arg : ","hello", name)
	}
	greet("naomi scott")
	fmt.Println()

	// 1 st class function with arg and return result

	resutl := func(a, b int) int {
		return a + b
	}(1, 2)
	fmt.Println("pass arg to 1st class function and return result : ","1 + 2 = ", resutl)
	fmt.Println()

	// higher order function with return function
	fmt.Println("higher order function")
	s := []string{"a", "a", "b", "c", "d", "d", "e"}
	picked := pick(s)("a")
	fmt.Println(picked)
	fmt.Println()

}
