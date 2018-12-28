package main

import (
	"fmt"
)

func main() {
	x := make([][]string, 2, 2)

	x[0] = []string{"James", "Bond", "Shaken, not stirred"}
	x[1] = []string{"Miss", "Moneypenny", "Helloooooo, James."}

	for _, i := range x{
		for _, j := range i {
			fmt.Println(j)
		}
	}
}
