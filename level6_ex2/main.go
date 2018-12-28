package main

import "fmt"

func main() {
	x := []int{5, 10, 15}
	fmt.Println(foo(x...))
	fmt.Println(bar(x))
}

func foo(i ...int) int {
	result := 0

	for _, x := range i {
		result += x
	}

	return result
}

func bar(i []int) int {
	result := 0

	for _, x := range i {
		result += x
	}

	return result
}
