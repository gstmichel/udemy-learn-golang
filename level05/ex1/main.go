package main

import "fmt"

type person struct {
	first       string
	last        string
	favIceCream []string
}

func main() {
	p1 := person{
		"Guiaume",
		"St-Michel",
		[]string{"Vanilla", "Nature"},
	}

	p2 := person{
		"Sadhana",
		"Venkatesan",
		[]string{"Chocolate", "Caramel"},
	}

	fmt.Println(p1.first, p1.last)
	for _, s := range p1.favIceCream {
		fmt.Println(s)
	}

	fmt.Println(p2.first, p2.last)
	for _, s := range p2.favIceCream {
		fmt.Println(s)
	}
}
