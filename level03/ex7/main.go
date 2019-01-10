package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		if i%3 == 0 {
			fmt.Printf("%v is multiple of 3\n", i)
		} else if i%3 == 1 {
			fmt.Printf("%v - 1 is multiple of 3\n", i)
		} else {
			fmt.Printf("%v + 1 is multiple of 3\n", i)
		}
	}
}
