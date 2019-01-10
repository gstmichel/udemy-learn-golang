package main

import "fmt"

func main() {
	x := []int{5, 10, 15}
	defer fmt.Printf("Defered calcul : %v\n", foo(x...))

	fmt.Println("After defer, but...")
}

func foo(i ...int) int {
	result := 0

	for _, x := range i {
		result += x
	}

	return result
}
