package main

import "fmt"

type rect struct {
	w, h int
}

func (r *rect) area() int {
	return r.w * r.h
}

func (r rect) perim() int {
	return 2*r.w + 2*r.h
}

func main() {
	r := rect{w: 10, h: 5}

	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", r.perim())
}
