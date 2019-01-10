package main

import "fmt"

func main() {
	execution(func() {
		fmt.Println("In the callback")
	})
}

func execution(callback func()) {
	fmt.Println("executing...")
	callback()
}
