package main

import "fmt"

type Person struct {
	first, last string
}

func main() {
	x := Person{"Name", "Other"}
	fmt.Println(x)

	changeMe(&x)
	fmt.Println(x)
}

func changeMe(p *Person) {
	(*p).last = "Oups"
	//p.last = "ok ?" // also work...
}
