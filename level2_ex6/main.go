package main

import (
	"fmt"
)

const (
	_ = iota
	a = 2018 + iota
	b
	c
	d
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
