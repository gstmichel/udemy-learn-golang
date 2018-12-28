package main

import "fmt"

type person struct {
	first, last string
	age int
}

func main() {
	x := person{"Guillaume", "St-Michel", 33}
	x.speak()
}

func (p person) speak() {
	fmt.Printf("Hi, my name is %v %v, I am %v years old.\n", p.first, p.last, p.age)
}