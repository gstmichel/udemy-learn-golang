package main

import "fmt"

func main() {
	x := funcFactory()

	x()
}

func funcFactory() func() {
	return func() {
		fmt.Println("Hello World factory")
	}
}
