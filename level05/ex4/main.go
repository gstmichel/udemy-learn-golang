package main

import "fmt"

func main() {
	truck := struct {
		doors int
		color string
	}{
		4,
		"black",
	}

	fmt.Println(truck)
}
