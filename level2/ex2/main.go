package main

import "fmt"

func main() {
	a := 10 == 9
	b := 4 <= 4
	c := 4 >= 5
	d := 10 != 9
	e := 4 < 4
	f := 4 > 5

	fmt.Println(a, b, c, d, e, f)
}
