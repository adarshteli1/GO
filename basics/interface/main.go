package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}
type Rect struct {
	a, b float64
}
type Circle struct {
	r float64
}

func (r Rect) Area() float64 {
	return r.a * r.b
}
func (c Circle) Area() float64 {
	return math.Pi * c.r * c.r
}

func CalArea(s Shape) float64 {
	return s.Area()
}

func main() {
	rect := Rect{a: 5, b: 3}
	cir := Circle{r: 3}

	fmt.Println("Area of Rect", CalArea(rect))

	fmt.Println("Area of Circle", CalArea(cir))
}
