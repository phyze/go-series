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
	//-------------------------
	fmt.Println("anonymous function")
	anonymous := func() {
		fmt.Println("i'm anonymous function")
	}
	anonymous()
	fmt.Println()

	//-------------------------
	fmt.Println("pass arg to anonymous function")
	greet := func(name string) {
		fmt.Println("hello", name)
	}
	greet("naomi scott")
	fmt.Println()

	//-------------------------
	fmt.Println("pass arg to anonymous function and return result")
	resutl := func(a, b int) int {
		return a + b
	}(1, 2)
	fmt.Println("1 + 2 = ", resutl)
	fmt.Println()

	//-------------------------
	fmt.Println("higher order function")
	s := []string{"a", "a", "b", "c", "d", "d", "e"}
	picked := pick(s)("a")
	fmt.Println(picked)
}
