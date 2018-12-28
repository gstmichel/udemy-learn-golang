package main

import "fmt"

func main() {
	vals := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	fmt.Println(vals[:5])
	fmt.Println(vals[5:])
	fmt.Println(vals[2:7])
	fmt.Println(vals[1:6])
}
