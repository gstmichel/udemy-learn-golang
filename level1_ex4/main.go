package main

import "fmt"

type guillaume int

var x guillaume

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}
