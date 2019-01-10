package main

import "fmt"

func main() {
	vals := [5]int{1, 2, 3, 4, 5}

	for _, v := range vals {
		fmt.Println(v)
	}

	fmt.Printf("%T\n", vals)
}
