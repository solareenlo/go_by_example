package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	w, h float64
}

type circle struct {
	r float64
}

func (r rect) area() float64 {
	return r.w * r.h
}

func (r rect) perim() float64 {
	return 2*r.w + 2*r.h
}

func (c circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.r
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{w: 3, h: 4}
	c := circle{r: 5}

	measure(r)
	measure(c)
}
