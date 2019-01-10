package main

import "fmt"

type person struct {
	first       string
	last        string
	favIceCream []string
}

func main() {
	person := map[string]person{
		"St-Michel": {
			"Guiaume",
			"St-Michel",
			[]string{"Vanilla", "Nature"},
		},
		"Venkatesan": {
			"Sadhana",
			"Venkatesan",
			[]string{"Chocolate", "Caramel"},
		},
	}

	fmt.Println(person["St-Michel"].first, person["St-Michel"].last)
	for _, s := range person["St-Michel"].favIceCream {
		fmt.Println(s)
	}

	fmt.Println(person["Venkatesan"].first, person["Venkatesan"].last)
	for _, s := range person["Venkatesan"].favIceCream {
		fmt.Println(s)
	}
}
