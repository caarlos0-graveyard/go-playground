package main

import "fmt"

type point struct {
	x, y int
}

type rect struct {
	pos    point
	width  int
	height int
}

func NewPoint(x, y int) point {
	return point{x, y}
}

func (r rect) area() int {
	return r.width * r.height
}

type circle struct {
	radius float64
}

func (c circle) diameter() float64 {
	return c.radius * 2
}

func main() {
	p := point{20, 40}
	fmt.Println(p)

	r := rect{
		pos:    NewPoint(20, 40),
		width:  300,
		height: 500,
	}
	fmt.Println(r)
	fmt.Println(r.area())

	c := circle{
		radius: 20,
	}
	fmt.Println(c.diameter())
}
