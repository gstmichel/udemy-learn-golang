package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	truck := truck{
		vehicle{4, "black"},
		false,
	}

	sedan := sedan{
		vehicle{2, "white"},
		false,
	}

	fmt.Println(truck)
	fmt.Println(sedan)
}
