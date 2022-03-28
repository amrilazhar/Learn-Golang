package main

import "fmt"

type square struct {
	sideLength float64
}
type triangle struct {
	height float64
	base   float64
}

type shape interface {
	getArea() float64
}

func main() {
	sq := square{6.5}
	tr := triangle{7.5, 10}

	printArea(sq)
	printArea(tr)
}

func printArea(sh shape) {
	fmt.Println(sh.getArea())
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}
