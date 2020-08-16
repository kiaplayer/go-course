package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

func (s triangle) getArea() float64 {
	return 0.5 * s.base * s.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s shape) {
	fmt.Println("Shape area is ", s.getArea())
}

func main() {
	t := triangle{height: 14, base: 8}
	printArea(t)

	s := square{sideLength: 10}
	printArea(s)
}
