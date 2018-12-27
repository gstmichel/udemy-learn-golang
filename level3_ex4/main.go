package main

import "fmt"

func main() {
	dob := 1985
	for {
		fmt.Println(dob)
		dob++
		if dob > 2018 {
			break
		}
	}
}
