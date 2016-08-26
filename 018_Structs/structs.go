package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	p := person{name: "jake", age: 10}
	fmt.Println(p)

	p.name = "jim"
	p.age = 20
	fmt.Println(p)

	var p2 person
	p2.name = "mike"
	p2.age = 30
	fmt.Println(p2)
}
