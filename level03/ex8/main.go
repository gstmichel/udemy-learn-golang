package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		switch {
		case i%2 == 0:
			fmt.Printf("%v is even\n", i)
		case i%2 == 1:
			fmt.Printf("%v is odd\n", i)
		}
	}
}
