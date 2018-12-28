package main

import "fmt"

func main() {
	x := count()
	y := count()

	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())

	fmt.Println(y())
	fmt.Println(y())
	fmt.Println(y())
	fmt.Println(y())
	fmt.Println(y())
}

func count() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}