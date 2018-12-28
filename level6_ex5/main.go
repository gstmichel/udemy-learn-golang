package main

import (
	"fmt"
	"math"
)

type Square struct {
	length float64
	height float64
}

type Circle struct {
	radius float64
}

type Shape interface {
	area() float64
}

func main() {
	info(Square{5, 10})
	info(Circle{5})
}

func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func (s Square) area() float64 {
	return s.length * s.height
}

func info(s Shape) {
	fmt.Printf("The Area is %v\n", s.area())
}