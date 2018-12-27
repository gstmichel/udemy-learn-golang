package main

import "fmt"

func main() {
	vals := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, v := range vals {
		fmt.Println(v)
	}

	fmt.Printf("%T\n", vals)
}
