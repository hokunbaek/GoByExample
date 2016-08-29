package main

import "fmt"

type rect struct {
	width  int
	height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	var r rect
	r.width = 10
	r.height = 5

	fmt.Println("Area: ", r.area())
	fmt.Println("Perim: ", r.perim())

	rp := &r
	fmt.Println("Area: ", rp.area())
	fmt.Println("Perim: ", rp.perim())
}
